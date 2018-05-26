package stub

import (
	"context"
	"fmt"
	"os"

	"github.com/neilpeterson/tweet-factory/pkg/apis/tweet-factory/v1alpha1"

	"github.com/operator-framework/operator-sdk/pkg/sdk"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"

	batchv1 "k8s.io/api/batch/v1"
	apiv1 "k8s.io/api/core/v1"
)

// NewHandler function
func NewHandler() sdk.Handler {
	return &Handler{}
}

// Handler function
type Handler struct {
	// Fill me
}

// Handle function
func (h *Handler) Handle(ctx context.Context, event sdk.Event) error {
	switch o := event.Object.(type) {
	case *v1alpha1.TweetFactory:

		// TODO - run Helm uninstall
		if event.Deleted {
			return nil
		}

		// Build Helm installation string
		runCommand := buildString(*o)

		// Run twitter-sentiment chart
		newTwitterSentiment(runCommand, o.Name)
	}
	return nil
}

// newTwitterSentiment creates a new instance of the twitter-sentiment application.
// Quick hack - currently starting a job / container to run Helm install.
// TODO - update to use a go package (k8s.io/helm/pkg/helm).
func newTwitterSentiment(s []string, n string) {

	// Initilization information - package rest
	var (
		config *rest.Config
		err    error
	)

	kubeconfig := os.Getenv("KUBECONFIG")

	// Authentication / connection object - package tools/clientcmd
	config, err = clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error creating client: %v", err)
		os.Exit(1)
	}

	// Kubernetes client - package kubernetes
	clientset := kubernetes.NewForConfigOrDie(config)

	jobsClient := clientset.BatchV1().Jobs("default")

	job := &batchv1.Job{
		ObjectMeta: metav1.ObjectMeta{
			Name:      n,
			Namespace: "default",
		},
		Spec: batchv1.JobSpec{
			Template: apiv1.PodTemplateSpec{
				Spec: apiv1.PodSpec{
					RestartPolicy: "Never",
					Containers: []apiv1.Container{
						{
							Name:    "demo",
							Image:   "neilpeterson/helm-test",
							Command: s,
						},
					},
				},
			},
		},
	}

	_, err = jobsClient.Create(job)
	if err != nil {
		fmt.Println(err)
	}
}

func buildString(o v1alpha1.TweetFactory) []string {

	s := []string{"helm", "install", "azure-samples/twitter-sentiment"}

	if len(o.Spec.ConsumerKey) > 0 {
		s = append(s, "--set", "consumerKey="+o.Spec.ConsumerKey)
	}
	if len(o.Spec.ConsumerSecret) > 0 {
		s = append(s, "--set", "consumerSecret="+o.Spec.ConsumerSecret)
	}
	if len(o.Spec.AccessToken) > 0 {
		s = append(s, "--set", "accessToken="+o.Spec.AccessToken)
	}
	if len(o.Spec.AccessTokenSecret) > 0 {
		s = append(s, "--set", "accessTokenSecret="+o.Spec.AccessTokenSecret)
	}
	if len(o.Spec.FilterText) > 0 {
		s = append(s, "--set", "filterText="+o.Spec.FilterText)
	}

	// I am setting the next two to use CR name.
	// Quick fix for running multiple instances of twitter-analytics.
	s = append(s, "--set", "resourceGroup=twitter-"+o.Name)
	s = append(s, "--set", "twitterSecretName=twitter-"+o.Name)

	return s
}

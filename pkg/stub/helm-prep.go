package stub

import (
	"github.com/neilpeterson/tweet-factory/pkg/apis/tweet-factory/v1alpha1"
	batchv1 "k8s.io/api/batch/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func prepJob(o v1alpha1.TweetFactory, n string, a string) *batchv1.Job {

	var s []string

	if a == "create" {
		s = stringCreate(o)
	} else {
		s = stringDelete(o)
	}

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

	return job
}

func stringCreate(o v1alpha1.TweetFactory) []string {

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
	s = append(s, "--name", "twitter-"+o.Name)

	return s
}

func stringDelete(o v1alpha1.TweetFactory) []string {
	s := []string{"helm", "delete", "twitter-" + o.Name, "--purge"}
	return s
}

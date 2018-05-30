package stub

import (
	"context"
	"fmt"

	"github.com/neilpeterson/tweet-factory/pkg/apis/tweet-factory/v1alpha1"

	"github.com/operator-framework/operator-sdk/pkg/sdk"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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

		if event.Deleted {
			// Delete Tweet Factory instance
			fmt.Println("Twitter Factory deleted: " + o.Name)
			deleteTwitterSentiment(*o)
		} else {
			// Create Tweet Factory instance
			fmt.Println("Twitter Factory created: " + o.Name)
			newTwitterSentiment(*o)
		}
	}
	return nil
}

// newTwitterSentiment creates a new instance of the twitter-sentiment application.
func newTwitterSentiment(o v1alpha1.TweetFactory) {

	var err error
	clientset := kubeAuth()

	// Prepare Helm install
	jobsClient := clientset.BatchV1().Jobs("default")
	job := prepJob(o, "create")

	// Run Helm install
	_, err = jobsClient.Create(job)
	if err != nil {
		fmt.Println(err)
	}

	// Quick hack for job clean-up
	// Need to improve this - function / go routine
	for {
		job, _ := jobsClient.Get(job.Name, metav1.GetOptions{})
		if job.Status.Succeeded == 1 {
			{
				jobsClient.Delete(job.Name, &metav1.DeleteOptions{})
				break
			}
		}
	}
}

func deleteTwitterSentiment(o v1alpha1.TweetFactory) {

	var err error
	clientset := kubeAuth()

	// Prepare Helm delete
	jobsClient := clientset.BatchV1().Jobs("default")
	job := prepJob(o, "delete")

	// Run Helm delete
	_, err = jobsClient.Create(job)
	if err != nil {
		fmt.Println(err)
	}

	// Quick hack for job clean-up
	// Need to improve this - function / go routine
	for {
		job, _ := jobsClient.Get(job.Name, metav1.GetOptions{})
		if job.Status.Succeeded == 1 {
			{
				jobsClient.Delete(job.Name, &metav1.DeleteOptions{})
				break
			}
		}
	}
}

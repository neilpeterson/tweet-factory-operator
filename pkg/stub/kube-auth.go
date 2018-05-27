package stub

import (
	"os"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/kubernetes"
)

func kubeAuth() clientset *Clientset {
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
}
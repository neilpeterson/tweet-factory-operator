package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// TweetFactoryList struct
type TweetFactoryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []TweetFactory `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// TweetFactory struct
type TweetFactory struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              TweetFactorySpec   `json:"spec"`
	Status            TweetFactoryStatus `json:"status,omitempty"`
}

// TweetFactorySpec struct
type TweetFactorySpec struct {
	ConsumerKey       string `json:"consumerKey"`
	ConsumerSecret    string `json:"consumerSecret"`
	AccessToken       string `json:"accessToken"`
	AccessTokenSecret string `json:"accessTokenSecret"`
	FilterText        string `json:"filterText"`
	ResourceGroup     string `json:"resourceGroup"`
	TwitterSecretName string `json:"twitterSecretName"`
}

// TweetFactoryStatus struct
type TweetFactoryStatus struct {
	// Fill me
}

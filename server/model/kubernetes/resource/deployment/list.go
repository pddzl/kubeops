package deployment

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

type DeploymentBrief struct {
	Name              string      `json:"name"`
	NameSpace         string      `json:"namespace"`
	Pods              string      `json:"pods"`
	CreationTimestamp metav1.Time `json:"creationTimestamp,omitempty"`
}

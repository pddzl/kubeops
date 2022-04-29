package pod

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

type PodBrief struct {
	Name              string      `json:"name,omitempty"`
	Namespace         string      `json:"namespace,omitempty"`
	Status            string      `json:"status,omitempty"`
	Node              string      `json:"node,omitempty"`
	CreationTimestamp metav1.Time `json:"creationTimestamp,omitempty"`
}

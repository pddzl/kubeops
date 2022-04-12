package deployment

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

type NewReplicaSet struct {
	Name              string            `json:"name"`
	NameSpace         string            `json:"namespace"`
	Labels            map[string]string `json:"labels,omitempty"`
	CreationTimestamp metav1.Time       `json:"creationTimestamp,omitempty"`
	Replicas          string            `json:"replicas"`
}

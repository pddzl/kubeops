package replicaSet

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

type ReplicaSetBrief struct {
	Name              string      `json:"name"`
	NameSpace         string      `json:"namespace"`
	AvailableReplicas int32       `json:"availableReplicas,omitempty"`
	Replicas          int32       `json:"replicas,omitempty"`
	CreationTimestamp metav1.Time `json:"creationTimestamp,omitempty"`
}

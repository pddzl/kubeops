package replicaSet

import (
	"github.com/pddzl/kubeops/server/model/kubernetes/api"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ReplicaSetDetail struct {
	ObjectMeta api.ObjectMeta   `json:"metadata"`
	Spec       replicaSetSpec   `json:"spec,omitempty"`
	Status     replicaSetStatus `json:"status,omitempty"`
}

type replicaSetSpec struct {
	Replicas        *int32            `json:"replicas,omitempty"`
	MinReadySeconds int32             `json:"minReadySeconds,omitempty"`
	Selector        map[string]string `json:"selector"`
}

type replicaSetStatus struct {
	Replicas             int32                 `json:"replicas"`
	FullyLabeledReplicas int32                 `json:"fullyLabeledReplicas,omitempty"`
	ReadyReplicas        int32                 `json:"readyReplicas,omitempty"`
	AvailableReplicas    int32                 `json:"availableReplicas,omitempty"`
	Conditions           []ReplicaSetCondition `json:"conditions,omitempty"`
}

type ReplicaSetCondition struct {
	Type               string      `json:"type"`
	Status             string      `json:"status"`
	LastTransitionTime metav1.Time `json:"lastTransitionTime,omitempty"`
	Reason             string      `json:"reason,omitempty"`
	Message            string      `json:"message,omitempty"`
}

package deployment

import (
	"github.com/pddzl/kubeops/server/model/kubernetes/api"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

type DeploymentDetail struct {
	ObjectMeta api.ObjectMeta   `json:"metadata"`
	Spec       spec             `json:"spec"`
	Status     deploymentStatus `json:"status,omitempty"`
}

type spec struct {
	Replicas             *int32             `json:"replicas,omitempty"`
	Selector             map[string]string  `json:"selector"`
	Strategy             deploymentStrategy `json:"strategy,omitempty"`
	MinReadySeconds      int32              `json:"minReadySeconds,omitempty"`
	RevisionHistoryLimit *int32             `json:"revisionHistoryLimit,omitempty"`
}

type deploymentStrategy struct {
	DeploymentStrategyType  string                  `json:"type,omitempty"`
	RollingUpdateDeployment rollingUpdateDeployment `json:"rollingUpdate,omitempty"`
}

type rollingUpdateDeployment struct {
	MaxUnavailable *intstr.IntOrString `json:"maxUnavailable,omitempty"`
	MaxSurge       *intstr.IntOrString `json:"maxSurge,omitempty"`
}

type deploymentStatus struct {
	Replicas            int32                 `json:"replicas,omitempty"`
	UpdatedReplicas     int32                 `json:"updatedReplicas,omitempty"`
	ReadyReplicas       int32                 `json:"readyReplicas,omitempty"`
	AvailableReplicas   int32                 `json:"availableReplicas,omitempty"`
	UnavailableReplicas int32                 `json:"unavailableReplicas,omitempty"`
	Conditions          []DeploymentCondition `json:"conditions,omitempty"`
}

type DeploymentCondition struct {
	Type               string      `json:"type"`
	Status             string      `json:"status"`
	LastUpdateTime     metav1.Time `json:"lastUpdateTime,omitempty"`
	LastTransitionTime metav1.Time `json:"lastTransitionTime,omitempty"`
	Reason             string      `json:"reason,omitempty"`
	Message            string      `json:"message,omitempty"`
}

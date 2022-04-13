package daemonSet

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"

	"github.com/pddzl/kubeops/server/model/kubernetes/api"
)

type DeploymentDetail struct {
	ObjectMeta api.ObjectMeta  `json:"metadata,omitempty"`
	Spec       daemonSetSpec   `json:"spec,omitempty"`
	Status     daemonSetStatus `json:"status,omitempty"`
}

type daemonSetSpec struct {
	Selector             map[string]string       `json:"selector"`
	UpdateStrategy       daemonSetUpdateStrategy `json:"updateStrategy,omitempty"`
	MinReadySeconds      int32                   `json:"minReadySeconds,omitempty"`
	RevisionHistoryLimit *int32                  `json:"revisionHistoryLimit,omitempty"`
}

type daemonSetStatus struct {
	CurrentNumberScheduled int32                `json:"currentNumberScheduled"`
	NumberMisscheduled     int32                `json:"numberMisscheduled"`
	DesiredNumberScheduled int32                `json:"desiredNumberScheduled"`
	NumberReady            int32                `json:"numberReady"`
	UpdatedNumberScheduled int32                `json:"updatedNumberScheduled,omitempty"`
	NumberAvailable        int32                `json:"numberAvailable,omitempty"`
	NumberUnavailable      int32                `json:"numberUnavailable,omitempty"`
	Conditions             []DaemonSetCondition `json:"conditions,omitempty"`
}

type DaemonSetCondition struct {
	Type               string      `json:"type"`
	Status             string      `json:"status"`
	LastTransitionTime metav1.Time `json:"lastTransitionTime,omitempty"`
	Reason             string      `json:"reason,omitempty"`
	Message            string      `json:"message,omitempty"`
}

type daemonSetUpdateStrategy struct {
	Type          string                 `json:"type,omitempty"`
	RollingUpdate rollingUpdateDaemonSet `json:"rollingUpdate,omitempty"`
}

type rollingUpdateDaemonSet struct {
	MaxUnavailable *intstr.IntOrString `json:"maxUnavailable,omitempty"`
	MaxSurge       *intstr.IntOrString `json:"maxSurge,omitempty"`
}

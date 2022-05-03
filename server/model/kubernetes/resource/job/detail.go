package job

import (
	"github.com/pddzl/kubeops/server/model/kubernetes/api"
	batchv1 "k8s.io/api/batch/v1"
)

type JobDetail struct {
	Metadata api.ObjectMeta `json:"metadata,omitempty"`
	Spec     jobSpec        `json:"spec,omitempty"`
	Status   jobStatus      `json:"status,omitempty"`
}

type jobSpec struct {
	Parallelism  *int32            `json:"parallelism,omitempty"`
	Completions  *int32            `json:"completions,omitempty"`
	BackoffLimit *int32            `json:"backoffLimit,omitempty"`
	Selector     map[string]string `json:"selector,omitempty"`
}

type jobStatus struct {
	Conditions []batchv1.JobCondition `json:"conditions,omitempty"`
}

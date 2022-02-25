package resource

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type NameSpaceBrief struct {
	Name              string            `json:"name"`
	Labels            map[string]string `json:"labels"`
	CreationTimestamp v1.Time           `json:"creationTimestamp"`
	Status            string            `json:"status"`
}

type NameSpaceDetail struct {
	Metadata NSMetadata `json:"metadata"`
	Status   string     `json:"status"`
	// ResourceQuotaList is list of resource quotas associated to the namespace
	ResourceQuotaList []ResourceQuotaDetail `json:"resourceQuotaList"`

	// ResourceLimits is list of limit ranges associated to the namespace
	ResourceLimits []LimitRangeDetail `json:"resourceLimits"`
}

type NSMetadata struct {
	Name              string            `json:"name"`
	Labels            map[string]string `json:"labels"`
	CreationTimestamp v1.Time           `json:"creationTimestamp"`
	Uid               string            `json:"uid"`
}

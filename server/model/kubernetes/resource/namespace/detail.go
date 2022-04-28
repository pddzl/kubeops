package namespace

import (
	"github.com/pddzl/kubeops/server/model/kubernetes/api"
	"github.com/pddzl/kubeops/server/model/kubernetes/resource"
)

type NameSpaceDetail struct {
	Metadata api.ObjectMeta `json:"metadata,omitempty"`
	Status   string         `json:"status"`
	// ResourceQuotaList is list of resource quotas associated to the namespace
	ResourceQuotaList []resource.ResourceQuotaDetail `json:"resourceQuotaList"`
	// ResourceLimits is list of limit ranges associated to the namespace
	ResourceLimits []resource.LimitRangeDetail `json:"resourceLimits"`
}

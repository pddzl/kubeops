package pod

import (
	"github.com/pddzl/kubeops/server/model/kubernetes/api"
)

type PodRefer struct {
	ObjectMeta api.ObjectMeta `json:"metadata"`
	Status     string         `json:"status"`
	Node       string         `json:"node"`
}

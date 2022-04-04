package replicaSet

import "github.com/pddzl/kubeops/server/model/kubernetes/api"

type Pod struct {
	ObjectMeta api.ObjectMeta `json:"objectMeta"`
	TypeMeta   api.TypeMeta   `json:"typeMeta"`
	Status     string         `json:"status"`
	NodeName   string         `json:"nodeName"`
}

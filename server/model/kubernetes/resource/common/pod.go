package common

import "github.com/pddzl/kubeops/server/model/kubernetes/api"

type RelatedPods struct {
	ObjectMeta api.ObjectMeta `json:"metadata"`
	Status     string         `json:"status"`
	NodeName   string         `json:"nodeName"`
}

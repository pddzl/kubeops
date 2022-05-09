package clusterRole

import metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"

type ClusterRoleBrief struct {
	Name              string      `json:"name,omitempty"`
	CreationTimestamp metaV1.Time `json:"creationTimestamp,omitempty"`
}

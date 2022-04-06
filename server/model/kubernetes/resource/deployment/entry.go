package deployment

import metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"

type DeploymentBrief struct {
	Name              string      `json:"name"`
	NameSpace         string      `json:"namespace"`
	Pods              string      `json:"pods"`
	CreationTimestamp metaV1.Time `json:"creationTimestamp,omitempty"`
}

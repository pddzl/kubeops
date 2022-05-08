package roleBinding

import metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"

type RoleBindingBrief struct {
	Name              string      `json:"name,omitempty"`
	Namespace         string      `json:"namespace,omitempty"`
	CreationTimestamp metaV1.Time `json:"creationTimestamp,omitempty"`
}

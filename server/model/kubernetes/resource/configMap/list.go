package configMap

import metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"

type ConfigMapBrief struct {
	Name              string      `json:"name"`
	Namespace         string      `json:"namespace"`
	CreationTimestamp metaV1.Time `json:"creationTimestamp,omitempty"`
}

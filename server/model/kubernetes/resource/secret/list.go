package secret

import metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"

type SecretBrief struct {
	Name              string      `json:"name"`
	Namespace         string      `json:"namespace"`
	Type              string      `json:"type"`
	CreationTimestamp metaV1.Time `json:"creationTimestamp,omitempty"`
}

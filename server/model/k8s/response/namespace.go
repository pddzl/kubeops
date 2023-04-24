package response

import metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"

type NamespaceBrief struct {
	Name              string            `json:"name"`
	Labels            map[string]string `json:"labels"`
	CreationTimestamp metaV1.Time       `json:"creationTimestamp"`
	Status            string            `json:"status"`
}

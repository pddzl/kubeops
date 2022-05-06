package namespace

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

type NamespaceBrief struct {
	Name              string            `json:"name"`
	Labels            map[string]string `json:"labels"`
	CreationTimestamp metav1.Time       `json:"creationTimestamp"`
	Status            string            `json:"status"`
}

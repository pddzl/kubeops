package job

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

type JobBrief struct {
	Name              string      `json:"name,omitempty"`
	Namespace         string      `json:"namespace,omitempty"`
	CreationTimestamp metav1.Time `json:"creationTimestamp,omitempty"`
	Completed         string      `json:"completed"`
}

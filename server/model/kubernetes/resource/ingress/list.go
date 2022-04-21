package ingress

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

type IngressBrief struct {
	Name              string      `json:"name,omitempty"`
	NameSpace         string      `json:"namespace,omitempty"`
	CreationTimestamp metav1.Time `json:"creationTimestamp,omitempty"`
	EndPoints         []string    `json:"endPoints,omitempty"`
	Hosts             []string    `json:"hosts,omitempty"`
}

package ingress

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

type IngressBrief struct {
	Name              string      `json:"name"`
	NameSpace         string      `json:"namespace"`
	CreationTimestamp metav1.Time `json:"creationTimestamp"`
	EndPoints         []string    `json:"endPoints"`
	Hosts             []string    `json:"hosts"`
}

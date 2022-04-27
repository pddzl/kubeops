package ingress

import (
	"github.com/pddzl/kubeops/server/model/kubernetes/api"
)

type IngressDetail struct {
	ObjectMeta api.ObjectMeta `json:"metadata"`
	Spec       ingressSpec    `json:"spec,omitempty"`
	Status     ingressStatus  `json:"status,omitempty"`
}

type ingressSpec struct {
	IngressClassName string        `json:"ingressClassName,omitempty"`
	Rules            []IngressRule `json:"rules,omitempty"`
}

type IngressRule struct {
	Host string               `json:"host,omitempty"`
	HTTP HTTPIngressRuleValue `json:"http,omitempty"`
}

type HTTPIngressRuleValue struct {
	Paths []HTTPIngressPath `json:"paths"`
}

type HTTPIngressPath struct {
	Path     string         `json:"path,omitempty"`
	PathType string         `json:"pathType"`
	Backend  IngressBackend `json:"backend"`
}

type IngressBackend struct {
	Service IngressServiceBackend `json:"service,omitempty"`
}

type IngressServiceBackend struct {
	Name string             `json:"name"`
	Port ServiceBackendPort `json:"port,omitempty"`
}

type ServiceBackendPort struct {
	Name   string `json:"name,omitempty"`
	Number int32  `json:"number,omitempty"`
}

type ingressStatus struct {
	EndPoints []string `json:"endPoints"`
}

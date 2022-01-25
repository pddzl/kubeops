package service

import (
	"github.com/pddzl/kubeops/server/service/exmaple"
	"github.com/pddzl/kubeops/server/service/kubernetes"
	"github.com/pddzl/kubeops/server/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup     system.ServiceGroup
	ExampleServiceGroup    exmaple.ServiceGroup
	KubernetesServiceGroup kubernetes.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)

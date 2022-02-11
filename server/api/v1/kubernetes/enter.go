package kubernetes

import (
	"github.com/pddzl/kubeops/server/service"
)

type ApiGroup struct {
	NodeApi
	PodApi
	NamespaceApi
}

var (
	nodeService      = service.ServiceGroupApp.KubernetesServiceGroup.NodeService
	podService       = service.ServiceGroupApp.KubernetesServiceGroup.PodService
	namespaceService = service.ServiceGroupApp.KubernetesServiceGroup.NamespaceService
)

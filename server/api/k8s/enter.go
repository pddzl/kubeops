package k8s

import "github.com/pddzl/kubeops/server/service"

type ApiGroup struct {
	NodeApi
	ResourceApi
	NamespaceApi
}

var (
	nodeService      = service.ServiceGroupApp.K8sServiceGroup.NodeService
	resourceService  = service.ServiceGroupApp.K8sServiceGroup.ResourceService
	namespaceService = service.ServiceGroupApp.K8sServiceGroup.NamespaceService
)

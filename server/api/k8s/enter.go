package k8s

import "github.com/pddzl/kubeops/server/service"

type ApiGroup struct {
	NodeApi
	ResourceApi
}

var (
	nodeService     = service.ServiceGroupApp.K8sServiceGroup.NodeService
	resourceService = service.ServiceGroupApp.K8sServiceGroup.ResourceService
)

package k8s

import "github.com/pddzl/kubeops/server/service"

type ApiGroup struct {
	NodeApi
}

var (
	nodeService = service.ServiceGroupApp.K8sServiceGroup.NodeService
)

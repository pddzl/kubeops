package kubernetes

import (
	"github.com/pddzl/kubeops/server/service"
)

type ApiGroup struct {
	NodeApi
}

var (
	nodeService = service.ServiceGroupApp.KubernetesServiceGroup.NodeService
)

package kubernetes

import (
	"github.com/pddzl/kubeops/server/service/kubernetes/node"
	"github.com/pddzl/kubeops/server/service/kubernetes/pod"
)

type ServiceGroup struct {
	NodeService node.NodeService
	PodService  pod.PodService
	NamespaceService
}

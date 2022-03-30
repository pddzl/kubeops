package kubernetes

import (
	"github.com/pddzl/kubeops/server/service/kubernetes/namespace"
	"github.com/pddzl/kubeops/server/service/kubernetes/node"
	"github.com/pddzl/kubeops/server/service/kubernetes/pod"
	"github.com/pddzl/kubeops/server/service/kubernetes/replicaSet"
)

type ServiceGroup struct {
	NodeService       node.NodeService
	PodService        pod.PodService
	NamespaceService  namespace.NamespaceService
	ReplicaSetService replicaSet.ReplicaSetService
}

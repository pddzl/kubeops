package kubernetes

import (
	"github.com/pddzl/kubeops/server/service/kubernetes/daemonSet"
	"github.com/pddzl/kubeops/server/service/kubernetes/deployment"
	"github.com/pddzl/kubeops/server/service/kubernetes/ingress"
	"github.com/pddzl/kubeops/server/service/kubernetes/namespace"
	"github.com/pddzl/kubeops/server/service/kubernetes/node"
	"github.com/pddzl/kubeops/server/service/kubernetes/pod"
	"github.com/pddzl/kubeops/server/service/kubernetes/replicaSet"
	"github.com/pddzl/kubeops/server/service/kubernetes/services"
)

type ServiceGroup struct {
	NodeService       node.NodeService
	PodService        pod.PodService
	NamespaceService  namespace.NamespaceService
	ReplicaSetService replicaSet.ReplicaSetService
	DeploymentService deployment.DeploymentService
	DaemonSetService  daemonSet.DaemonSetService
	ServicesService   services.ServicesService
	IngressService    ingress.IngressService
}

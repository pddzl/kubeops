package kubernetes

import (
	"github.com/pddzl/kubeops/server/service"
)

type ApiGroup struct {
	NodeApi
	PodApi
	NamespaceApi
	ReplicaSetApi
	DeploymentApi
	DaemonSetApi
	ServicesApi
	IngressApi
	JobApi
}

var (
	nodeService       = service.ServiceGroupApp.KubernetesServiceGroup.NodeService
	podService        = service.ServiceGroupApp.KubernetesServiceGroup.PodService
	namespaceService  = service.ServiceGroupApp.KubernetesServiceGroup.NamespaceService
	replicaSetService = service.ServiceGroupApp.KubernetesServiceGroup.ReplicaSetService
	deploymentService = service.ServiceGroupApp.KubernetesServiceGroup.DeploymentService
	daemonSetService  = service.ServiceGroupApp.KubernetesServiceGroup.DaemonSetService
	servicesService   = service.ServiceGroupApp.KubernetesServiceGroup.ServicesService
	ingressService    = service.ServiceGroupApp.KubernetesServiceGroup.IngressService
	jobService        = service.ServiceGroupApp.KubernetesServiceGroup.JobService
)

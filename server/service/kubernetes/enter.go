package kubernetes

import (
	"github.com/pddzl/kubeops/server/service/kubernetes/clusterRole"
	"github.com/pddzl/kubeops/server/service/kubernetes/clusterRoleBinding"
	"github.com/pddzl/kubeops/server/service/kubernetes/configMap"
	"github.com/pddzl/kubeops/server/service/kubernetes/daemonSet"
	"github.com/pddzl/kubeops/server/service/kubernetes/deployment"
	"github.com/pddzl/kubeops/server/service/kubernetes/ingress"
	"github.com/pddzl/kubeops/server/service/kubernetes/job"
	"github.com/pddzl/kubeops/server/service/kubernetes/namespace"
	"github.com/pddzl/kubeops/server/service/kubernetes/node"
	"github.com/pddzl/kubeops/server/service/kubernetes/pod"
	"github.com/pddzl/kubeops/server/service/kubernetes/replicaSet"
	role "github.com/pddzl/kubeops/server/service/kubernetes/role"
	"github.com/pddzl/kubeops/server/service/kubernetes/roleBinding"
	"github.com/pddzl/kubeops/server/service/kubernetes/secret"
	"github.com/pddzl/kubeops/server/service/kubernetes/serviceAccount"
	"github.com/pddzl/kubeops/server/service/kubernetes/services"
)

type ServiceGroup struct {
	NodeService               node.NodeService
	PodService                pod.PodService
	NamespaceService          namespace.NamespaceService
	ReplicaSetService         replicaSet.ReplicaSetService
	DeploymentService         deployment.DeploymentService
	DaemonSetService          daemonSet.DaemonSetService
	ServicesService           services.ServicesService
	IngressService            ingress.IngressService
	JobService                job.JobService
	ServiceAccount            serviceAccount.ServiceAccountService
	RoleService               role.RoleService
	RoleBindingService        roleBinding.RoleBindingService
	ClusterRoleService        clusterRole.ClusterRoleService
	ClusterRoleBindingService clusterRoleBinding.ClusterRoleBindingService
	SecretService             secret.SecretService
	ConfigMapService          configMap.ConfigMapService
}

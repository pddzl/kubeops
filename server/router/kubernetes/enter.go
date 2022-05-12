package kubernetes

type RouterGroup struct {
	NamespaceRouter
	NodeRouter
	PodRouter
	ReplicaSetRouter
	DeploymentRouter
	DaemonSetRouter
	ServicesRouter
	IngressRouter
	JobRouter
	ServiceAccountRouter
	RoleRouter
	RoleBindingRouter
	ClusterRoleRouter
	ClusterRoleBindingRouter
	SecretRouter
}

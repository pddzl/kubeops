package kubernetes

type RouterGroup struct {
	NamespaceRouter
	NodeRouter
	PodRouter
	ReplicaSetRouter
	DeploymentRouter
	DaemonSetRouter
}

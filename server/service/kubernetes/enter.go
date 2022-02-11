package kubernetes

type ServiceGroup struct {
	NodeService
	PodService
	NamespaceService
}

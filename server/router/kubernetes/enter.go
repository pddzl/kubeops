package kubernetes

type RouterGroup struct {
	NodeRouter
	PodRouter
	NamespaceRouter
}

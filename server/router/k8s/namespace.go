package k8s

import (
	"github.com/gin-gonic/gin"
	"github.com/pddzl/kubeops/server/api"
)

type NamespaceRouter struct{}

func (nr *NamespaceRouter) InitNamespaceRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	namespaceRouter := Router.Group("namespace")
	namespaceApi := api.ApiGroupApp.K8sApiGroup.NamespaceApi
	{
		namespaceRouter.POST("getNamespaces", namespaceApi.GetNamespaces)
		namespaceRouter.GET("getNamespaceDetail", namespaceApi.GetNamespaceDetail)
	}
	return namespaceRouter
}

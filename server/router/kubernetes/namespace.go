package kubernetes

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/pddzl/kubeops/server/api/v1"
	"github.com/pddzl/kubeops/server/middleware"
)

type NamespaceRouter struct{}

func (n *NamespaceRouter) InitNamespaceRouter(Router *gin.RouterGroup) {
	namespaceRouter := Router.Group("namespace").Use(middleware.OperationRecord())
	namespaceRouterWithoutRecord := Router.Group("namespace")
	namespaceApi := v1.ApiGroupApp.KubernetesApiGroup.NamespaceApi
	{
		namespaceRouter.POST("getNamespaceDetail", namespaceApi.GetNamespaceDetail) // 获取namespace详情
		namespaceRouter.POST("getNamespaceRaw", namespaceApi.GetNamespaceRaw)       // 获取namespace in raw
	}
	{
		namespaceRouterWithoutRecord.POST("getNamespaceList", namespaceApi.GetNamespaceList)        // 获取所有namespace
		namespaceRouterWithoutRecord.GET("getNamespaceOnlyName", namespaceApi.GetNamespaceOnlyName) // 获取所有namespace(only name)
	}
}

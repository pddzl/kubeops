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
	namespaceRouterApi := v1.ApiGroupApp.KubernetesApiGroup.NamespaceApi
	{
		namespaceRouter.POST("getNamespaceDetail", namespaceRouterApi.GetNamespaceDetail) // 获取namespace详情
	}
	{
		namespaceRouterWithoutRecord.POST("getNamespaceList", namespaceRouterApi.GetNamespaceList)        // 获取所有namespace
		namespaceRouterWithoutRecord.GET("getNamespaceOnlyName", namespaceRouterApi.GetNamespaceOnlyName) // 获取所有namespace(only name)
	}
}

package kubernetes

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/pddzl/kubeops/server/api/v1"
)

type NamespaceRouter struct{}

func (n *NamespaceRouter) InitNamespaceRouter(Router *gin.RouterGroup) {
	//namespaceRouter := Router.Group("namespace").Use(middleware.OperationRecord())
	namespaceRouterWithoutRecord := Router.Group("namespace")
	namespaceRouterApi := v1.ApiGroupApp.KubernetesApiGroup.NamespaceApi
	//{
	//	nodeRouter.POST("getNodeDetail", nodeRouterApi.GetNodeDetail) // 获取node详情
	//}
	{
		namespaceRouterWithoutRecord.POST("getNamespaceList", namespaceRouterApi.GetNamespaceList) // 获取所有node
	}
}

package kubernetes

import (
	"github.com/gin-gonic/gin"
	"github.com/pddzl/kubeops/server/middleware"

	v1 "github.com/pddzl/kubeops/server/api/v1"
)

type NodeRouter struct{}

func (s *NodeRouter) InitNodeRouter(Router *gin.RouterGroup) {
	nodeRouter := Router.Group("node").Use(middleware.OperationRecord())
	nodeRouterWithoutRecord := Router.Group("node")
	nodeRouterApi := v1.ApiGroupApp.KubernetesApiGroup.NodeApi
	{
		nodeRouter.POST("getNodeDetail", nodeRouterApi.GetNodeDetail) // 获取node详情
		nodeRouter.POST("getNodeRaw", nodeRouterApi.GetNodeRaw)       // 获取node in raw
		nodeRouter.POST("getNodePods", nodeRouterApi.GetNodePods)     // 获取node pods
	}
	{
		nodeRouterWithoutRecord.POST("getNodeList", nodeRouterApi.GetNodeList) // 获取所有node
	}
}

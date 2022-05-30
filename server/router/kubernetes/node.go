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
	nodeApi := v1.ApiGroupApp.KubernetesApiGroup.NodeApi
	{
		nodeRouter.POST("getNodeDetail", nodeApi.GetNodeDetail) // 获取node详情
		nodeRouter.POST("getNodeRaw", nodeApi.GetNodeRaw)       // 获取node in raw
		nodeRouter.POST("getNodePods", nodeApi.GetNodePods)     // 获取node pods
		nodeRouter.GET("getNodeType", nodeApi.GetNodeType)      // 获取node类型
		nodeRouter.GET("getNodeStatus", nodeApi.GetNodeStatus)  // 获取node状态
	}
	{
		nodeRouterWithoutRecord.POST("getNodeList", nodeApi.GetNodeList) // 获取所有node
	}
}

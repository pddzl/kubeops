package kubernetes

import (
	"github.com/gin-gonic/gin"

	v1 "github.com/pddzl/kubeops/server/api/v1"
)

type NodeRouter struct{}

func (s *NodeRouter) InitNodeRouter(Router *gin.RouterGroup) {
	//nodeRouter := Router.Group("node").Use(middleware.OperationRecord())
	nodeRouterWithoutRecord := Router.Group("node")
	nodeRouterApi := v1.ApiGroupApp.KubernetesApiGroup.NodeApi
	//{
	//
	//}
	{
		nodeRouterWithoutRecord.POST("getNodeList", nodeRouterApi.GetNodeList) // 获取所有node
	}
}

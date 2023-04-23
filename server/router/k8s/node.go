package k8s

import (
	"github.com/gin-gonic/gin"
	"github.com/pddzl/kubeops/server/api"
)

type NodeRouter struct{}

func (nr *NodeRouter) InitNodeRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	nodeRouter := Router.Group("node")
	nodeApi := api.ApiGroupApp.K8sApiGroup.NodeApi
	{
		nodeRouter.POST("getNodes", nodeApi.GetNodes)
		nodeRouter.GET("getNodeDetail", nodeApi.GetNodeDetail)
	}
	return nodeRouter
}

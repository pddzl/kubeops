package kubernetes

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/pddzl/kubeops/server/api/v1"
	"github.com/pddzl/kubeops/server/middleware"
)

type DeploymentRouter struct{}

func (s *NodeRouter) InitDeploymentRouter(Router *gin.RouterGroup) {
	DeploymentRouter := Router.Group("deployment").Use(middleware.OperationRecord())
	DeploymentRouterWithoutRecord := Router.Group("deployment")
	DeploymentApi := v1.ApiGroupApp.KubernetesApiGroup.DeploymentApi
	{
		//DeploymentRouter.POST("getReplicaSetDetail", DeploymentApi)     // 获取deployment详情
		DeploymentRouter.POST("getDeploymentRaw", DeploymentApi.GetDeploymentRaw) // 获取deployment in 编排
	}
	{
		DeploymentRouterWithoutRecord.POST("getDeploymentList", DeploymentApi.GetDeploymentList) // 获取所有deployment
	}
}

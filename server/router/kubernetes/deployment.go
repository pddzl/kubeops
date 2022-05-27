package kubernetes

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/pddzl/kubeops/server/api/v1"
	"github.com/pddzl/kubeops/server/middleware"
)

type DeploymentRouter struct{}

func (d *DeploymentRouter) InitDeploymentRouter(Router *gin.RouterGroup) {
	deploymentRouter := Router.Group("deployment").Use(middleware.OperationRecord())
	deploymentRouterWithoutRecord := Router.Group("deployment")
	deploymentApi := v1.ApiGroupApp.KubernetesApiGroup.DeploymentApi
	{
		deploymentRouter.POST("getDeploymentDetail", deploymentApi.GetDeploymentDetail) // 获取deployment详情
		deploymentRouter.POST("getDeploymentRaw", deploymentApi.GetDeploymentRaw)       // 获取deployment in 编排
		deploymentRouter.POST("getNewReplicaSet", deploymentApi.GetNewReplicaSet)       // 获取deployment关联的replicaSet
		deploymentRouter.POST("deleteDeployment", deploymentApi.DeleteDeployment)       // 删除deployment
	}
	{
		deploymentRouterWithoutRecord.POST("getDeploymentList", deploymentApi.GetDeploymentList) // 获取所有deployment
	}
}

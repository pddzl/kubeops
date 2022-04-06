package kubernetes

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/pddzl/kubeops/server/api/v1"
)

type DeploymentRouter struct{}

func (s *NodeRouter) InitDeploymentRouter(Router *gin.RouterGroup) {
	//DeploymentRouter := Router.Group("deployment").Use(middleware.OperationRecord())
	DeploymentRouterWithoutRecord := Router.Group("deployment")
	DeploymentApi := v1.ApiGroupApp.KubernetesApiGroup.DeploymentApi
	//{
	//	DeploymentRouter.POST("getReplicaSetDetail", DeploymentApi)     // 获取replicaSet详情
	//	DeploymentRouter.POST("getReplicaSetRaw", ReplicaSetApi.GetReplicaSetRaw)           // 获取replicaSet in raw
	//}
	{
		DeploymentRouterWithoutRecord.POST("getDeploymentList", DeploymentApi.GetDeploymentList) // 获取所有replicaSet
	}
}

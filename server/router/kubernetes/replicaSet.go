package kubernetes

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/pddzl/kubeops/server/api/v1"
	"github.com/pddzl/kubeops/server/middleware"
)

type ReplicaSetRouter struct{}

func (s *ReplicaSetRouter) InitReplicaSetRouter(Router *gin.RouterGroup) {
	replicaSetRouter := Router.Group("replicaSet").Use(middleware.OperationRecord())
	replicaSetRouterWithoutRecord := Router.Group("replicaSet")
	ReplicaSetApi := v1.ApiGroupApp.KubernetesApiGroup.ReplicaSetApi
	{
		replicaSetRouter.POST("getReplicaSetDetail", ReplicaSetApi.GetReplicaSetDetail)     // 获取replicaSet详情
		replicaSetRouter.POST("getReplicaSetRaw", ReplicaSetApi.GetReplicaSetRaw)           // 获取replicaSet in raw
		replicaSetRouter.POST("getReplicaSetPods", ReplicaSetApi.GetReplicaSetPods)         // 获取replicaSet pods
		replicaSetRouter.POST("getReplicaSetServices", ReplicaSetApi.GetReplicaSetServices) // 获取replicaSet services
		replicaSetRouter.POST("deleteReplicaSet", ReplicaSetApi.DeleteReplicaSet)           // 删除replicaSet
	}
	{
		replicaSetRouterWithoutRecord.POST("getReplicaSetList", ReplicaSetApi.GetReplicaSetList) // 获取所有replicaSet
	}
}

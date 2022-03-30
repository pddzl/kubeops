package kubernetes

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/pddzl/kubeops/server/api/v1"
	"github.com/pddzl/kubeops/server/middleware"
)

type ReplicaSetRouter struct{}

func (s *NodeRouter) InitReplicaSetRouter(Router *gin.RouterGroup) {
	ReplicaSetRouter := Router.Group("replicaSet").Use(middleware.OperationRecord())
	ReplicaSetRouterWithoutRecord := Router.Group("replicaSet")
	ReplicaSetApi := v1.ApiGroupApp.KubernetesApiGroup.ReplicaSetApi
	{
		//	ReplicaSetRouter.POST("getPodDetail", podRouterApi.GetPodDetail) // 获取pod详情
		ReplicaSetRouter.POST("getReplicaSetRaw", ReplicaSetApi.GetReplicaSetRaw) // 获取replicaSet in raw
		//	ReplicaSetRouter.POST("getPodLog", podRouterApi.GetPodLog)       // 获取pod日志
	}
	{
		ReplicaSetRouterWithoutRecord.POST("getReplicaSetList", ReplicaSetApi.GetReplicaSetList) // 获取所有replicaSet
	}
}

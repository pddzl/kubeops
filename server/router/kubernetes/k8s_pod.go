package kubernetes

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/pddzl/kubeops/server/api/v1"
	"github.com/pddzl/kubeops/server/middleware"
)

type PodRouter struct{}

func (s *NodeRouter) InitPodRouter(Router *gin.RouterGroup) {
	podRouter := Router.Group("pod").Use(middleware.OperationRecord())
	podRouterWithoutRecord := Router.Group("pod")
	podRouterApi := v1.ApiGroupApp.KubernetesApiGroup.PodApi
	{
		podRouter.POST("getPodDetail", podRouterApi.GetPodDetail) // 获取pod详情
		podRouter.POST("getPodRaw", podRouterApi.GetPodRaw)       // 获取pod in raw
		podRouter.POST("getPodLog", podRouterApi.GetPodLog)       // 获取pod日志
	}
	{
		podRouterWithoutRecord.POST("getPodList", podRouterApi.GetPodList) // 获取所有pod

	}
}

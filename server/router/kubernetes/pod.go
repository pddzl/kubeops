package kubernetes

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/pddzl/kubeops/server/api/v1"
	"github.com/pddzl/kubeops/server/middleware"
)

type PodRouter struct{}

func (s *PodRouter) InitPodRouter(Router *gin.RouterGroup) {
	podRouter := Router.Group("pod").Use(middleware.OperationRecord())
	podRouterWithoutRecord := Router.Group("pod")
	podApi := v1.ApiGroupApp.KubernetesApiGroup.PodApi
	{
		podRouter.POST("getPodDetail", podApi.GetPodDetail) // 获取pod详情
		podRouter.POST("getPodRaw", podApi.GetPodRaw)       // 获取pod in raw
		podRouter.POST("getPodLog", podApi.GetPodLog)       // 获取pod日志
		podRouter.POST("deletePod", podApi.DeletePod)       // 删除pod
		podRouter.GET("getPodStatus", podApi.GetPodStatus)  // 获取pod状态
	}
	{
		podRouterWithoutRecord.POST("getPodList", podApi.GetPodList)        // 获取所有pod
		podRouterWithoutRecord.GET("getPodTerminal", podApi.GetPodTerminal) // 获取pod webShell websocket携带不了jwt
	}
}

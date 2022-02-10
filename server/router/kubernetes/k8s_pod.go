package kubernetes

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/pddzl/kubeops/server/api/v1"
)

type PodRouter struct{}

func (s *NodeRouter) InitPodRouter(Router *gin.RouterGroup) {
	//podRouter := Router.Group("pod").Use(middleware.OperationRecord())
	podRouterWithoutRecord := Router.Group("pod")
	podRouterApi := v1.ApiGroupApp.KubernetesApiGroup.PodApi
	//{
	//	podRouter.POST("getPodDetail") // 获取node详情
	//}
	{
		podRouterWithoutRecord.POST("getPodList", podRouterApi.GetPodList)
		podRouterWithoutRecord.POST("getPodLog", podRouterApi.GetPodLog) // 获取pod日志
	}
}

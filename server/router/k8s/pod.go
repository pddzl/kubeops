package k8s

import (
	"github.com/gin-gonic/gin"
	"github.com/pddzl/kubeops/server/api"
)

type PodRouter struct{}

func (pr *PodRouter) InitPodRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	podRouter := Router.Group("pod")
	podApi := api.ApiGroupApp.K8sApiGroup.PodApi
	{
		podRouter.POST("getPods", podApi.GetPods)
	}
	return podRouter
}
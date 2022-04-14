package kubernetes

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/pddzl/kubeops/server/api/v1"
	"github.com/pddzl/kubeops/server/middleware"
)

type DaemonSetRouter struct{}

func (d *DaemonSetRouter) InitDaemonSetRouter(Router *gin.RouterGroup) {
	daemonSetRouter := Router.Group("daemonSet").Use(middleware.OperationRecord())
	daemonSetRouterWithoutRecord := Router.Group("daemonSet")
	daemonSetApi := v1.ApiGroupApp.KubernetesApiGroup.DaemonSetApi
	{
		daemonSetRouter.POST("getDaemonSetDetail", daemonSetApi.GetDaemonSetDetail) // 获取daemonSet详情
		daemonSetRouter.POST("getDaemonSetRaw", daemonSetApi.GetDaemonSetRaw)       // 获取daemonSet in 编排
		daemonSetRouter.POST("getDaemonSetPods", daemonSetApi.GetDaemonSetPods)     // 获取daemonSet关联的pods
	}
	{
		daemonSetRouterWithoutRecord.POST("getDaemonSetList", daemonSetApi.GetDaemonSetList) // 获取所有daemonSet
	}
}

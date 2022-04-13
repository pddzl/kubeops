package kubernetes

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/pddzl/kubeops/server/api/v1"
)

type DaemonSetRouter struct{}

func (d *DaemonSetRouter) InitDaemonSetRouter(Router *gin.RouterGroup) {
	//daemonSetRouter := Router.Group("daemonSet").Use(middleware.OperationRecord())
	daemonSetRouterWithoutRecord := Router.Group("daemonSet")
	daemonSetApi := v1.ApiGroupApp.KubernetesApiGroup.DaemonSetApi
	//{
	//	daemonSetRouter.POST("getDeploymentDetail") // 获取deployment详情
	//	daemonSetRouter.POST("getDeploymentRaw")       // 获取deployment in 编排
	//}
	{
		daemonSetRouterWithoutRecord.POST("getDaemonSetList", daemonSetApi.GetDaemonSetList) // 获取所有daemonSet
	}
}

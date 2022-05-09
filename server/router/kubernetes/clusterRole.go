package kubernetes

import (
	"github.com/gin-gonic/gin"

	v1 "github.com/pddzl/kubeops/server/api/v1"
	"github.com/pddzl/kubeops/server/middleware"
)

type ClusterRoleRouter struct{}

func (c *ClusterRoleRouter) InitClusterRoleRouter(Router *gin.RouterGroup) {
	clusterRoleRouter := Router.Group("clusterRole").Use(middleware.OperationRecord())
	clusterRoleRouterWithoutRecord := Router.Group("clusterRole")
	clusterRoleApi := v1.ApiGroupApp.KubernetesApiGroup.ClusterRoleApi
	{
		clusterRoleRouter.POST("getClusterRoleRaw", clusterRoleApi.GetClusterRoleRaw)       // 获取clusterRole编排
		clusterRoleRouter.POST("getClusterRoleDetail", clusterRoleApi.GetClusterRoleDetail) // 获取clusterRole详情
	}
	{
		clusterRoleRouterWithoutRecord.POST("getClusterRoleList", clusterRoleApi.GetClusterRoleList) // 获取所有clusterRole
	}
}

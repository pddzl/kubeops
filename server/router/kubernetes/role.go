package kubernetes

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/pddzl/kubeops/server/api/v1"
	"github.com/pddzl/kubeops/server/middleware"
)

type RoleRouter struct{}

func (role *RoleRouter) InitRoleRouter(Router *gin.RouterGroup) {
	roleRouter := Router.Group("role").Use(middleware.OperationRecord())
	roleRouterWithoutRecord := Router.Group("role")
	roleRouterApi := v1.ApiGroupApp.KubernetesApiGroup.RoleApi
	{
		roleRouter.POST("getRoleRaw", roleRouterApi.GetRoleRaw)       // 获取pod编排
		roleRouter.POST("getRoleDetail", roleRouterApi.GetRoleDetail) // 获取role详情
	}
	{
		roleRouterWithoutRecord.POST("getRoleList", roleRouterApi.GetRoleList) // 获取所有role
	}
}

package kubernetes

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/pddzl/kubeops/server/api/v1"
	"github.com/pddzl/kubeops/server/middleware"
)

type RoleBindingRouter struct{}

func (r *RoleBindingRouter) InitRoleBindingRouter(Router *gin.RouterGroup) {
	roleBindingRouter := Router.Group("roleBinding").Use(middleware.OperationRecord())
	roleBindingRouterWithoutRecord := Router.Group("roleBinding")
	roleBindingApi := v1.ApiGroupApp.KubernetesApiGroup.RoleBindingApi
	{
		roleBindingRouter.POST("getRoleBindingRaw", roleBindingApi.GetRoleBindingRaw)       // 获取pod编排
		roleBindingRouter.POST("getRoleBindingDetail", roleBindingApi.GetRoleBindingDetail) // 获取roleBinding详情
	}
	{
		roleBindingRouterWithoutRecord.POST("getRoleBindingList", roleBindingApi.GetRoleBindingList) // 获取所有roleBinding
	}
}

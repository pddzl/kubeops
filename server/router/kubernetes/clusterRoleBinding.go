package kubernetes

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/pddzl/kubeops/server/api/v1"
	"github.com/pddzl/kubeops/server/middleware"
)

type ClusterRoleBindingRouter struct{}

func (c *ClusterRoleBindingRouter) InitClusterRoleBindingRouter(Router *gin.RouterGroup) {
	clusterRoleBindingRouter := Router.Group("clusterRoleBinding").Use(middleware.OperationRecord())
	clusterRoleBindingRouterWithoutRecord := Router.Group("clusterRoleBinding")
	clusterRoleBindingApi := v1.ApiGroupApp.KubernetesApiGroup.ClusterRoleBindingApi
	{
		clusterRoleBindingRouter.POST("getClusterRoleBindingRaw", clusterRoleBindingApi.GetClusterRoleBindingRaw)       // 获取clusterRoleBinding编排
		clusterRoleBindingRouter.POST("getClusterRoleBindingDetail", clusterRoleBindingApi.GetClusterRoleBindingDetail) // 获取clusterRoleBinding详情
		clusterRoleBindingRouter.POST("deleteClusterRoleBinding", clusterRoleBindingApi.DeleteClusterRoleBinding)       // 删除clusterRoleBinding
	}
	{
		clusterRoleBindingRouterWithoutRecord.POST("getClusterRoleBindingList", clusterRoleBindingApi.GetClusterRoleBindingList) // 获取所有clusterRoleBinding
	}
}

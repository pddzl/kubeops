package kubernetes

import (
	"github.com/gin-gonic/gin"

	v1 "github.com/pddzl/kubeops/server/api/v1"
	"github.com/pddzl/kubeops/server/middleware"
)

type ResourceRouter struct{}

func (r *ResourceRouter) InitResourceRouter(Router *gin.RouterGroup) {
	resourceRouter := Router.Group("resource").Use(middleware.OperationRecord())
	resourceApi := v1.ApiGroupApp.KubernetesApiGroup.ResourceApi
	{
		resourceRouter.POST("createResource", resourceApi.CreateDynamicResource) // 创建资源
	}
}

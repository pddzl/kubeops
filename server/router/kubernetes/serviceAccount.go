package kubernetes

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/pddzl/kubeops/server/api/v1"
	"github.com/pddzl/kubeops/server/middleware"
)

type ServiceAccountRouter struct{}

func (s *ServiceAccountRouter) InitServiceAccountRouter(Router *gin.RouterGroup) {
	serviceAccountRouter := Router.Group("serviceAccount").Use(middleware.OperationRecord())
	ServiceAccountRouterWithoutRecord := Router.Group("serviceAccount")
	serviceAccountApi := v1.ApiGroupApp.KubernetesApiGroup.ServiceAccountApi
	{
		serviceAccountRouter.POST("getServiceAccountDetail", serviceAccountApi.GetServiceAccountDetail) // 获取serviceAccount详情
		serviceAccountRouter.POST("getServiceAccountRaw", serviceAccountApi.GetServiceAccountRaw)       // 获取serviceAccount编排
		serviceAccountRouter.POST("deleteServiceAccount", serviceAccountApi.DeleteServiceAccount)       // 删除serviceAccount
	}
	{
		ServiceAccountRouterWithoutRecord.POST("getServiceAccountList", serviceAccountApi.GetServiceAccountList) // 获取所有serviceAccount
	}
}

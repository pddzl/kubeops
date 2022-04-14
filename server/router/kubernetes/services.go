package kubernetes

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/pddzl/kubeops/server/api/v1"
	"github.com/pddzl/kubeops/server/middleware"
)

type ServicesRouter struct{}

func (s *ServicesRouter) InitServicesRouter(Router *gin.RouterGroup) {
	servicesRouter := Router.Group("services").Use(middleware.OperationRecord())
	servicesRouterWithoutRecord := Router.Group("services")
	servicesApi := v1.ApiGroupApp.KubernetesApiGroup.ServicesApi
	{
		servicesRouter.POST("getServicesDetail", servicesApi.GetServicesDetail) // 获取services详情
		servicesRouter.POST("getServicesRaw", servicesApi.GetServicesRaw)       // 获取services in raw
		servicesRouter.POST("getServicesPods", servicesApi.GetServicesPods)     // 获取services pods
	}
	{
		servicesRouterWithoutRecord.POST("getServicesList", servicesApi.GetServicesList) // 获取所有services
	}
}

package kubernetes

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/pddzl/kubeops/server/api/v1"
	"github.com/pddzl/kubeops/server/middleware"
)

type IngressRouter struct{}

func (i *IngressRouter) InitIngressRouter(Router *gin.RouterGroup) {
	ingressRouter := Router.Group("ingress").Use(middleware.OperationRecord())
	ingressRouterWithoutRecord := Router.Group("ingress")
	ingressApi := v1.ApiGroupApp.KubernetesApiGroup.IngressApi
	{
		ingressRouter.POST("getIngressDetail", ingressApi.GetIngressDetail) // 获取Ingress详情
	}
	{
		ingressRouterWithoutRecord.POST("getIngressList", ingressApi.GetIngressList) // 获取所有Ingress
	}
}

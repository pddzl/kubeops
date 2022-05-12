package kubernetes

import (
	"github.com/gin-gonic/gin"

	v1 "github.com/pddzl/kubeops/server/api/v1"
	"github.com/pddzl/kubeops/server/middleware"
)

type SecretRouter struct{}

func (s *SecretRouter) InitSecretRouter(Router *gin.RouterGroup) {
	secretRouter := Router.Group("secret").Use(middleware.OperationRecord())
	secretRouterWithoutRecord := Router.Group("secret")
	secretApi := v1.ApiGroupApp.KubernetesApiGroup.SecretApi
	{
		secretRouter.POST("getSecretRaw", secretApi.GetSecretRaw)       // 获取secret编排
		secretRouter.POST("getSecretDetail", secretApi.GetSecretDetail) // 获取secret详情
	}
	{
		secretRouterWithoutRecord.POST("getSecretList", secretApi.GetSecretList) // 获取所有secret
	}
}

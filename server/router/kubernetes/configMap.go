package kubernetes

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/pddzl/kubeops/server/api/v1"
	"github.com/pddzl/kubeops/server/middleware"
)

type ConfigMapRouter struct{}

func (cm *ConfigMapRouter) InitConfigMapRouter(Router *gin.RouterGroup) {
	configMapRouter := Router.Group("configMap").Use(middleware.OperationRecord())
	configMapRouterWithoutRecord := Router.Group("configMap")
	configMapApi := v1.ApiGroupApp.KubernetesApiGroup.ConfigMapApi
	{
		configMapRouter.POST("getConfigMapRaw", configMapApi.GetConfigMapRaw)       // 获取configMap编排
		configMapRouter.POST("getConfigMapDetail", configMapApi.GetConfigMapDetail) // 获取configMap详情
		configMapRouter.POST("deleteConfigMap", configMapApi.DeleteConfigMap)       // 删除configMap
	}
	{
		configMapRouterWithoutRecord.POST("getConfigMapList", configMapApi.GetConfigMapList) // 获取所有configMap
	}
}

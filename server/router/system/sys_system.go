package system

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/pddzl/kubeops/server/api/v1"
	"github.com/pddzl/kubeops/server/middleware"
)

type SysRouter struct{}

func (s *SysRouter) InitSystemRouter(Router *gin.RouterGroup) {
	sysRouter := Router.Group("system").Use(middleware.OperationRecord())
	systemApi := v1.ApiGroupApp.SystemApiGroup.SystemApi
	{
		sysRouter.POST("reloadSystem", systemApi.ReloadSystem) // 重启服务
	}
}

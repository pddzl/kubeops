package kubernetes

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/pddzl/kubeops/server/api/v1"
	"github.com/pddzl/kubeops/server/middleware"
)

type ScaleRouter struct{}

func (s *ScaleRouter) InitScaleRouter(Router *gin.RouterGroup) {
	scaleRouter := Router.Group("scale").Use(middleware.OperationRecord())
	scaleApi := v1.ApiGroupApp.KubernetesApiGroup.ScaleApi
	{
		scaleRouter.POST("", scaleApi.Scale) // deployment, replicaset scale
	}
}

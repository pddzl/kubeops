package system

import (
	"github.com/gin-gonic/gin"
	"github.com/pddzl/kubeops/server/api"
)

type ApiRouter struct{}

func (u *ApiRouter) InitApiRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	apiRouter := Router.Group("api")
	apiApi := api.ApiGroupApp.SystemApiGroup.ApiApi
	{
		apiRouter.POST("addApi", apiApi.AddApi)
		apiRouter.POST("getApis", apiApi.GetApis)
		apiRouter.POST("deleteApi", apiApi.DeleteApi)
		apiRouter.POST("editApi", apiApi.EditApi)
		apiRouter.POST("getElTreeApis", apiApi.GetElTreeApis)
	}
	return apiRouter
}

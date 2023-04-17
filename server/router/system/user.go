package system

import (
	"github.com/gin-gonic/gin"
	"github.com/pddzl/kubeops/server/api"
)

type UserRouter struct{}

func (u *UserRouter) InitUserRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	userRouter := Router.Group("user")
	userApi := api.ApiGroupApp.SystemApiGroup.UserApi
	{
		userRouter.GET("getUserInfo", userApi.GetUserInfo)
		userRouter.POST("getUsers", userApi.GetUsers)
		userRouter.POST("deleteUser", userApi.DeleteUser)
		userRouter.POST("addUser", userApi.AddUser)
		userRouter.POST("editUser", userApi.EditUser)
		userRouter.POST("modifyPass", userApi.ModifyPass)
		userRouter.POST("switchActive", userApi.SwitchActive)
	}
	return userRouter
}

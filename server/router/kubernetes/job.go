package kubernetes

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/pddzl/kubeops/server/api/v1"
	"github.com/pddzl/kubeops/server/middleware"
)

type JobRouter struct{}

func (j *JobRouter) InitJobRouter(Router *gin.RouterGroup) {
	jobRouter := Router.Group("job").Use(middleware.OperationRecord())
	jobRouterWithoutRecord := Router.Group("job")
	jobRouterApi := v1.ApiGroupApp.KubernetesApiGroup.JobApi
	{
		jobRouter.POST("getJobDetail", jobRouterApi.GetJobDetail) // 获取job详情
		//jobRouter.POST("getPodRaw", podRouterApi.GetPodRaw)       // 获取pod in raw
		//jobRouter.POST("getPodLog", podRouterApi.GetPodLog)       // 获取pod日志
	}
	{
		jobRouterWithoutRecord.POST("getJobList", jobRouterApi.GetJobList) // 获取所有job
	}
}

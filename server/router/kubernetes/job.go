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
	jobApi := v1.ApiGroupApp.KubernetesApiGroup.JobApi
	{
		jobRouter.POST("getJobRaw", jobApi.GetJobRaw)       // 获取pod编排
		jobRouter.POST("getJobDetail", jobApi.GetJobDetail) // 获取job详情
		jobRouter.POST("getJobPods", jobApi.GetJobPods)     // 获取job关联的pods
	}
	{
		jobRouterWithoutRecord.POST("getJobList", jobApi.GetJobList) // 获取所有job
	}
}

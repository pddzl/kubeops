package kubernetes

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"

	"github.com/pddzl/kubeops/server/global"
	"github.com/pddzl/kubeops/server/model/common/response"
	"github.com/pddzl/kubeops/server/model/kubernetes/request"
)

type JobApi struct{}

func (j *JobApi) GetJobList(c *gin.Context) {
	var pageInfo request.ListParams
	_ = c.ShouldBindJSON(&pageInfo)
	// 校验字段
	validate := validator.New()
	if err := validate.Struct(&pageInfo.PageInfo); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	list, total, err := jobService.GetJobList(pageInfo.NameSpace, pageInfo.PageInfo)
	if err != nil {
		response.FailWithMessage("获取失败", c)
		global.KOP_LOG.Error("获取失败", zap.Error(err))
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    int64(total),
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

func (j *JobApi) GetJobRaw(c *gin.Context) {
	var job request.JobCommon
	_ = c.ShouldBindJSON(&job)
	// 校验字段
	validate := validator.New()
	if err := validate.Struct(&job); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	raw, err := jobService.GetJobRaw(job.NameSpace, job.Job)
	if err != nil {
		response.FailWithMessage("获取失败", c)
		global.KOP_LOG.Error("获取失败", zap.Error(err))
	} else {
		response.OkWithDetailed(raw, "获取成功", c)
	}
}

func (j *JobApi) GetJobDetail(c *gin.Context) {
	var job request.JobCommon
	_ = c.ShouldBindJSON(&job)
	// 校验字段
	validate := validator.New()
	if err := validate.Struct(&job); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	detail, err := jobService.GetJobDetail(job.NameSpace, job.Job)
	if err != nil {
		response.FailWithMessage("获取失败", c)
		global.KOP_LOG.Error("获取失败", zap.Error(err))
	} else {
		response.OkWithDetailed(detail, "获取成功", c)
	}
}

// DeleteJob 删除job
func (j *JobApi) DeleteJob(c *gin.Context) {
	var jobCommon request.JobCommon
	_ = c.ShouldBindJSON(&jobCommon)
	// 校验字段
	validate := validator.New()
	if err := validate.Struct(&jobCommon); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := jobService.DeleteJob(jobCommon.NameSpace, jobCommon.Job); err != nil {
		response.FailWithMessage(err.Error(), c)
		global.KOP_LOG.Error("删除失败", zap.Error(err))
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

func (j *JobApi) GetJobPods(c *gin.Context) {
	var job request.JobPods
	_ = c.ShouldBindJSON(&job)
	// 校验字段
	validate := validator.New()
	if err := validate.Struct(&job); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	pods, total, err := jobService.GetJobPods(job.NameSpace, job.Job, job.PageInfo)
	if err != nil {
		response.FailWithMessage("获取失败", c)
		global.KOP_LOG.Error("获取失败", zap.Error(err))
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     pods,
			Total:    int64(total),
			Page:     job.Page,
			PageSize: job.PageSize,
		}, "获取成功", c)
	}
}

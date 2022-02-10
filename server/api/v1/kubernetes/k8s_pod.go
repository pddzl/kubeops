package kubernetes

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/pddzl/kubeops/server/global"

	"github.com/pddzl/kubeops/server/model/common/response"
	"github.com/pddzl/kubeops/server/model/kubernetes/request"
)

type PodApi struct{}

func (p *PodApi) GetPodLog(c *gin.Context) {
	var podLog request.PodLog
	_ = c.ShouldBindJSON(&podLog)
	// 字段校验
	validate := validator.New()
	if err := validate.Struct(&podLog); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	log, err := podService.GetPodLog(podLog.NameSpace, podLog.Pod, podLog.Lines)
	if err != nil {
		global.KOP_LOG.Error(err.Error())
		response.FailWithMessage("获取日志失败", c)
		return
	}
	response.OkWithDetailed(log, "获取日志成功", c)
}

package kubernetes

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/pddzl/kubeops/server/global"
	"github.com/pddzl/kubeops/server/model/common/response"
	"github.com/pddzl/kubeops/server/model/kubernetes/request"
	"go.uber.org/zap"
)

type DaemonSetApi struct{}

// 获取deployment list

func (d *DaemonSetApi) GetDaemonSetList(c *gin.Context) {
	var pageInfo request.ListParams
	_ = c.ShouldBindJSON(&pageInfo)
	// 校验字段
	validate := validator.New()
	if err := validate.Struct(&pageInfo.PageInfo); err != nil {
		global.KOP_LOG.Error("请求参数有误", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	list, total, err := daemonSetService.GetDaemonSetList(pageInfo.NameSpace, pageInfo.PageInfo)
	if err != nil {
		response.FailWithMessage("获取失败", c)
		global.KOP_LOG.Error("获取失败", zap.Error(err))
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    int64(total),
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

// 获取deployment in 编排

func (d *DaemonSetApi) GetDaemonSetRaw(c *gin.Context) {
	var daemonSetRaw request.DaemonSetCommon
	_ = c.ShouldBindJSON(&daemonSetRaw)
	// 校验
	validate := validator.New()
	if err := validate.Struct(&daemonSetRaw); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	info, err := daemonSetService.GetDaemonSetRaw(daemonSetRaw.NameSpace, daemonSetRaw.DaemonSet)
	if err != nil {
		response.FailWithMessage("获取失败", c)
		global.KOP_LOG.Error("获取失败", zap.Error(err))
		return
	}
	response.OkWithDetailed(info, "获取成功", c)
}

// 获取daemonSet detail

func (d *DaemonSetApi) GetDaemonSetDetail(c *gin.Context) {
	var daemonSetDetail request.DaemonSetCommon
	_ = c.ShouldBindJSON(&daemonSetDetail)
	// 校验
	validate := validator.New()
	if err := validate.Struct(&daemonSetDetail); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	detail, err := daemonSetService.GetDaemonSetDetail(daemonSetDetail.NameSpace, daemonSetDetail.DaemonSet)
	if err != nil {
		response.FailWithMessage("获取失败", c)
		global.KOP_LOG.Error("获取失败", zap.Error(err))
		return
	}
	response.OkWithDetailed(detail, "获取成功", c)
}

package kubernetes

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"

	"github.com/pddzl/kubeops/server/global"
	"github.com/pddzl/kubeops/server/model/common/response"
	"github.com/pddzl/kubeops/server/model/kubernetes/request"
)

type ServicesApi struct{}

// 获取services列表

func (s *ServicesApi) GetServicesList(c *gin.Context) {
	var pageInfo request.ListParams
	_ = c.ShouldBindJSON(&pageInfo)
	// 校验字段
	validate := validator.New()
	if err := validate.Struct(&pageInfo.PageInfo); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	list, total, err := servicesService.GetServicesList(pageInfo.NameSpace, pageInfo.PageInfo)
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

// 获取services in 编排

func (s *ServicesApi) GetServicesRaw(c *gin.Context) {}

// 获取services详情

func (s *ServicesApi) GetServicesDetail(c *gin.Context) {}

// 获取services关联的pods

func (s *ServicesApi) GetServicesPods(c *gin.Context) {}

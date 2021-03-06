package kubernetes

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"

	"github.com/pddzl/kubeops/server/global"
	"github.com/pddzl/kubeops/server/model/common/response"
	"github.com/pddzl/kubeops/server/model/kubernetes/request"
)

type ServiceAccountApi struct{}

func (sa *ServiceAccountApi) GetServiceAccountList(c *gin.Context) {
	var pageInfo request.ListParams
	_ = c.ShouldBindJSON(&pageInfo)
	// 校验字段
	validate := validator.New()
	if err := validate.Struct(&pageInfo.PageInfo); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	list, total, err := serviceAccountService.GetServiceAccountList(pageInfo.NameSpace, pageInfo.PageInfo)
	if err != nil {
		response.FailWithMessage("获取失败", c)
		global.KOP_LOG.Info("获取失败", zap.Error(err))
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    int64(total),
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

func (sa *ServiceAccountApi) GetServiceAccountRaw(c *gin.Context) {
	var serviceAccount request.ServiceAccountCommon
	_ = c.ShouldBindJSON(&serviceAccount)
	// 校验
	validate := validator.New()
	if err := validate.Struct(&serviceAccount); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	raw, err := serviceAccountService.GetServiceAccountRaw(serviceAccount.NameSpace, serviceAccount.ServiceAccount)
	if err != nil {
		response.FailWithMessage("获取失败", c)
		global.KOP_LOG.Error("获取失败", zap.Error(err))
	} else {
		response.OkWithDetailed(raw, "获取成功", c)
	}
}

func (sa *ServiceAccountApi) GetServiceAccountDetail(c *gin.Context) {
	var serviceAccount request.ServiceAccountCommon
	_ = c.ShouldBindJSON(&serviceAccount)
	// 校验
	validate := validator.New()
	if err := validate.Struct(&serviceAccount); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	detail, err := serviceAccountService.GetServiceAccountDetail(serviceAccount.NameSpace, serviceAccount.ServiceAccount)
	if err != nil {
		response.FailWithMessage("获取失败", c)
		global.KOP_LOG.Error("获取失败", zap.Error(err))
	} else {
		response.OkWithDetailed(detail, "获取成功", c)
	}
}

func (sa *ServiceAccountApi) DeleteServiceAccount(c *gin.Context) {
	var serviceAccount request.ServiceAccountCommon
	_ = c.ShouldBindJSON(&serviceAccount)
	// 校验
	validate := validator.New()
	if err := validate.Struct(&serviceAccount); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := serviceAccountService.DeleteServiceAccount(serviceAccount.NameSpace, serviceAccount.ServiceAccount); err != nil {
		response.FailWithMessage(err.Error(), c)
		global.KOP_LOG.Error("删除失败", zap.Error(err))
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

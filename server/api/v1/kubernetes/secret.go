package kubernetes

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"

	"github.com/pddzl/kubeops/server/global"
	"github.com/pddzl/kubeops/server/model/common/response"
	"github.com/pddzl/kubeops/server/model/kubernetes/request"
)

type SecretApi struct{}

func (s *SecretApi) GetSecretList(c *gin.Context) {
	var pageInfo request.ListParams
	_ = c.ShouldBindJSON(&pageInfo)
	// 校验字段
	validate := validator.New()
	if err := validate.Struct(&pageInfo.PageInfo); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	list, total, err := SecretService.GetSecretList(pageInfo.NameSpace, pageInfo.PageInfo)
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

func (s *SecretApi) GetSecretRaw(c *gin.Context) {
	var secret request.SecretCommon
	_ = c.ShouldBindJSON(&secret)
	// 校验
	validate := validator.New()
	if err := validate.Struct(&secret); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	raw, err := SecretService.GetSecretRaw(secret.Namespace, secret.Secret)
	if err != nil {
		response.FailWithMessage("获取失败", c)
		global.KOP_LOG.Error("获取失败", zap.Error(err))
	} else {
		response.OkWithDetailed(raw, "获取成功", c)
	}
}

func (s *SecretApi) GetSecretDetail(c *gin.Context) {
	var secret request.SecretCommon
	_ = c.ShouldBindJSON(&secret)
	// 校验
	validate := validator.New()
	if err := validate.Struct(&secret); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	detail, err := SecretService.GetSecretDetail(secret.Namespace, secret.Secret)
	if err != nil {
		response.FailWithMessage("获取失败", c)
		global.KOP_LOG.Error("获取失败", zap.Error(err))
	} else {
		response.OkWithDetailed(detail, "获取成功", c)
	}
}

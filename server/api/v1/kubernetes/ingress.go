package kubernetes

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/pddzl/kubeops/server/global"
	"github.com/pddzl/kubeops/server/model/common/response"
	"github.com/pddzl/kubeops/server/model/kubernetes/request"
	"go.uber.org/zap"
)

type IngressApi struct{}

// GetIngressList 获取Ingress列表
func (i *IngressApi) GetIngressList(c *gin.Context) {
	var pageInfo request.ListParams
	_ = c.ShouldBindJSON(&pageInfo)
	// 校验字段
	validate := validator.New()
	if err := validate.Struct(&pageInfo.PageInfo); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	list, total, err := ingressService.GetIngressList(pageInfo.NameSpace, pageInfo.PageInfo)
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

// GetIngressRaw 获取Ingress in raw
func (i *IngressApi) GetIngressRaw(c *gin.Context) {
	var ingress request.IngressCommon
	_ = c.ShouldBindJSON(&ingress)
	// 校验
	validate := validator.New()
	if err := validate.Struct(&ingress); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	raw, err := ingressService.GetReplicaSetRaw(ingress.NameSpace, ingress.Ingress)
	if err != nil {
		response.FailWithMessage("获取失败", c)
		global.KOP_LOG.Error("获取失败", zap.Error(err))
		return
	}
	response.OkWithDetailed(raw, "获取成功", c)
}

// GetIngressDetail 获取Ingress detail
func (i *IngressApi) GetIngressDetail(c *gin.Context) {
	var ingress request.IngressCommon
	_ = c.ShouldBindJSON(&ingress)
	// 校验
	validate := validator.New()
	if err := validate.Struct(&ingress); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	detail, err := ingressService.GetIngressDetail(ingress.NameSpace, ingress.Ingress)
	if err != nil {
		response.FailWithMessage("获取失败", c)
		global.KOP_LOG.Error("获取失败", zap.Error(err))
		return
	}
	response.OkWithDetailed(detail, "获取成功", c)
}

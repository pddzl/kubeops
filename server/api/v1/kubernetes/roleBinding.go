package kubernetes

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/pddzl/kubeops/server/global"
	"github.com/pddzl/kubeops/server/model/common/response"
	"github.com/pddzl/kubeops/server/model/kubernetes/request"
	"go.uber.org/zap"
)

type RoleBindingApi struct{}

func (r *RoleBindingApi) GetRoleBindingList(c *gin.Context) {
	var pageInfo request.ListParams
	_ = c.ShouldBindJSON(&pageInfo)
	// 校验字段
	validate := validator.New()
	if err := validate.Struct(&pageInfo.PageInfo); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	list, total, err := roleBindingService.GetRoleBindingList(pageInfo.NameSpace, pageInfo.PageInfo)
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

func (r *RoleBindingApi) GetRoleBindingRaw(c *gin.Context) {
	var roleBinding request.RoleBindingCommon
	_ = c.ShouldBindJSON(&roleBinding)
	// 校验
	validate := validator.New()
	if err := validate.Struct(&roleBinding); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	raw, err := roleBindingService.GetRoleBindingRaw(roleBinding.Namespace, roleBinding.RoleBinding)
	if err != nil {
		response.FailWithMessage("获取失败", c)
		global.KOP_LOG.Error("获取失败", zap.Error(err))
	} else {
		response.OkWithDetailed(raw, "获取成功", c)
	}
}

func (r *RoleBindingApi) GetRoleBindingDetail(c *gin.Context) {
	var roleBinding request.RoleBindingCommon
	_ = c.ShouldBindJSON(&roleBinding)
	// 校验
	validate := validator.New()
	if err := validate.Struct(&roleBinding); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	detail, err := roleBindingService.GetRoleBindingDetail(roleBinding.Namespace, roleBinding.RoleBinding)
	if err != nil {
		response.FailWithMessage("获取失败", c)
		global.KOP_LOG.Error("获取失败", zap.Error(err))
	} else {
		response.OkWithDetailed(detail, "获取成功", c)
	}
}

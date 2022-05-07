package kubernetes

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/pddzl/kubeops/server/global"
	"go.uber.org/zap"

	"github.com/pddzl/kubeops/server/model/common/response"
	"github.com/pddzl/kubeops/server/model/kubernetes/request"
)

type RoleApi struct{}

func (r *RoleApi) GetRoleList(c *gin.Context) {
	var pageInfo request.ListParams
	_ = c.ShouldBindJSON(&pageInfo)
	// 校验字段
	validate := validator.New()
	if err := validate.Struct(&pageInfo.PageInfo); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	list, total, err := roleService.GetRoleList(pageInfo.NameSpace, pageInfo.PageInfo)
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

func (r *RoleApi) GetRoleRaw(c *gin.Context) {
	var role request.RoleCommon
	_ = c.ShouldBindJSON(&role)
	// 校验
	validate := validator.New()
	if err := validate.Struct(&role); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	raw, err := roleService.GetRoleRaw(role.Namespace, role.Role)
	if err != nil {
		response.FailWithMessage("获取失败", c)
		global.KOP_LOG.Error("获取失败", zap.Error(err))
	} else {
		response.OkWithDetailed(raw, "获取成功", c)
	}
}

func (r *RoleApi) GetRoleDetail(c *gin.Context) {
	var role request.RoleCommon
	_ = c.ShouldBindJSON(&role)
	// 校验
	validate := validator.New()
	if err := validate.Struct(&role); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	detail, err := roleService.GetRoleDetail(role.Namespace, role.Role)
	if err != nil {
		response.FailWithMessage("获取失败", c)
		global.KOP_LOG.Error("获取失败", zap.Error(err))
	} else {
		response.OkWithDetailed(detail, "获取成功", c)
	}
}

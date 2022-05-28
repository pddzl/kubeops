package kubernetes

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"

	"github.com/pddzl/kubeops/server/global"
	"github.com/pddzl/kubeops/server/model/common/request"
	"github.com/pddzl/kubeops/server/model/common/response"
	k8sRequest "github.com/pddzl/kubeops/server/model/kubernetes/request"
)

type ClusterRoleApi struct{}

func (cr *ClusterRoleApi) GetClusterRoleList(c *gin.Context) {
	var pageInfo request.PageInfo
	_ = c.ShouldBindJSON(&pageInfo)
	// 校验字段
	validate := validator.New()
	if err := validate.Struct(&pageInfo); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	list, total, err := clusterRoleService.GetClusterRoleList(pageInfo)
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

func (cr *ClusterRoleApi) GetClusterRoleRaw(c *gin.Context) {
	var clusterRole k8sRequest.ClusterRoleCommon
	_ = c.ShouldBindJSON(&clusterRole)
	// 校验
	validate := validator.New()
	if err := validate.Struct(&clusterRole); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	raw, err := clusterRoleService.GetClusterRoleRaw(clusterRole.ClusterRole)
	if err != nil {
		response.FailWithMessage("获取失败", c)
		global.KOP_LOG.Error("获取失败", zap.Error(err))
	} else {
		response.OkWithDetailed(raw, "获取成功", c)
	}
}

func (cr *ClusterRoleApi) GetClusterRoleDetail(c *gin.Context) {
	var clusterRole k8sRequest.ClusterRoleCommon
	_ = c.ShouldBindJSON(&clusterRole)
	// 校验
	validate := validator.New()
	if err := validate.Struct(&clusterRole); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	detail, err := clusterRoleService.GetClusterRoleDetail(clusterRole.ClusterRole)
	if err != nil {
		response.FailWithMessage("获取失败", c)
		global.KOP_LOG.Error("获取失败", zap.Error(err))
	} else {
		response.OkWithDetailed(detail, "获取成功", c)
	}
}

// DeleteClusterRole 删除ClusterRole
func (cr *ClusterRoleApi) DeleteClusterRole(c *gin.Context) {
	var clusterRole k8sRequest.ClusterRoleCommon
	_ = c.ShouldBindJSON(&clusterRole)
	// 校验
	validate := validator.New()
	if err := validate.Struct(&clusterRole); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := clusterRoleService.DeleteClusterRole(clusterRole.ClusterRole); err != nil {
		response.FailWithMessage(err.Error(), c)
		global.KOP_LOG.Error("删除失败", zap.Error(err))
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

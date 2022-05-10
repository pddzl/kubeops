package kubernetes

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	k8sRequest "github.com/pddzl/kubeops/server/model/kubernetes/request"
	"go.uber.org/zap"

	"github.com/pddzl/kubeops/server/global"
	"github.com/pddzl/kubeops/server/model/common/request"
	"github.com/pddzl/kubeops/server/model/common/response"
)

type ClusterRoleBindingApi struct{}

func (cb *ClusterRoleBindingApi) GetClusterRoleBindingList(c *gin.Context) {
	var pageInfo request.PageInfo
	_ = c.ShouldBindJSON(&pageInfo)
	// 校验字段
	validate := validator.New()
	if err := validate.Struct(&pageInfo); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	list, total, err := ClusterRoleBindingService.GetClusterRoleBindingList(pageInfo)
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

func (cb *ClusterRoleBindingApi) GetClusterRoleBindingRaw(c *gin.Context) {
	var clusterRoleBinding k8sRequest.ClusterRoleBindingCommon
	_ = c.ShouldBindJSON(&clusterRoleBinding)
	// 校验
	validate := validator.New()
	if err := validate.Struct(&clusterRoleBinding); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	raw, err := ClusterRoleBindingService.GetClusterRoleBindingRaw(clusterRoleBinding.ClusterRoleBinding)
	if err != nil {
		response.FailWithMessage("获取失败", c)
		global.KOP_LOG.Error("获取失败", zap.Error(err))
	} else {
		response.OkWithDetailed(raw, "获取成功", c)
	}
}

func (cb *ClusterRoleBindingApi) GetClusterRoleBindingDetail(c *gin.Context) {
	var clusterRoleBinding k8sRequest.ClusterRoleBindingCommon
	_ = c.ShouldBindJSON(&clusterRoleBinding)
	// 校验
	validate := validator.New()
	if err := validate.Struct(&clusterRoleBinding); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	detail, err := ClusterRoleBindingService.GetClusterRoleBinding(clusterRoleBinding.ClusterRoleBinding)
	if err != nil {
		response.FailWithMessage("获取失败", c)
		global.KOP_LOG.Error("获取失败", zap.Error(err))
	} else {
		response.OkWithDetailed(detail, "获取成功", c)
	}
}

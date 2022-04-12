package kubernetes

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/pddzl/kubeops/server/global"
	"github.com/pddzl/kubeops/server/model/common/response"
	"github.com/pddzl/kubeops/server/model/kubernetes/request"
	"go.uber.org/zap"
)

type DeploymentApi struct{}

// 获取deployment list

func (d *DeploymentApi) GetDeploymentList(c *gin.Context) {
	var pageInfo request.ListParams
	_ = c.ShouldBindJSON(&pageInfo)
	// 校验字段
	validate := validator.New()
	if err := validate.Struct(&pageInfo.PageInfo); err != nil {
		global.KOP_LOG.Error("请求参数有误", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	list, total, err := deploymentService.GetDeploymentList(pageInfo.NameSpace, pageInfo.PageInfo)
	if err != nil {
		global.KOP_LOG.Error("获取失败", zap.Error(err))
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    int64(total),
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

// 获取deployment in 编排

func (d *DeploymentApi) GetDeploymentRaw(c *gin.Context) {
	var deploymentRaw request.DeploymentCommon
	_ = c.ShouldBindJSON(&deploymentRaw)
	// 校验
	validate := validator.New()
	if err := validate.Struct(&deploymentRaw); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	info, err := deploymentService.GetDeploymentRaw(deploymentRaw.NameSpace, deploymentRaw.Deployment)
	if err != nil {
		response.FailWithMessage("获取失败", c)
		global.KOP_LOG.Error("获取失败", zap.Error(err))
	}
	response.OkWithDetailed(info, "获取成功", c)
}

// 获取deployment detail

func (d *DeploymentApi) GetDeploymentDetail(c *gin.Context) {
	var deploymentRaw request.DeploymentCommon
	_ = c.ShouldBindJSON(&deploymentRaw)
	// 校验
	validate := validator.New()
	if err := validate.Struct(&deploymentRaw); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	detail, err := deploymentService.GetDeploymentDetail(deploymentRaw.NameSpace, deploymentRaw.Deployment)
	if err != nil {
		response.FailWithMessage("获取失败", c)
		global.KOP_LOG.Error("获取失败", zap.Error(err))
	}
	response.OkWithDetailed(detail, "获取成功", c)
}

// 获取deployment关联的replicaSet

func (d *DeploymentApi) GetNewReplicaSet(c *gin.Context) {
	var deploymentCommon request.DeploymentCommon
	_ = c.ShouldBindJSON(&deploymentCommon)
	// 校验
	validate := validator.New()
	if err := validate.Struct(&deploymentCommon); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	newReplicaSet, err := deploymentService.GetNewReplicaSet(deploymentCommon.NameSpace, deploymentCommon.Deployment)
	if err != nil {
		response.FailWithMessage("获取失败", c)
		global.KOP_LOG.Error("获取失败", zap.Error(err))
	}
	response.OkWithDetailed(newReplicaSet, "获取成功", c)
}

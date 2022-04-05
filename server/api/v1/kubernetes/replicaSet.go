package kubernetes

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"

	"github.com/pddzl/kubeops/server/global"
	"github.com/pddzl/kubeops/server/model/common/response"
	"github.com/pddzl/kubeops/server/model/kubernetes/request"
)

type ReplicaSetApi struct{}

func (r *ReplicaSetApi) GetReplicaSetList(c *gin.Context) {
	var pageInfo request.SearchPodParams
	_ = c.ShouldBindJSON(&pageInfo)
	// 校验字段
	validate := validator.New()
	if err := validate.Struct(&pageInfo.PageInfo); err != nil {
		global.KOP_LOG.Error("请求参数有误", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := replicaSetService.GetReplicaSetList(pageInfo.NameSpace, pageInfo.PageInfo)
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

// get pod detail

func (r *ReplicaSetApi) GetReplicaSetDetail(c *gin.Context) {
	var replicaSetInRaw request.ReplicaSetCommon
	_ = c.ShouldBindJSON(&replicaSetInRaw)
	// 校验
	validate := validator.New()
	if err := validate.Struct(&replicaSetInRaw); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	detail, err := replicaSetService.GetReplicaSetDetail(replicaSetInRaw.NameSpace, replicaSetInRaw.ReplicaSet)
	if err != nil {
		response.FailWithMessage("获取失败", c)
		global.KOP_LOG.Error("获取失败", zap.Error(err))
	}
	response.OkWithDetailed(detail, "获取成功", c)
}

// 获取replicaSet in raw

func (r *ReplicaSetApi) GetReplicaSetRaw(c *gin.Context) {
	var replicaSetInRaw request.ReplicaSetCommon
	_ = c.ShouldBindJSON(&replicaSetInRaw)
	// 校验
	validate := validator.New()
	if err := validate.Struct(&replicaSetInRaw); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	info, err := replicaSetService.GetReplicaSetRaw(replicaSetInRaw.NameSpace, replicaSetInRaw.ReplicaSet)
	if err != nil {
		response.FailWithMessage("获取失败", c)
		global.KOP_LOG.Error("获取失败", zap.Error(err))
	}
	response.OkWithDetailed(info, "获取成功", c)
}

// 获取replicaSet pods

func (r *ReplicaSetApi) GetReplicaSetPods(c *gin.Context) {
	var replicaSetPods request.ReplicaSetPods
	_ = c.ShouldBindJSON(&replicaSetPods)
	// 校验
	validate := validator.New()
	if err := validate.Struct(&replicaSetPods); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	pods, total, err := replicaSetService.GetReplicaSetPods(replicaSetPods.NameSpace, replicaSetPods.ReplicaSet, replicaSetPods.PageInfo)
	if err != nil {
		response.FailWithMessage("获取失败", c)
		global.KOP_LOG.Error("获取失败", zap.Error(err))
	}
	response.OkWithDetailed(response.PageResult{
		List:     pods,
		Total:    int64(total),
		Page:     replicaSetPods.Page,
		PageSize: replicaSetPods.PageSize,
	}, "获取成功", c)
}

// 获取replicaSet services

func (r *ReplicaSetApi) GetReplicaSetServices(c *gin.Context) {
	var replicaSetServices request.ReplicaSetCommon
	_ = c.ShouldBindJSON(&replicaSetServices)
	// 校验
	validate := validator.New()
	if err := validate.Struct(&replicaSetServices); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	services, err := replicaSetService.GetReplicaSetServices(replicaSetServices.NameSpace, replicaSetServices.ReplicaSet)
	if err != nil {
		response.FailWithMessage("获取失败", c)
		global.KOP_LOG.Error("获取失败", zap.Error(err))
	}
	response.OkWithDetailed(services, "获取成功", c)
}

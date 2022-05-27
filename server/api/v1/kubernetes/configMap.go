package kubernetes

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"

	"github.com/pddzl/kubeops/server/global"
	"github.com/pddzl/kubeops/server/model/common/response"
	"github.com/pddzl/kubeops/server/model/kubernetes/request"
)

type ConfigMapApi struct{}

func (cm *ConfigMapApi) GetConfigMapList(c *gin.Context) {
	var pageInfo request.ListParams
	_ = c.ShouldBindJSON(&pageInfo)
	// 校验字段
	validate := validator.New()
	if err := validate.Struct(&pageInfo.PageInfo); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	list, total, err := ConfigMapService.GetConfigMapList(pageInfo.NameSpace, pageInfo.PageInfo)
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

func (cm *ConfigMapApi) GetConfigMapRaw(c *gin.Context) {
	var configMap request.ConfigMapCommon
	_ = c.ShouldBindJSON(&configMap)
	// 校验
	validate := validator.New()
	if err := validate.Struct(&configMap); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	raw, err := ConfigMapService.GetConfigMapRaw(configMap.Namespace, configMap.ConfigMap)
	if err != nil {
		response.FailWithMessage("获取失败", c)
		global.KOP_LOG.Error("获取失败", zap.Error(err))
	} else {
		response.OkWithDetailed(raw, "获取成功", c)
	}
}

func (cm *ConfigMapApi) GetConfigMapDetail(c *gin.Context) {
	var configMap request.ConfigMapCommon
	_ = c.ShouldBindJSON(&configMap)
	// 校验
	validate := validator.New()
	if err := validate.Struct(&configMap); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	detail, err := ConfigMapService.GetConfigMapDetail(configMap.Namespace, configMap.ConfigMap)
	if err != nil {
		response.FailWithMessage("获取失败", c)
		global.KOP_LOG.Error("获取失败", zap.Error(err))
	} else {
		response.OkWithDetailed(detail, "获取成功", c)
	}
}

// DeleteConfigMap 删除configMap
func (cm *ConfigMapApi) DeleteConfigMap(c *gin.Context) {
	var configMap request.ConfigMapCommon
	_ = c.ShouldBindJSON(&configMap)
	// 校验
	validate := validator.New()
	if err := validate.Struct(&configMap); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := ConfigMapService.DeleteConfigMap(configMap.Namespace, configMap.ConfigMap); err != nil {
		response.FailWithMessage(err.Error(), c)
		global.KOP_LOG.Error("删除失败", zap.Error(err))
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

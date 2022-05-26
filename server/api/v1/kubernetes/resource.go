package kubernetes

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"

	"github.com/pddzl/kubeops/server/global"
	"github.com/pddzl/kubeops/server/model/common/response"
	"github.com/pddzl/kubeops/server/model/kubernetes/request"
)

type ResourceApi struct{}

func (r *ResourceApi) CreateDynamicResource(c *gin.Context) {
	var dynamicResource request.DynamicResource
	_ = c.ShouldBindJSON(&dynamicResource)
	// 校验字段
	validate := validator.New()
	if err := validate.Struct(&dynamicResource); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := resourceService.CreateDynamicResource(dynamicResource.Content); err != nil {
		global.KOP_LOG.Error("创建失败", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

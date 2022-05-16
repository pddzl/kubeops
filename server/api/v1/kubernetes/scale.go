package kubernetes

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"

	"github.com/pddzl/kubeops/server/global"
	"github.com/pddzl/kubeops/server/model/common/response"
	"github.com/pddzl/kubeops/server/model/kubernetes/request"
)

type ScaleApi struct{}

func (s *ScaleApi) Scale(c *gin.Context) {
	var common request.ScaleCommon
	_ = c.ShouldBindJSON(&common)
	// 校验
	validate := validator.New()
	if err := validate.Struct(&common); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	num, err := ScaleService.Scale(common.Namespace, common.Name, common.Kind, common.Num)
	if err != nil {
		global.KOP_LOG.Error("scale失败", zap.Error(err))
		response.OkWithDetailed(num, "scale失败", c)
	} else {
		response.OkWithDetailed(num, "scale成功", c)
	}
}

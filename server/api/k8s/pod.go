package k8s

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/pddzl/kubeops/server/global"
	"github.com/pddzl/kubeops/server/model/common/response"
	k8sRequest "github.com/pddzl/kubeops/server/model/k8s/request"
	"go.uber.org/zap"
)

type PodApi struct{}

// GetPods 获取pod list
func (pa *PodApi) GetPods(c *gin.Context) {
	var info k8sRequest.PodListReq
	_ = c.ShouldBindJSON(&info)
	// 校验
	validate := validator.New()
	if err := validate.Struct(&info); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if list, total, err := podService.GetPods(info.Namespace, info.Page, info.PageSize); err != nil {
		global.KOP_LOG.Error("获取失败", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    int64(total),
			Page:     info.Page,
			PageSize: info.PageSize},
			"获取成功", c)
	}
}

// GetPodDetail 获取pod detail
func (pa *PodApi) GetPodDetail(c *gin.Context) {
	//name := c.Query("name")
}
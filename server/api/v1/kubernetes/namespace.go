package kubernetes

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"

	"github.com/pddzl/kubeops/server/global"
	"github.com/pddzl/kubeops/server/model/common/request"
	"github.com/pddzl/kubeops/server/model/common/response"
)

type NamespaceApi struct{}

// GetNamespaceList 获取集群所有namespace
func (n *NamespaceApi) GetNamespaceList(c *gin.Context) {
	var pageInfo request.PageInfo
	_ = c.ShouldBindJSON(&pageInfo)
	// 校验
	validate := validator.New()
	if err := validate.Struct(&pageInfo); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	list, total, err := namespaceService.GetNamespaceList(pageInfo)
	if err != nil {
		global.KOP_LOG.Error("获取失败", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    int64(total),
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

// GetNamespaceOnlyName 获取集群所以namespace（只包括name）
func (n *NamespaceApi) GetNamespaceOnlyName(c *gin.Context) {
	list, err := namespaceService.GetNamespaceOnlyName()
	if err != nil {
		global.KOP_LOG.Error("获取失败", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(list, "获取成功", c)
}

// GetNamespaceDetail 获取namespace详情
func (n *NamespaceApi) GetNamespaceDetail(c *gin.Context) {
	var nameInfo request.GetByName
	_ = c.ShouldBindJSON(&nameInfo)
	// 校验
	validate := validator.New()
	if err := validate.Struct(&nameInfo); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	detail, err := namespaceService.GetNamespaceDetail(nameInfo.Name)
	if err != nil {
		global.KOP_LOG.Error("获取失败", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(detail, "获取成功", c)
}

// GetNamespaceRaw 获取namespace in raw
func (n *NamespaceApi) GetNamespaceRaw(c *gin.Context) {
	var nameInfo request.GetByName
	_ = c.ShouldBindJSON(&nameInfo)
	// 校验
	validate := validator.New()
	if err := validate.Struct(&nameInfo); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	raw, err := namespaceService.GetNamespaceRaw(nameInfo.Name)
	if err != nil {
		global.KOP_LOG.Error("获取失败", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(raw, "获取成功", c)
}

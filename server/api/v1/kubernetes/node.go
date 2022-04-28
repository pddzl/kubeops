package kubernetes

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"

	"github.com/pddzl/kubeops/server/global"
	commonRequest "github.com/pddzl/kubeops/server/model/common/request"
	"github.com/pddzl/kubeops/server/model/common/response"
	"github.com/pddzl/kubeops/server/model/kubernetes/request"
)

type NodeApi struct{}

// GetNodeList 获取所有节点
func (n *NodeApi) GetNodeList(c *gin.Context) {
	var pageInfo commonRequest.PageInfo
	_ = c.ShouldBindJSON(&pageInfo)
	// 校验
	validate := validator.New()
	if err := validate.Struct(&pageInfo); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if list, total, err := nodeService.GetNodeList(&pageInfo); err != nil {
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

// GetNodeDetail 获取节点详情
func (n *NodeApi) GetNodeDetail(c *gin.Context) {
	var nameInfo commonRequest.GetByName
	_ = c.ShouldBindJSON(&nameInfo)
	// 校验
	validate := validator.New()
	if err := validate.Struct(&nameInfo); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if nodeDetail, err := nodeService.GetNodeDetail(nameInfo.Name); err != nil {
		response.FailWithMessage("获取失败", c)
		global.KOP_LOG.Error("获取失败", zap.Error(err))
	} else {
		response.OkWithDetailed(nodeDetail, "获取成功", c)
	}
}

// GetNodeRaw 获取节点 in raw
func (n *NodeApi) GetNodeRaw(c *gin.Context) {
	var nameInfo commonRequest.GetByName
	_ = c.ShouldBindJSON(&nameInfo)
	// 校验
	validate := validator.New()
	if err := validate.Struct(&nameInfo); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if nodeDetail, err := nodeService.GetNodeRaw(nameInfo.Name); err != nil {
		response.FailWithMessage("获取失败", c)
		global.KOP_LOG.Error("获取失败", zap.Error(err))
	} else {
		response.OkWithDetailed(nodeDetail, "获取成功", c)
	}
}

// GetNodePods 获取pods
func (n *NodeApi) GetNodePods(c *gin.Context) {
	var node request.NodePods
	_ = c.ShouldBindJSON(&node)
	// 校验
	validate := validator.New()
	if err := validate.Struct(&node); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	pods, total, err := nodeService.GetNodePods(node.NodeName, node.PageInfo)
	if err != nil {
		response.FailWithMessage("获取失败", c)
		global.KOP_LOG.Error("获取失败", zap.Error(err))
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     pods,
			Total:    int64(total),
			Page:     node.Page,
			PageSize: node.PageSize,
		}, "获取成功", c)
	}
}

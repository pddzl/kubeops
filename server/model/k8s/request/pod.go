package request

import "github.com/pddzl/kubeops/server/model/common/request"

// PodListReq 获取pod列表
type PodListReq struct {
	request.PageInfo
	Namespace string `json:"namespace"`
}

// PodDetailReq 获取pod详情
type PodDetailReq struct {
	Namespace string `json:"namespace" validate:"required"`
	Name      string `json:"name" validate:"required"`
}

// PodLogReq 获取pod日志
type PodLogReq struct {
	Namespace string `json:"namespace" validate:"required"`
	Pod       string `json:"pod" validate:"required"`
	Container string `json:"container"`
	Lines     uint   `json:"lines"`
	Follow    bool   `json:"follow"`
}

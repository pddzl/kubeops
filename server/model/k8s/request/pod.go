package request

import "github.com/pddzl/kubeops/server/model/common/request"

// PodListReq 获取pod列表
type PodListReq struct {
	request.PageInfo
	Namespace string `json:"namespace"`
}

type PodReq struct {
	Namespace string `json:"namespace" validate:"required"`
	Pod       string `json:"pod" validate:"required"`
}

// PodLogReq 获取pod日志
type PodLogReq struct {
	PodReq
	Container string `json:"container"`
	Lines     uint   `json:"lines"`
	Follow    bool   `json:"follow"`
}

package request

import "github.com/pddzl/kubeops/server/model/common/request"

type PodListReq struct {
	request.PageInfo
	Namespace string `json:"namespace"`
}

type PodDetailReq struct {
	Namespace string `json:"namespace" validate:"required"`
	Name      string `json:"name" validate:"required"`
}

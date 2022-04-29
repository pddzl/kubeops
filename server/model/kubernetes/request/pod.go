package request

import (
	"github.com/pddzl/kubeops/server/model/common/request"
)

type PodLog struct {
	NameSpace string `json:"namespace" validate:"required"`
	Pod       string `json:"pod" validate:"required"`
	Container string `json:"container"`
	Lines     int64  `json:"lines"`
}

type SearchPodParams struct {
	NameSpace string `json:"namespace"`
	request.PageInfo
}

type PodList struct {
	Namespace string `json:"namespace"`
	request.PageInfo
}

type PodDetail struct {
	Pod       string `json:"pod" validate:"required"`
	NameSpace string `json:"namespace" validate:"required"`
}

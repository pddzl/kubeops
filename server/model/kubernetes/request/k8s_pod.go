package request

import (
	"github.com/pddzl/kubeops/server/model/common/request"
)

type PodLog struct {
	NameSpace string `json:"namespace" default:"default"`
	Pod       string `json:"pod" validate:"required"`
	Lines     int64  `json:"lines" default:"20"`
}

type SearchPodParams struct {
	NameSpace string `json:"namespace"`
	request.PageInfo
}

package request

import "github.com/pddzl/kubeops/server/model/common/request"

type JobCommon struct {
	NameSpace string `json:"namespace" validate:"required"`
	Job       string `json:"job" validate:"required"`
}

type JobPods struct {
	JobCommon
	request.PageInfo
}

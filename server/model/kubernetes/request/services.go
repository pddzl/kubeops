package request

import "github.com/pddzl/kubeops/server/model/common/request"

type ServicesCommon struct {
	NameSpace string `json:"namespace" validate:"required"`
	Service   string `json:"service" validate:"required"`
}

type ServicesPods struct {
	ServicesCommon
	request.PageInfo
}

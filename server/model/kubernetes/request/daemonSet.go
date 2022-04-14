package request

import "github.com/pddzl/kubeops/server/model/common/request"

type DaemonSetCommon struct {
	NameSpace string `json:"namespace" validate:"required"`
	DaemonSet string `json:"daemonSet" validate:"required"`
}

type DaemonSetPods struct {
	DaemonSetCommon
	request.PageInfo
}

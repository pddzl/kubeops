package request

import "github.com/pddzl/kubeops/server/model/common/request"

type ReplicaSetCommon struct {
	NameSpace  string `json:"namespace" validate:"required"`
	ReplicaSet string `json:"replicaSet" validate:"required"`
}

type ReplicaSetPods struct {
	ReplicaSetCommon
	request.PageInfo
}

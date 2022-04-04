package request

import "github.com/pddzl/kubeops/server/model/common/request"

type ReplicaSetInRaw struct {
	NameSpace  string `json:"namespace" validate:"required"`
	ReplicaSet string `json:"replicaSet" validate:"required"`
}

type ReplicaSetPods struct {
	NameSpace  string `json:"namespace" validate:"required"`
	ReplicaSet string `json:"replicaSet" validate:"required"`
	request.PageInfo
}

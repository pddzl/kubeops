package request

import "github.com/pddzl/kubeops/server/model/common/request"

type NodePods struct {
	NodeName string `json:"node_name"`
	request.PageInfo
}

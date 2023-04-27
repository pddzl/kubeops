package request

import "github.com/pddzl/kubeops/server/model/common/request"

type PodListReq struct {
	request.PageInfo
	Namespace string `json:"namespace"`
}

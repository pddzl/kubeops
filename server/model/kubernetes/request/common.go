package request

import "github.com/pddzl/kubeops/server/model/common/request"

type ListParams struct {
	NameSpace string `json:"namespace"`
	request.PageInfo
}

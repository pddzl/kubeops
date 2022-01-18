package response

import (
	"github.com/pddzl/kubeops/server/model/system/request"
)

type PolicyPathResponse struct {
	Paths []request.CasbinInfo `json:"paths"`
}

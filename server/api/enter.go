package api

import (
	"github.com/pddzl/kubeops/server/api/system"
)

type ApiGroup struct {
	SystemApiGroup system.ApiGroup
}

var ApiGroupApp = new(ApiGroup)

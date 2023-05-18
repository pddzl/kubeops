package api

import (
	"github.com/pddzl/kubeops/server/api/k8s"
	"github.com/pddzl/kubeops/server/api/system"
)

type ApiGroup struct {
	SystemApiGroup system.ApiGroup
	K8sApiGroup    k8s.ApiGroup
}

var ApiGroupApp = new(ApiGroup)

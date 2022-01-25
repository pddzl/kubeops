package v1

import (
	"github.com/pddzl/kubeops/server/api/v1/example"
	"github.com/pddzl/kubeops/server/api/v1/kubernetes"
	"github.com/pddzl/kubeops/server/api/v1/system"
)

type ApiGroup struct {
	SystemApiGroup     system.ApiGroup
	ExampleApiGroup    example.ApiGroup
	KubernetesApiGroup kubernetes.ApiGroup
}

var ApiGroupApp = new(ApiGroup)

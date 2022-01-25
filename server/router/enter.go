package router

import (
	"github.com/pddzl/kubeops/server/router/example"
	"github.com/pddzl/kubeops/server/router/kubernetes"
	"github.com/pddzl/kubeops/server/router/system"
)

type RouterGroup struct {
	System     system.RouterGroup
	Example    example.RouterGroup
	Kubernetes kubernetes.RouterGroup
}

var RouterGroupApp = new(RouterGroup)

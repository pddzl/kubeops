package router

import (
	"github.com/pddzl/kubeops/server/router/system"
)

type RouterGroup struct {
	System system.RouterGroup
}

var RouterGroupApp = new(RouterGroup)

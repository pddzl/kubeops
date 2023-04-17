package system

import (
	"github.com/pddzl/kubeops/server/global"
)

type JwtBlacklist struct {
	global.TD27_MODEL
	Jwt string `gorm:"type:text;comment:jwt"`
}

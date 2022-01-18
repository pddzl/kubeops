package system

import (
	"github.com/pddzl/kubeops/server/global"
)

type SysApi struct {
	global.KOP_MODEL
	Path        string `json:"path" gorm:"comment:api路径" validate:"required"`             // api路径
	Description string `json:"description" gorm:"comment:api中文描述" validate:"required"`    // api中文描述
	ApiGroup    string `json:"apiGroup" gorm:"comment:api组" validate:"required"`          // api组
	Method      string `json:"method" gorm:"default:POST;comment:方法" validate:"required"` // 方法:创建POST(默认)|查看GET|更新PUT|删除DELETE
}

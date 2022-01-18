package system

import (
	"github.com/pddzl/kubeops/server/global"
	"github.com/pddzl/kubeops/server/model/system"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

var Authority = new(authority)

type authority struct{}

func (a *authority) TableName() string {
	return "sys_authorities"
}

func (a *authority) Initialize() error {
	entities := []system.SysAuthority{
		{AuthorityId: "888", AuthorityName: "管理员", DefaultRouter: "dashboard"},
		{AuthorityId: "9528", AuthorityName: "测试账号", DefaultRouter: "dashboard"},
	}
	if err := global.KOP_DB.Create(&entities).Error; err != nil {
		return errors.Wrapf(err, "%s表数据初始化失败!", a.TableName())
	}
	return nil
}

func (a *authority) CheckDataExist() bool {
	if errors.Is(global.KOP_DB.Where("authority_id = ?", "888").First(&system.SysAuthority{}).Error, gorm.ErrRecordNotFound) { // 判断是否存在数据
		return false
	}
	return true
}

package system

import (
	adapter "github.com/casbin/gorm-adapter/v3"
	"github.com/pddzl/kubeops/server/global"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

var Casbin = new(casbin)

type casbin struct{}

func (c *casbin) TableName() string {
	var entity adapter.CasbinRule
	return entity.TableName()
}

func (c *casbin) Initialize() error {
	entities := []adapter.CasbinRule{
		{PType: "p", V0: "888", V1: "/base/login", V2: "POST"},
		{PType: "p", V0: "888", V1: "/user/admin_register", V2: "POST"},

		{PType: "p", V0: "888", V1: "/api/createApi", V2: "POST"},
		{PType: "p", V0: "888", V1: "/api/getApiList", V2: "POST"},
		{PType: "p", V0: "888", V1: "/api/getApiById", V2: "POST"},
		{PType: "p", V0: "888", V1: "/api/deleteApi", V2: "POST"},
		{PType: "p", V0: "888", V1: "/api/updateApi", V2: "POST"},
		{PType: "p", V0: "888", V1: "/api/getAllApis", V2: "POST"},
		{PType: "p", V0: "888", V1: "/api/deleteApisByIds", V2: "DELETE"},

		{PType: "p", V0: "888", V1: "/authority/copyAuthority", V2: "POST"},
		{PType: "p", V0: "888", V1: "/authority/updateAuthority", V2: "PUT"},
		{PType: "p", V0: "888", V1: "/authority/createAuthority", V2: "POST"},
		{PType: "p", V0: "888", V1: "/authority/deleteAuthority", V2: "POST"},
		{PType: "p", V0: "888", V1: "/authority/getAuthorityList", V2: "POST"},
		{PType: "p", V0: "888", V1: "/authority/setDataAuthority", V2: "POST"},

		{PType: "p", V0: "888", V1: "/menu/getMenu", V2: "POST"},
		{PType: "p", V0: "888", V1: "/menu/getMenuList", V2: "POST"},
		{PType: "p", V0: "888", V1: "/menu/addBaseMenu", V2: "POST"},
		{PType: "p", V0: "888", V1: "/menu/getBaseMenuTree", V2: "POST"},
		{PType: "p", V0: "888", V1: "/menu/addMenuAuthority", V2: "POST"},
		{PType: "p", V0: "888", V1: "/menu/getMenuAuthority", V2: "POST"},
		{PType: "p", V0: "888", V1: "/menu/deleteBaseMenu", V2: "POST"},
		{PType: "p", V0: "888", V1: "/menu/updateBaseMenu", V2: "POST"},
		{PType: "p", V0: "888", V1: "/menu/getBaseMenuById", V2: "POST"},

		{PType: "p", V0: "888", V1: "/user/getUserInfo", V2: "GET"},
		{PType: "p", V0: "888", V1: "/user/setUserInfo", V2: "PUT"},
		{PType: "p", V0: "888", V1: "/user/setSelfInfo", V2: "PUT"},
		{PType: "p", V0: "888", V1: "/user/getUserList", V2: "POST"},
		{PType: "p", V0: "888", V1: "/user/deleteUser", V2: "DELETE"},
		{PType: "p", V0: "888", V1: "/user/changePassword", V2: "POST"},
		{PType: "p", V0: "888", V1: "/user/setUserAuthority", V2: "POST"},
		{PType: "p", V0: "888", V1: "/user/setUserAuthorities", V2: "POST"},
		{PType: "p", V0: "888", V1: "/user/resetPassword", V2: "POST"},

		{PType: "p", V0: "888", V1: "/fileUploadAndDownload/upload", V2: "POST"},
		{PType: "p", V0: "888", V1: "/fileUploadAndDownload/deleteFile", V2: "POST"},
		{PType: "p", V0: "888", V1: "/fileUploadAndDownload/getFileList", V2: "POST"},

		{PType: "p", V0: "888", V1: "/casbin/updateCasbin", V2: "POST"},
		{PType: "p", V0: "888", V1: "/casbin/getPolicyPathByAuthorityId", V2: "POST"},

		{PType: "p", V0: "888", V1: "/jwt/jsonInBlacklist", V2: "POST"},

		{PType: "p", V0: "888", V1: "/sysOperationRecord/findSysOperationRecord", V2: "GET"},
		{PType: "p", V0: "888", V1: "/sysOperationRecord/updateSysOperationRecord", V2: "PUT"},
		{PType: "p", V0: "888", V1: "/sysOperationRecord/createSysOperationRecord", V2: "POST"},
		{PType: "p", V0: "888", V1: "/sysOperationRecord/getSysOperationRecordList", V2: "GET"},
		{PType: "p", V0: "888", V1: "/sysOperationRecord/deleteSysOperationRecord", V2: "DELETE"},
		{PType: "p", V0: "888", V1: "/sysOperationRecord/deleteSysOperationRecordByIds", V2: "DELETE"},

		{PType: "p", V0: "9528", V1: "/base/login", V2: "POST"},
		{PType: "p", V0: "9528", V1: "/user/admin_register", V2: "POST"},
		{PType: "p", V0: "9528", V1: "/api/createApi", V2: "POST"},
		{PType: "p", V0: "9528", V1: "/api/getApiList", V2: "POST"},
		{PType: "p", V0: "9528", V1: "/api/getApiById", V2: "POST"},
		{PType: "p", V0: "9528", V1: "/api/deleteApi", V2: "POST"},
		{PType: "p", V0: "9528", V1: "/api/updateApi", V2: "POST"},
		{PType: "p", V0: "9528", V1: "/api/getAllApis", V2: "POST"},

		{PType: "p", V0: "9528", V1: "/authority/createAuthority", V2: "POST"},
		{PType: "p", V0: "9528", V1: "/authority/deleteAuthority", V2: "POST"},
		{PType: "p", V0: "9528", V1: "/authority/getAuthorityList", V2: "POST"},
		{PType: "p", V0: "9528", V1: "/authority/setDataAuthority", V2: "POST"},

		{PType: "p", V0: "9528", V1: "/menu/getMenu", V2: "POST"},
		{PType: "p", V0: "9528", V1: "/menu/getMenuList", V2: "POST"},
		{PType: "p", V0: "9528", V1: "/menu/addBaseMenu", V2: "POST"},
		{PType: "p", V0: "9528", V1: "/menu/getBaseMenuTree", V2: "POST"},
		{PType: "p", V0: "9528", V1: "/menu/addMenuAuthority", V2: "POST"},
		{PType: "p", V0: "9528", V1: "/menu/getMenuAuthority", V2: "POST"},
		{PType: "p", V0: "9528", V1: "/menu/deleteBaseMenu", V2: "POST"},
		{PType: "p", V0: "9528", V1: "/menu/updateBaseMenu", V2: "POST"},
		{PType: "p", V0: "9528", V1: "/menu/getBaseMenuById", V2: "POST"},
		{PType: "p", V0: "9528", V1: "/user/changePassword", V2: "POST"},
		{PType: "p", V0: "9528", V1: "/user/getUserList", V2: "POST"},
		{PType: "p", V0: "9528", V1: "/user/setUserAuthority", V2: "POST"},
		{PType: "p", V0: "9528", V1: "/casbin/updateCasbin", V2: "POST"},
		{PType: "p", V0: "9528", V1: "/casbin/getPolicyPathByAuthorityId", V2: "POST"},
		{PType: "p", V0: "9528", V1: "/jwt/jsonInBlacklist", V2: "POST"},
		{PType: "p", V0: "9528", V1: "/user/getUserInfo", V2: "GET"},
	}
	if err := global.KOP_DB.Create(&entities).Error; err != nil {
		return errors.Wrap(err, c.TableName()+"????????????????????????!")
	}
	return nil
}

func (c *casbin) CheckDataExist() bool {
	if errors.Is(global.KOP_DB.Where(adapter.CasbinRule{PType: "p", V0: "9528", V1: "GET", V2: "/user/getUserInfo"}).First(&adapter.CasbinRule{}).Error, gorm.ErrRecordNotFound) { // ????????????????????????
		return false
	}
	return true
}

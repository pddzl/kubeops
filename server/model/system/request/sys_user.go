package request

import model "github.com/pddzl/kubeops/server/model/system"

// User register structure

type Register struct {
	Username     string   `json:"userName" validate:"required"`
	Password     string   `json:"passWord" validate:"required"`
	NickName     string   `json:"nickName" gorm:"default:'QMPlusUser'" validate:"required"`
	HeaderImg    string   `json:"headerImg" gorm:"default:'https://qmplusimg.henrongyi.top/gva_header.jpg'"`
	AuthorityId  string   `json:"authorityId" gorm:"default:888" validate:"required"`
	AuthorityIds []string `json:"authorityIds"`
}

// User login structure

type Login struct {
	Username  string `json:"username" validate:"required"`  // 用户名
	Password  string `json:"password" validate:"required"`  // 密码
	Captcha   string `json:"captcha" validate:"required"`   // 验证码
	CaptchaId string `json:"captchaId" validate:"required"` // 验证码ID
}

// Modify password structure

type ChangePasswordStruct struct {
	Username    string `json:"username" validate:"required"`    // 用户名
	Password    string `json:"password" validate:"required"`    // 密码
	NewPassword string `json:"newPassword" validate:"required"` // 新密码
}

// Modify  user's auth structure

type SetUserAuth struct {
	AuthorityId string `json:"authorityId" validate:"required"` // 角色ID
}

// Modify  user's auth structure

type SetUserAuthorities struct {
	ID           uint
	AuthorityIds []string `json:"authorityIds"` // 角色ID
}

type ChangeUserInfo struct {
	ID           uint                 `gorm:"primarykey"`                                                                           // 主键ID
	NickName     string               `json:"nickName" gorm:"default:系统用户;comment:用户昵称"`                                            // 用户昵称
	Phone        string               `json:"phone"  gorm:"comment:用户手机号"`                                                          // 用户角色ID
	AuthorityIds []string             `json:"authorityIds" gorm:"-"`                                                                // 角色ID
	Email        string               `json:"email"  gorm:"comment:用户邮箱"`                                                           // 用户邮箱
	HeaderImg    string               `json:"headerImg" gorm:"default:https://qmplusimg.henrongyi.top/gva_header.jpg;comment:用户头像"` // 用户头像
	Authorities  []model.SysAuthority `json:"-" gorm:"many2many:sys_user_authority;"`
}

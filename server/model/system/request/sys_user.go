package request

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

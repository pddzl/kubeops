package system

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"

	"github.com/pddzl/kubeops/server/global"
	"github.com/pddzl/kubeops/server/model/common/request"
	"github.com/pddzl/kubeops/server/model/common/response"
	"github.com/pddzl/kubeops/server/model/system"
	systemReq "github.com/pddzl/kubeops/server/model/system/request"
	systemRes "github.com/pddzl/kubeops/server/model/system/response"
)

type AuthorityApi struct{}

// @Tags Authority
// @Summary 创建角色
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysAuthority true "权限id, 权限名, 父角色id"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /authority/createAuthority [post]

func (a *AuthorityApi) CreateAuthority(c *gin.Context) {
	var authority system.SysAuthority
	_ = c.ShouldBindJSON(&authority)
	// 校验
	validate := validator.New()
	if err := validate.Struct(&authority); err != nil {
		global.KOP_LOG.Error("请求参数有误", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err, authBack := authorityService.CreateAuthority(authority); err != nil {
		global.KOP_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败"+err.Error(), c)
	} else {
		_ = menuService.AddMenuAuthority(systemReq.DefaultMenu(), authority.AuthorityId)
		_ = casbinService.UpdateCasbin(authority.AuthorityId, systemReq.DefaultCasbin())
		response.OkWithDetailed(systemRes.SysAuthorityResponse{Authority: authBack}, "创建成功", c)
	}
}

// @Tags Authority
// @Summary 拷贝角色
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body response.SysAuthorityCopyResponse true "旧角色id, 新权限id, 新权限名, 新父角色id"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"拷贝成功"}"
// @Router /authority/copyAuthority [post]

func (a *AuthorityApi) CopyAuthority(c *gin.Context) {
	var copyInfo systemRes.SysAuthorityCopyResponse
	_ = c.ShouldBindJSON(&copyInfo)
	// 校验
	validate := validator.New()
	if err := validate.Struct(&copyInfo); err != nil {
		global.KOP_LOG.Error("请求参数有误", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err, authBack := authorityService.CopyAuthority(copyInfo); err != nil {
		global.KOP_LOG.Error("拷贝失败!", zap.Error(err))
		response.FailWithMessage("拷贝失败"+err.Error(), c)
	} else {
		response.OkWithDetailed(systemRes.SysAuthorityResponse{Authority: authBack}, "拷贝成功", c)
	}
}

// @Tags Authority
// @Summary 删除角色
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysAuthority true "删除角色"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /authority/deleteAuthority [post]

func (a *AuthorityApi) DeleteAuthority(c *gin.Context) {
	var authority system.SysAuthority
	_ = c.ShouldBindJSON(&authority)
	// 校验
	if authority.AuthorityId == "" {
		response.FailWithMessage("authorityId为空", c)
		global.KOP_LOG.Error("请求参数有误", zap.String("authorityId", "authorityId为空"))
		return
	}
	if err := authorityService.DeleteAuthority(&authority); err != nil { // 删除角色之前需要判断是否有用户正在使用此角色
		global.KOP_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败"+err.Error(), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Authority
// @Summary 更新角色信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysAuthority true "权限id, 权限名, 父角色id"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /authority/updateAuthority [post]

func (a *AuthorityApi) UpdateAuthority(c *gin.Context) {
	var auth system.SysAuthority
	_ = c.ShouldBindJSON(&auth)
	// 校验
	validate := validator.New()
	if err := validate.Struct(&auth); err != nil {
		global.KOP_LOG.Error("请求参数有误", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err, authority := authorityService.UpdateAuthority(auth); err != nil {
		global.KOP_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败"+err.Error(), c)
	} else {
		response.OkWithDetailed(systemRes.SysAuthorityResponse{Authority: authority}, "更新成功", c)
	}
}

// @Tags Authority
// @Summary 分页获取角色列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "页码, 每页大小"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /authority/getAuthorityList [post]

func (a *AuthorityApi) GetAuthorityList(c *gin.Context) {
	var pageInfo request.PageInfo
	_ = c.ShouldBindJSON(&pageInfo)
	// 校验
	validate := validator.New()
	if err := validate.Struct(&pageInfo); err != nil {
		global.KOP_LOG.Error("请求参数有误", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err, list, total := authorityService.GetAuthorityInfoList(pageInfo); err != nil {
		global.KOP_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败"+err.Error(), c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

// @Tags Authority
// @Summary 设置角色资源权限
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysAuthority true "设置角色资源权限"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"设置成功"}"
// @Router /authority/setDataAuthority [post]

func (a *AuthorityApi) SetDataAuthority(c *gin.Context) {
	var auth system.SysAuthority
	_ = c.ShouldBindJSON(&auth)
	// 校验
	if auth.AuthorityId == "" {
		response.FailWithMessage("authorityId为空", c)
		global.KOP_LOG.Error("请求参数有误", zap.String("authorityId", "authorityId为空"))
		return
	}
	if err := authorityService.SetDataAuthority(auth); err != nil {
		global.KOP_LOG.Error("设置失败!", zap.Error(err))
		response.FailWithMessage("设置失败"+err.Error(), c)
	} else {
		response.OkWithMessage("设置成功", c)
	}
}

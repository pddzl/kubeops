package system

import (
	"github.com/gin-gonic/gin"
	"github.com/pddzl/kubeops/server/global"
	"github.com/pddzl/kubeops/server/model/common/response"
	"github.com/pddzl/kubeops/server/model/system/request"
	systemRes "github.com/pddzl/kubeops/server/model/system/response"
	"go.uber.org/zap"
)

type CasbinApi struct{}

// @Tags Casbin
// @Summary 更新角色api权限
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.CasbinInReceive true "权限id, 权限模型列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /casbin/UpdateCasbin [post]

func (cas *CasbinApi) UpdateCasbin(c *gin.Context) {
	var cmr request.CasbinInReceive
	_ = c.ShouldBindJSON(&cmr)
	// 校验
	if cmr.AuthorityId == "" {
		response.FailWithMessage("authorityId为空", c)
		global.KOP_LOG.Error("请求参数有误", zap.String("authorityId", "authorityId为空"))
		return
	}
	if err := casbinService.UpdateCasbin(cmr.AuthorityId, cmr.CasbinInfos); err != nil {
		global.KOP_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags Casbin
// @Summary 获取权限列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.CasbinInReceive true "权限id, 权限模型列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /casbin/getPolicyPathByAuthorityId [post]

func (cas *CasbinApi) GetPolicyPathByAuthorityId(c *gin.Context) {
	var casbin request.CasbinInReceive
	_ = c.ShouldBindJSON(&casbin)
	// 校验
	if casbin.AuthorityId == "" {
		response.FailWithMessage("authorityId为空", c)
		global.KOP_LOG.Error("请求参数有误", zap.String("authorityId", "authorityId为空"))
		return
	}
	paths := casbinService.GetPolicyPathByAuthorityId(casbin.AuthorityId)
	response.OkWithDetailed(systemRes.PolicyPathResponse{Paths: paths}, "获取成功", c)
}

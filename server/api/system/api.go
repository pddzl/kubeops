package system

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/pddzl/kubeops/server/global"
	"github.com/pddzl/kubeops/server/model/common/request"
	"github.com/pddzl/kubeops/server/model/common/response"
	systemModel "github.com/pddzl/kubeops/server/model/system"
	systemReq "github.com/pddzl/kubeops/server/model/system/request"
	systemRep "github.com/pddzl/kubeops/server/model/system/response"
	"go.uber.org/zap"
)

type ApiApi struct{}

// AddApi 添加api
func (a *ApiApi) AddApi(c *gin.Context) {
	var apiReq systemModel.ApiModel
	_ = c.ShouldBindJSON(&apiReq)

	// 参数校验
	validate := validator.New()
	if err := validate.Struct(&apiReq); err != nil {
		response.FailWithMessage("请求参数错误", c)
		global.KOP_LOG.Error("请求参数错误", zap.Error(err))
		return
	}

	if instance, err := apiService.AddApi(apiReq); err != nil {
		response.FailWithMessage("添加失败", c)
		global.KOP_LOG.Error("添加失败", zap.Error(err))
	} else {
		response.OkWithDetailed(instance, "添加成功", c)
	}
}

// GetApis 列出所有api
func (a *ApiApi) GetApis(c *gin.Context) {
	var apiSp systemReq.ApiSearchParams
	_ = c.ShouldBindJSON(&apiSp)

	// 参数校验
	validate := validator.New()
	if err := validate.Struct(&apiSp); err != nil {
		response.FailWithMessage("请求参数错误", c)
		global.KOP_LOG.Error("请求参数错误", zap.Error(err))
		return
	}

	if list, total, err := apiService.GetApis(apiSp); err != nil {
		response.FailWithMessage("获取失败", c)
		global.KOP_LOG.Error("获取失败", zap.Error(err))
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     apiSp.Page,
			PageSize: apiSp.PageSize,
		}, "获取成功", c)
	}
}

// DeleteApi 删除指定api
func (a *ApiApi) DeleteApi(c *gin.Context) {
	var cId request.CId
	_ = c.ShouldBindJSON(&cId)

	// 参数校验
	validate := validator.New()
	if err := validate.Struct(&cId); err != nil {
		response.FailWithMessage("请求参数错误", c)
		global.KOP_LOG.Error("请求参数错误", zap.Error(err))
		return
	}

	if err := apiService.DeleteApi(cId.ID); err != nil {
		response.FailWithMessage("删除失败", c)
		global.KOP_LOG.Error("删除失败", zap.Error(err))
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// EditApi 编辑api
func (a *ApiApi) EditApi(c *gin.Context) {
	var eApi systemReq.EditApi
	_ = c.ShouldBindJSON(&eApi)

	// 参数校验
	validate := validator.New()
	if err := validate.Struct(&eApi); err != nil {
		response.FailWithMessage("请求参数错误", c)
		global.KOP_LOG.Error("请求参数错误", zap.Error(err))
		return
	}

	if err := apiService.EditApi(eApi); err != nil {
		response.FailWithMessage("编辑失败", c)
		global.KOP_LOG.Error("编辑失败", zap.Error(err))
	} else {
		response.OkWithMessage("编辑成功", c)
	}
}

// GetElTreeApis 格式化列出所有api
func (a *ApiApi) GetElTreeApis(c *gin.Context) {
	var cId request.CId
	_ = c.ShouldBindJSON(&cId)

	// 参数校验
	validate := validator.New()
	if err := validate.Struct(&cId); err != nil {
		response.FailWithMessage("请求参数错误", c)
		global.KOP_LOG.Error("请求参数错误", zap.Error(err))
		return
	}

	list, checkedKey, err := apiService.GetElTreeApis(cId.ID)
	if err != nil {
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(systemRep.ApiTree{
			List:       list,
			CheckedKey: checkedKey,
		}, "获取成功", c)
	}
}

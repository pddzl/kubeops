package kubernetes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
	"net/http"

	"github.com/pddzl/kubeops/server/global"
	"github.com/pddzl/kubeops/server/model/common/response"
	"github.com/pddzl/kubeops/server/model/kubernetes/request"
)

type PodApi struct{}

// GetPodList 获取集群所有pod
func (p *PodApi) GetPodList(c *gin.Context) {
	var podList request.PodList
	_ = c.ShouldBindJSON(&podList)
	// 校验字段
	validate := validator.New()
	if err := validate.Struct(&podList.PageInfo); err != nil {
		global.KOP_LOG.Error("请求参数有误", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	list, total, err := podService.GetPodList(podList.Namespace, podList.PageInfo)
	if err != nil {
		response.FailWithMessage("获取失败", c)
		global.KOP_LOG.Error("获取失败", zap.Error(err))
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    int64(total),
			Page:     podList.Page,
			PageSize: podList.PageSize,
		}, "获取成功", c)
	}
}

// GetPodDetail 获取pod详情
func (p *PodApi) GetPodDetail(c *gin.Context) {
	var podDetail request.PodDetail
	_ = c.ShouldBindJSON(&podDetail)
	// 校验
	validate := validator.New()
	if err := validate.Struct(&podDetail); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	info, err := podService.GetPodDetail(podDetail.NameSpace, podDetail.Pod)
	if err != nil {
		global.KOP_LOG.Error("获取失败", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	}
	response.OkWithDetailed(info, "获取成功", c)
}

// GetPodRaw 获取pod in raw
func (p *PodApi) GetPodRaw(c *gin.Context) {
	var podDetail request.PodDetail
	_ = c.ShouldBindJSON(&podDetail)
	// 校验
	validate := validator.New()
	if err := validate.Struct(&podDetail); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	info, err := podService.GetPodRaw(podDetail.NameSpace, podDetail.Pod)
	if err != nil {
		global.KOP_LOG.Error("获取失败", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	}
	response.OkWithDetailed(info, "获取成功", c)
}

// GetPodLog 获取pod日志
func (p *PodApi) GetPodLog(c *gin.Context) {
	var podLog request.PodLog
	_ = c.ShouldBindJSON(&podLog)
	// 字段校验
	validate := validator.New()
	if err := validate.Struct(&podLog); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	log, err := podService.GetPodLog(podLog.NameSpace, podLog.Container, podLog.Pod, podLog.Lines)
	if err != nil {
		global.KOP_LOG.Error(err.Error())
		response.FailWithMessage("获取日志失败", c)
		return
	}
	response.OkWithDetailed(log, "获取日志成功", c)
}

// DeletePod 删除pod
func (p *PodApi) DeletePod(c *gin.Context) {
	var podDetail request.PodDetail
	_ = c.ShouldBindJSON(&podDetail)
	// 校验
	validate := validator.New()
	if err := validate.Struct(&podDetail); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := podService.DeletePod(podDetail.NameSpace, podDetail.Pod); err != nil {
		response.FailWithMessage(err.Error(), c)
		global.KOP_LOG.Error("删除失败", zap.Error(err))
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// 获取pod webShell
var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (p *PodApi) GetPodTerminal(c *gin.Context) {
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		global.KOP_LOG.Error(fmt.Sprintf("初始化websocket失败 %s", c.Errors))
	}
	defer ws.Close()

	session, err := podService.NewTerminalSession(ws)
	if err != nil {
		global.KOP_LOG.Error(fmt.Sprintf("初始化session失败 %s", c.Errors))
	}

	namespace, _ := c.GetQuery("namespace")
	pod, _ := c.GetQuery("pod")
	container, _ := c.GetQuery("container")

	err = podService.GetPodTerminal(namespace, pod, container, global.KUBERNETES_CONFIG, session)
	if err != nil {
		global.KOP_LOG.Error("terminal失败", zap.Error(err))
	}
}

// GetPodStatus 获取pod状态
func (p *PodApi) GetPodStatus(c *gin.Context) {
	chart, err := podService.GetPodStatus()
	if err != nil {
		response.FailWithMessage("获取失败", c)
		global.KOP_LOG.Error("获取失败", zap.Error(err))
	} else {
		response.OkWithDetailed(chart, "获取成功", c)
	}
}

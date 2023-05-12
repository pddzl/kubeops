package k8s

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/websocket"
	"github.com/pddzl/kubeops/server/global"
	"github.com/pddzl/kubeops/server/model/common/response"
	k8sRequest "github.com/pddzl/kubeops/server/model/k8s/request"
	"go.uber.org/zap"
	"net/http"
)

type PodApi struct{}

// GetPods 获取pod list
func (pa *PodApi) GetPods(c *gin.Context) {
	var info k8sRequest.PodListReq
	_ = c.ShouldBindJSON(&info)
	// 校验
	validate := validator.New()
	if err := validate.Struct(&info); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if list, total, err := podService.GetPods(info.Namespace, info.Page, info.PageSize); err != nil {
		response.FailWithMessage("获取失败", c)
		global.KOP_LOG.Error("获取失败", zap.Error(err))
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    int64(total),
			Page:     info.Page,
			PageSize: info.PageSize},
			"获取成功", c)
	}
}

// GetPodDetail 获取pod detail
func (pa *PodApi) GetPodDetail(c *gin.Context) {
	var pdReq k8sRequest.PodReq
	_ = c.ShouldBindJSON(&pdReq)
	// 校验
	validate := validator.New()
	if err := validate.Struct(&pdReq); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if detail, err := podService.GetPodDetail(pdReq.Namespace, pdReq.Pod); err != nil {
		response.FailWithMessage("获取失败", c)
		global.KOP_LOG.Error("获取失败", zap.Error(err))
	} else {
		response.OkWithDetailed(detail, "获取成功", c)
	}
}

// GetPodLog 获取pod日志
func (pa *PodApi) GetPodLog(c *gin.Context) {
	var plReq k8sRequest.PodLogReq
	_ = c.ShouldBindJSON(&plReq)

	// 校验
	validate := validator.New()
	if err := validate.Struct(&plReq); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	log, err := podService.GetPodLog(plReq.Namespace, plReq.Pod, plReq.Container, int64(plReq.Lines), plReq.Follow)
	if err != nil {
		response.FailWithMessage("获取失败", c)
		global.KOP_LOG.Error("获取pod日志失败", zap.Error(err))
	} else {
		response.OkWithDetailed(log, "获取成功", c)
	}
}

// DeletePod 删除pod
func (pa *PodApi) DeletePod(c *gin.Context) {
	var pdReq k8sRequest.PodReq
	_ = c.ShouldBindJSON(&pdReq)
	// 校验
	validate := validator.New()
	if err := validate.Struct(&pdReq); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := podService.DeletePod(pdReq.Namespace, pdReq.Pod); err != nil {
		response.FailWithMessage("删除失败", c)
		global.KOP_LOG.Error("删除失败", zap.Error(err))
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// GetPodWebSSH pod WebSSH
func (pa *PodApi) GetPodWebSSH(c *gin.Context) {
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

	err = podService.GetPodWebSSH(namespace, pod, container, global.KOP_K8S_CONFIG, session)
	if err != nil {
		global.KOP_LOG.Error("terminal失败", zap.Error(err))
	}
}

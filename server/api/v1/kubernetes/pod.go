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
	"github.com/pddzl/kubeops/server/model/kubernetes/resource"
)

type PodApi struct{}

// 获取集群所有pod

func (p *PodApi) GetPodList(c *gin.Context) {
	var pageInfo request.SearchPodParams
	var list []resource.Pod
	_ = c.ShouldBindJSON(&pageInfo)
	// 校验字段
	validate := validator.New()
	if err := validate.Struct(&pageInfo.PageInfo); err != nil {
		global.KOP_LOG.Error("请求参数有误", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}
	podList, total, err := podService.GetPodList(pageInfo.NameSpace, pageInfo.PageInfo)
	if err != nil {
		global.KOP_LOG.Error("获取pod列表失败", zap.Error(err))
	}
	// 处理pod原始数据
	for _, podRaw := range podList {
		var pod resource.Pod
		pod.Name = podRaw.Name
		pod.Namespace = podRaw.Namespace
		pod.Image = podRaw.Spec.Containers[0].Image
		pod.Node = podRaw.Spec.NodeName
		pod.Status = string(podRaw.Status.Phase)
		pod.CreationTimestamp = podRaw.CreationTimestamp
		// append
		list = append(list, pod)
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    int64(total),
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

// 获取pod详情

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

// 获取pod in raw

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

// 获取pod日志

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
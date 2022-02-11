package kubernetes

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"

	"github.com/pddzl/kubeops/server/global"
	"github.com/pddzl/kubeops/server/model/common/response"
	"github.com/pddzl/kubeops/server/model/kubernetes"
	"github.com/pddzl/kubeops/server/model/kubernetes/request"
)

type PodApi struct{}

// 获取集群所有pod

func (p *PodApi) GetPodList(c *gin.Context) {
	var pageInfo request.SearchPodParams
	var list []kubernetes.Pod
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
		var pod kubernetes.Pod
		pod.Name = podRaw.Name
		pod.Namespace = podRaw.Namespace
		pod.Image = podRaw.Spec.Containers[0].Image
		pod.Node = podRaw.Spec.NodeName
		pod.Resource.CpuLimit = podRaw.Spec.Containers[0].Resources.Limits.Cpu().String()
		pod.Resource.MemoryLimit = podRaw.Spec.Containers[0].Resources.Limits.Memory().String()
		pod.Resource.CpuRequests = podRaw.Spec.Containers[0].Resources.Requests.Cpu().String()
		pod.Resource.MemoryRequests = podRaw.Spec.Containers[0].Resources.Requests.Memory().String()
		pod.Status = string(podRaw.Status.Phase)
		pod.CreateTimestamp = podRaw.CreationTimestamp.Time
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

func (p *PodApi) GetPodLog(c *gin.Context) {
	var podLog request.PodLog
	_ = c.ShouldBindJSON(&podLog)
	// 字段校验
	validate := validator.New()
	if err := validate.Struct(&podLog); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	log, err := podService.GetPodLog(podLog.NameSpace, podLog.Pod, podLog.Lines)
	if err != nil {
		global.KOP_LOG.Error(err.Error())
		response.FailWithMessage("获取日志失败", c)
		return
	}
	response.OkWithDetailed(log, "获取日志成功", c)
}
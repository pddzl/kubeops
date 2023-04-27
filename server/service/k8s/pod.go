package k8s

import (
	"context"
	"github.com/pddzl/kubeops/server/global"
	modelK8s "github.com/pddzl/kubeops/server/model/k8s"
	coreV1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type PodService struct{}

func (ps *PodService) GetPods(namespace string, page int, pageSize int) ([]modelK8s.PodBrief, int, error) {
	// 获取pod list
	list, err := global.KOP_K8S_Client.CoreV1().Pods(namespace).List(context.TODO(), metaV1.ListOptions{})
	if err != nil {
		return nil, 0, err
	}

	var podList coreV1.PodList
	// 分页
	end := pageSize * page
	offset := pageSize * (page - 1)
	total := len(list.Items)
	if total <= offset {
		return nil, total, nil
	}
	if total < end {
		podList.Items = list.Items[offset:]
	} else {
		podList.Items = list.Items[offset:end]
	}

	// 处理list数据
	var podBriefList []modelK8s.PodBrief
	for _, pod := range podList.Items {
		var podBrief modelK8s.PodBrief
		podBrief.Name = pod.Name
		podBrief.Namespace = pod.Namespace
		podBrief.Node = pod.Spec.NodeName
		podBrief.Status = string(pod.Status.Phase)
		podBrief.CreationTimestamp = pod.CreationTimestamp.Time
		// append
		podBriefList = append(podBriefList, podBrief)
	}
	return podBriefList, total, nil
}

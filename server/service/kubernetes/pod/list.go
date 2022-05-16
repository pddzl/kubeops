package pod

import (
	"context"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/pddzl/kubeops/server/global"
	"github.com/pddzl/kubeops/server/model/common/request"
	resourcePod "github.com/pddzl/kubeops/server/model/kubernetes/resource/pod"
)

// 获取pod list

func (p *PodService) GetPodList(namespace string, info request.PageInfo) ([]resourcePod.PodBrief, int, error) {
	// 获取pod list
	list, err := global.KOP_KUBERNETES.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, 0, err
	}

	var podList corev1.PodList
	// 分页
	end := info.PageSize * info.Page
	offset := info.PageSize * (info.Page - 1)
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
	var podBriefList []resourcePod.PodBrief
	for _, pod := range podList.Items {
		var podBrief resourcePod.PodBrief
		podBrief.Name = pod.Name
		podBrief.Namespace = pod.Namespace
		podBrief.Node = pod.Spec.NodeName
		podBrief.Status = string(pod.Status.Phase)
		podBrief.CreationTimestamp = pod.CreationTimestamp
		// append
		podBriefList = append(podBriefList, podBrief)
	}
	return podBriefList, total, nil
}

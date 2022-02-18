package pod

import (
	"context"
	v1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/pddzl/kubeops/server/global"
	"github.com/pddzl/kubeops/server/model/common/request"
)

// 获取pod list

func (p *PodService) GetPodList(namespace string, info request.PageInfo) (list []v1.Pod, total int, err error) {
	end := info.PageSize * info.Page
	offset := info.PageSize * (info.Page - 1)
	podList, err := global.KOP_KUBERNETES.CoreV1().Pods(namespace).List(context.TODO(), metaV1.ListOptions{})
	if err != nil {
		return nil, 0, err
	}
	total = len(podList.Items)
	if total <= offset {
		return nil, total, nil
	}
	if total < end {
		list = podList.Items[offset:]
	} else {
		list = podList.Items[offset:end]
	}
	return list, total, nil
}

package daemonSet

import (
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"

	"github.com/pddzl/kubeops/server/global"
	"github.com/pddzl/kubeops/server/model/common/request"
	"github.com/pddzl/kubeops/server/model/kubernetes/api"
	"github.com/pddzl/kubeops/server/model/kubernetes/resource/common"
)

func (d *DaemonSetService) GetDaemonSetPods(namespace string, name string, info request.PageInfo) ([]common.RelatedPod, int, error) {
	// 获取原生daemonSet数据
	daemonSet, err := global.KOP_KUBERNETES.AppsV1().DaemonSets(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		return nil, 0, err
	}

	//获取daemonSet的selector
	selector := labels.SelectorFromSet(daemonSet.Spec.Selector.MatchLabels)
	options := metav1.ListOptions{LabelSelector: selector.String()}

	// 获取daemonSet关联pods
	podList, err := global.KOP_KUBERNETES.CoreV1().Pods(namespace).List(context.TODO(), options)
	if err != nil {
		return nil, 0, err
	}

	// 处理related Pods
	var relatedPodList []common.RelatedPod
	for _, pod := range podList.Items {
		var relatedPod common.RelatedPod
		relatedPod.ObjectMeta = api.NewObjectMeta(pod.ObjectMeta)
		relatedPod.Status = string(pod.Status.Phase)
		relatedPod.NodeName = pod.Spec.NodeName
		// append
		relatedPodList = append(relatedPodList, relatedPod)
	}

	// 分页
	end := info.PageSize * info.Page
	offset := info.PageSize * (info.Page - 1)
	total := len(podList.Items)
	if total <= offset {
		return nil, total, nil
	}
	if total < end {
		return relatedPodList[offset:], total, nil
	} else {
		return relatedPodList[offset:end], total, nil
	}
}

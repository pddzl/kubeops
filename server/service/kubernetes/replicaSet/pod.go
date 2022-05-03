package replicaSet

import (
	"context"
	"github.com/pddzl/kubeops/server/model/kubernetes/resource/common"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"

	"github.com/pddzl/kubeops/server/global"
	"github.com/pddzl/kubeops/server/model/common/request"
	"github.com/pddzl/kubeops/server/model/kubernetes/api"
)

func (r *ReplicaSetService) GetReplicaSetPods(namespace string, name string, info request.PageInfo) ([]common.RelatedPod, int, error) {
	// 获取replicaSet原始数据
	rs, err := global.KOP_KUBERNETES.AppsV1().ReplicaSets(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		return nil, 0, err
	}

	// 获取replicaSet的selector
	selector := labels.SelectorFromSet(rs.Spec.Selector.MatchLabels)
	options := metav1.ListOptions{LabelSelector: selector.String()}

	// 获取pods
	podList, err := global.KOP_KUBERNETES.CoreV1().Pods(namespace).List(context.TODO(), options)
	if err != nil {
		return nil, 0, err
	}

	// 处理related pod
	var relatedPodList []common.RelatedPod
	for _, pod := range podList.Items {
		var relatedPod common.RelatedPod
		relatedPod.ObjectMeta = api.NewObjectMeta(pod.ObjectMeta)
		relatedPod.NodeName = pod.Spec.NodeName
		relatedPod.Status = string(pod.Status.Phase)
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

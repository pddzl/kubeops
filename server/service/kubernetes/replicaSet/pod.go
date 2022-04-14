package replicaSet

import (
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"

	"github.com/pddzl/kubeops/server/global"
	"github.com/pddzl/kubeops/server/model/common/request"
	"github.com/pddzl/kubeops/server/model/kubernetes/api"
	"github.com/pddzl/kubeops/server/model/kubernetes/resource/replicaSet"
)

func (r *ReplicaSetService) GetReplicaSetPods(namespace string, name string, info request.PageInfo) ([]replicaSet.Pod, int, error) {
	// 获取replicaSet原始数据
	rs, err := global.KOP_KUBERNETES.AppsV1().ReplicaSets(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		return nil, 0, err
	}

	// 获取replicaSet的selector
	selector := labels.SelectorFromSet(rs.Spec.Selector.MatchLabels)
	options := metav1.ListOptions{LabelSelector: selector.String()}

	// 获取pods
	var replicaSetPods []replicaSet.Pod
	podList, err := global.KOP_KUBERNETES.CoreV1().Pods(namespace).List(context.TODO(), options)
	if err != nil {
		return nil, 0, err
	}

	// 处理replicaSet Pods
	for _, pod := range podList.Items {
		replicaSetPod := replicaSet.Pod{}
		replicaSetPod.ObjectMeta = api.NewObjectMeta(pod.ObjectMeta)
		replicaSetPod.Status = string(pod.Status.Phase)
		replicaSetPod.NodeName = pod.Spec.NodeName
		// append
		replicaSetPods = append(replicaSetPods, replicaSetPod)
	}

	// 分页
	end := info.PageSize * info.Page
	offset := info.PageSize * (info.Page - 1)
	total := len(podList.Items)
	if total <= offset {
		return nil, total, nil
	}
	if total < end {
		return replicaSetPods[offset:], total, nil
	} else {
		return replicaSetPods[offset:end], total, nil
	}
}

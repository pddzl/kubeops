package daemonSet

import (
	"context"
	"github.com/pddzl/kubeops/server/global"
	"github.com/pddzl/kubeops/server/model/common/request"
	"github.com/pddzl/kubeops/server/model/kubernetes/api"
	resourceDaemonSet "github.com/pddzl/kubeops/server/model/kubernetes/resource/daemonSet"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
)

func (d *DaemonSetService) GetDaemonSetPods(namespace string, name string, info request.PageInfo) ([]resourceDaemonSet.Pod, int, error) {
	// 获取原生daemonSet数据
	daemonSet, err := global.KOP_KUBERNETES.AppsV1().DaemonSets(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		return nil, 0, err
	}

	//获取daemonSet的selector
	selector := labels.SelectorFromSet(daemonSet.Spec.Selector.MatchLabels)
	options := metav1.ListOptions{LabelSelector: selector.String()}

	// 获取pods
	podList, err := global.KOP_KUBERNETES.CoreV1().Pods(namespace).List(context.TODO(), options)
	if err != nil {
		return nil, 0, err
	}

	// 处理replicaSet Pods
	var daemonSetPods []resourceDaemonSet.Pod
	for _, pod := range podList.Items {
		daemonSetPod := resourceDaemonSet.Pod{}
		daemonSetPod.ObjectMeta = api.NewObjectMeta(pod.ObjectMeta)
		daemonSetPod.Status = string(pod.Status.Phase)
		daemonSetPod.NodeName = pod.Spec.NodeName
		// append
		daemonSetPods = append(daemonSetPods, daemonSetPod)
	}

	// 分页
	end := info.PageSize * info.Page
	offset := info.PageSize * (info.Page - 1)
	total := len(podList.Items)
	if total <= offset {
		return nil, total, nil
	}
	if total < end {
		return daemonSetPods[offset:], total, nil
	} else {
		return daemonSetPods[offset:end], total, nil
	}

}

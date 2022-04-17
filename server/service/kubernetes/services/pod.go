package services

import (
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"

	"github.com/pddzl/kubeops/server/global"
	"github.com/pddzl/kubeops/server/model/common/request"
	"github.com/pddzl/kubeops/server/model/kubernetes/api"
	"github.com/pddzl/kubeops/server/model/kubernetes/resource/common"
)

func (s *ServicesService) GetServicesPods(namespace string, name string, info request.PageInfo) ([]common.RelatedPods, int, error) {
	// 获取原生service数据
	service, err := global.KOP_KUBERNETES.CoreV1().Services(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		return nil, 0, err
	}

	// 获取services的selector
	selector := labels.SelectorFromSet(service.Spec.Selector)
	options := metav1.ListOptions{LabelSelector: selector.String()}

	// 获取service关联的pod
	podList, err := global.KOP_KUBERNETES.CoreV1().Pods(namespace).List(context.TODO(), options)
	if err != nil {
		return nil, 0, err
	}

	var servicePods []common.RelatedPods

	// 处理replicaSet Pods
	for _, pod := range podList.Items {
		var servicePod common.RelatedPods
		servicePod.ObjectMeta = api.NewObjectMeta(pod.ObjectMeta)
		servicePod.Status = string(pod.Status.Phase)
		servicePod.NodeName = pod.Spec.NodeName
		// append
		servicePods = append(servicePods, servicePod)
	}

	// 分页
	end := info.PageSize * info.Page
	offset := info.PageSize * (info.Page - 1)
	total := len(podList.Items)
	if total <= offset {
		return nil, total, nil
	}
	if total < end {
		return servicePods[offset:], total, nil
	} else {
		return servicePods[offset:end], total, nil
	}
}

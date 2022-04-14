package replicaSet

import (
	"context"
	"github.com/pddzl/kubeops/server/global"
	"github.com/pddzl/kubeops/server/model/kubernetes/resource/common"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (r *ReplicaSetService) GetReplicaSetServices(namespace string, name string) ([]v1.Service, error) {
	// 获取原生replicaSet数据
	replicaSet, err := global.KOP_KUBERNETES.AppsV1().ReplicaSets(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}

	// -> why List Services with selector doesn't work
	//selector := labels.SelectorFromSet(replicaSet.Spec.Selector.MatchLabels)
	//options := metav1.ListOptions{LabelSelector: selector.String()}

	// 获取services
	services, err := global.KOP_KUBERNETES.CoreV1().Services(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	matchingServices := common.FilterNamespacedServicesBySelector(services.Items, namespace, replicaSet.Spec.Selector.MatchLabels)

	return matchingServices, nil
}

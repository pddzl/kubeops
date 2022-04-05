package replicaSet

import (
	"context"
	"github.com/pddzl/kubeops/server/global"
	"github.com/pddzl/kubeops/server/model/kubernetes/resource/common"
	v1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (r *ReplicaSetService) GetReplicaSetServices(namespace string, replicaSetName string) ([]v1.Service, error) {
	replicaSet, err := global.KOP_KUBERNETES.AppsV1().ReplicaSets(namespace).Get(context.TODO(), replicaSetName, metaV1.GetOptions{})
	if err != nil {
		return nil, err
	}

	channels := &common.ResourceChannels{
		ServiceList: common.GetServiceListChannel(common.NewSameNamespaceQuery(namespace), 1),
	}
	services := <-channels.ServiceList.List
	err = <-channels.ServiceList.Error
	if err != nil {
		return nil, err
	}

	matchingServices := common.FilterNamespacedServicesBySelector(services.Items, namespace, replicaSet.Spec.Selector.MatchLabels)

	return matchingServices, nil
}

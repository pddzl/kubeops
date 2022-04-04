package replicaSet

import (
	"context"
	"github.com/pddzl/kubeops/server/global"
	"github.com/pddzl/kubeops/server/model/kubernetes/api"
	"github.com/pddzl/kubeops/server/model/kubernetes/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (r *ReplicaSetService) GetReplicaSetDetail(namespace string, name string) (*resource.ReplicaSetDetail, error) {
	var replicaSetDetail resource.ReplicaSetDetail

	replicaSet, err := global.KOP_KUBERNETES.AppsV1().ReplicaSets(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}

	podInfo, err := getReplicaSetPodInfo(replicaSet)
	if err != nil {
		return nil, err
	}

	replicaSetDetail.ObjectMeta = api.NewObjectMeta(replicaSet.ObjectMeta)
	replicaSetDetail.TypeMeta = api.NewTypeMeta(api.ResourceKindReplicaSet)
	replicaSetDetail.Selector = replicaSet.Spec.Selector
	replicaSetDetail.Pods = *podInfo

	return &replicaSetDetail, nil
}

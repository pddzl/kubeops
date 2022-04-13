package replicaSet

import (
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/pddzl/kubeops/server/global"
	"github.com/pddzl/kubeops/server/model/kubernetes/api"
	replicaSet2 "github.com/pddzl/kubeops/server/model/kubernetes/resource/replicaSet"
)

func (r *ReplicaSetService) GetReplicaSetDetail(namespace string, replicaSetName string) (*replicaSet2.ReplicaSetDetail, error) {
	replicaSet, err := global.KOP_KUBERNETES.AppsV1().ReplicaSets(namespace).Get(context.TODO(), replicaSetName, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}

	// 处理replicaSet数据
	var replicaSetDetail replicaSet2.ReplicaSetDetail
	// metadata
	replicaSetDetail.ObjectMeta = api.NewObjectMeta(replicaSet.ObjectMeta)
	// spec
	replicaSetDetail.Spec.Replicas = replicaSet.Spec.Replicas
	replicaSetDetail.Spec.MinReadySeconds = replicaSet.Spec.MinReadySeconds
	replicaSetDetail.Spec.Selector = replicaSet.Spec.Selector.MatchLabels
	// status
	replicaSetDetail.Status.Replicas = replicaSet.Status.Replicas
	replicaSetDetail.Status.FullyLabeledReplicas = replicaSet.Status.FullyLabeledReplicas
	replicaSetDetail.Status.ReadyReplicas = replicaSet.Status.ReadyReplicas
	replicaSetDetail.Status.AvailableReplicas = replicaSet.Status.AvailableReplicas
	// status -> condition
	for _, condition := range replicaSet.Status.Conditions {
		var replicaSetCondition replicaSet2.ReplicaSetCondition
		replicaSetCondition.Type = string(condition.Type)
		replicaSetCondition.Status = string(condition.Status)
		replicaSetCondition.LastTransitionTime = condition.LastTransitionTime
		replicaSetCondition.Reason = condition.Reason
		replicaSetCondition.Message = condition.Message
		// append
		replicaSetDetail.Status.Conditions = append(replicaSetDetail.Status.Conditions, replicaSetCondition)
	}

	return &replicaSetDetail, nil
}

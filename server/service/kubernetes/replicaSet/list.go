package replicaSet

import (
	"context"
	"fmt"
	"github.com/pddzl/kubeops/server/model/kubernetes/resource/replicaSet"
	v1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/pddzl/kubeops/server/global"
	"github.com/pddzl/kubeops/server/model/common/request"
)

func (r *ReplicaSetService) GetReplicaSetList(namespace string, info request.PageInfo) (list interface{}, total int, err error) {
	end := info.PageSize * info.Page
	offset := info.PageSize * (info.Page - 1)
	var replicaSetBriefList []replicaSet.ReplicaSetBrief
	var replicaSetList v1.ReplicaSetList

	replicaSets, err := global.KOP_KUBERNETES.AppsV1().ReplicaSets(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, 0, err
	}

	// 分页
	total = len(replicaSets.Items)
	if total <= offset {
		return nil, total, nil
	}
	if total < end {
		replicaSetList.Items = replicaSets.Items[offset:]
	} else {
		replicaSetList.Items = replicaSets.Items[offset:end]
	}

	// 处理replicaSet数据
	for _, rs := range replicaSets.Items {
		var replicaSetBrief replicaSet.ReplicaSetBrief
		replicaSetBrief.Name = rs.Name
		replicaSetBrief.NameSpace = rs.Namespace
		replicaSetBrief.Pods = fmt.Sprintf("%d / %d", rs.Status.AvailableReplicas, rs.Status.Replicas)
		replicaSetBrief.CreationTimestamp = rs.CreationTimestamp
		// append
		replicaSetBriefList = append(replicaSetBriefList, replicaSetBrief)
	}

	return replicaSetBriefList, total, nil
}

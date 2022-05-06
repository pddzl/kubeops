package replicaSet

import (
	"context"
	"fmt"
	"github.com/pddzl/kubeops/server/model/kubernetes/resource/replicaSet"
	appsV1 "k8s.io/api/apps/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/pddzl/kubeops/server/global"
	"github.com/pddzl/kubeops/server/model/common/request"
)

func (r *ReplicaSetService) GetReplicaSetList(namespace string, info request.PageInfo) ([]replicaSet.ReplicaSetBrief, int, error) {
	// 获取replicaSet list
	list, err := global.KOP_KUBERNETES.AppsV1().ReplicaSets(namespace).List(context.TODO(), metaV1.ListOptions{})
	if err != nil {
		return nil, 0, err
	}

	var replicaSetList appsV1.ReplicaSetList
	// 分页
	end := info.PageSize * info.Page
	offset := info.PageSize * (info.Page - 1)
	total := len(list.Items)
	if total <= offset {
		return nil, total, nil
	}
	if total < end {
		replicaSetList.Items = list.Items[offset:]
	} else {
		replicaSetList.Items = list.Items[offset:end]
	}

	// 处理replicaSet数据
	var replicaSetBriefList []replicaSet.ReplicaSetBrief
	for _, rs := range list.Items {
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

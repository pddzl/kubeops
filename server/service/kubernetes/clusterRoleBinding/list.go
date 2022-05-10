package clusterRoleBinding

import (
	"context"
	rbacV1 "k8s.io/api/rbac/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/pddzl/kubeops/server/global"
	"github.com/pddzl/kubeops/server/model/common/request"
	resourceClusterRoleBinding "github.com/pddzl/kubeops/server/model/kubernetes/resource/clusterRoleBinding"
)

func (cb *ClusterRoleBindingService) GetClusterRoleBindingList(info request.PageInfo) ([]resourceClusterRoleBinding.ClusterRoleBindingBrief, int, error) {
	list, err := global.KOP_KUBERNETES.RbacV1().ClusterRoleBindings().List(context.TODO(), metaV1.ListOptions{})
	if err != nil {
		return nil, 0, err
	}

	var clusterRoleBindingList rbacV1.ClusterRoleBindingList
	// 分页
	end := info.PageSize * info.Page
	offset := info.PageSize * (info.Page - 1)
	total := len(list.Items)
	if total <= offset {
		return nil, total, nil
	}
	if total < end {
		clusterRoleBindingList.Items = list.Items[offset:]
	} else {
		clusterRoleBindingList.Items = list.Items[offset:end]
	}

	var clusterRoleBindingBriefList []resourceClusterRoleBinding.ClusterRoleBindingBrief
	// 处理数据
	for _, role := range clusterRoleBindingList.Items {
		var clusterRoleBindingBrief resourceClusterRoleBinding.ClusterRoleBindingBrief
		clusterRoleBindingBrief.Name = role.Name
		clusterRoleBindingBrief.CreationTimestamp = role.CreationTimestamp
		// append
		clusterRoleBindingBriefList = append(clusterRoleBindingBriefList, clusterRoleBindingBrief)
	}

	return clusterRoleBindingBriefList, total, nil
}

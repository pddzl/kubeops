package clusterRole

import (
	"context"
	rbacV1 "k8s.io/api/rbac/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/pddzl/kubeops/server/global"
	"github.com/pddzl/kubeops/server/model/common/request"
	resourceClusterRole "github.com/pddzl/kubeops/server/model/kubernetes/resource/clusterRole"
)

func (c *ClusterRoleService) GetClusterRoleList(info request.PageInfo) ([]resourceClusterRole.ClusterRoleBrief, int, error) {
	// 获取clusterRole list
	list, err := global.KOP_KUBERNETES.RbacV1().ClusterRoles().List(context.TODO(), metaV1.ListOptions{})
	if err != nil {
		return nil, 0, err
	}

	var clusterRoleList rbacV1.ClusterRoleList
	// 分页
	end := info.PageSize * info.Page
	offset := info.PageSize * (info.Page - 1)
	total := len(list.Items)
	if total <= offset {
		return nil, total, nil
	}
	if total < end {
		clusterRoleList.Items = list.Items[offset:]
	} else {
		clusterRoleList.Items = list.Items[offset:end]
	}

	// 处理数据
	var clusterRoleBriefList []resourceClusterRole.ClusterRoleBrief
	for _, clusterRole := range clusterRoleList.Items {
		var clusterRoleBrief resourceClusterRole.ClusterRoleBrief
		clusterRoleBrief.Name = clusterRole.Name
		clusterRoleBrief.CreationTimestamp = clusterRole.CreationTimestamp
		// append
		clusterRoleBriefList = append(clusterRoleBriefList, clusterRoleBrief)
	}

	return clusterRoleBriefList, total, nil
}

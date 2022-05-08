package roleBinding

import (
	"context"
	"github.com/pddzl/kubeops/server/global"
	"github.com/pddzl/kubeops/server/model/common/request"
	resourceRoleBinding "github.com/pddzl/kubeops/server/model/kubernetes/resource/roleBinding"
	rbacV1 "k8s.io/api/rbac/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (r *RoleBindingService) GetRoleBindingList(namespace string, info request.PageInfo) ([]resourceRoleBinding.RoleBindingBrief, int, error) {
	// 获取roleBinding list
	list, err := global.KOP_KUBERNETES.RbacV1().RoleBindings(namespace).List(context.TODO(), metaV1.ListOptions{})
	if err != nil {
		return nil, 0, err
	}

	var roleBindingList rbacV1.RoleBindingList
	// 分页
	end := info.PageSize * info.Page
	offset := info.PageSize * (info.Page - 1)
	total := len(list.Items)
	if total <= offset {
		return nil, total, nil
	}
	if total < end {
		roleBindingList.Items = list.Items[offset:]
	} else {
		roleBindingList.Items = list.Items[offset:end]
	}

	var roleBindingBriefList []resourceRoleBinding.RoleBindingBrief
	// 处理数据
	for _, roleBinding := range roleBindingList.Items {
		var roleBindingBrief resourceRoleBinding.RoleBindingBrief
		roleBindingBrief.Name = roleBinding.Name
		roleBindingBrief.Namespace = roleBinding.Namespace
		roleBindingBrief.CreationTimestamp = roleBinding.CreationTimestamp
		// append
		roleBindingBriefList = append(roleBindingBriefList, roleBindingBrief)
	}

	return roleBindingBriefList, total, nil
}

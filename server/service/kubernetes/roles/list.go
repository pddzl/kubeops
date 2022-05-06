package role

import (
	"context"
	resourceRoles "github.com/pddzl/kubeops/server/model/kubernetes/resource/role"
	rbacV1 "k8s.io/api/rbac/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/pddzl/kubeops/server/global"
	"github.com/pddzl/kubeops/server/model/common/request"
)

func (role *RoleService) GetRoleList(namespace string, info request.PageInfo) ([]resourceRoles.RolesBrief, int, error) {
	// 获取roles list
	list, err := global.KOP_KUBERNETES.RbacV1().Roles(namespace).List(context.TODO(), metaV1.ListOptions{})
	if err != nil {
		return nil, 0, err
	}

	var roleList rbacV1.RoleList
	// 分页
	end := info.PageSize * info.Page
	offset := info.PageSize * (info.Page - 1)
	total := len(list.Items)
	if total <= offset {
		return nil, total, nil
	}
	if total < end {
		roleList.Items = list.Items[offset:]
	} else {
		roleList.Items = list.Items[offset:end]
	}

	// 处理数据
	var roleBriefList []resourceRoles.RolesBrief
	for _, role := range roleList.Items {
		var roleBrief resourceRoles.RolesBrief
		roleBrief.Name = role.Name
		roleBrief.Namespace = role.Namespace
		roleBrief.CreationTimestamp = role.CreationTimestamp
		// append
		roleBriefList = append(roleBriefList, roleBrief)
	}

	return roleBriefList, total, nil
}

package role

import (
	"context"
	"github.com/pddzl/kubeops/server/global"
	rbacV1 "k8s.io/api/rbac/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (r *RoleService) GetRoleDetail(namespace string, name string) (*rbacV1.Role, error) {
	detail, err := global.KOP_KUBERNETES.RbacV1().Roles(namespace).Get(context.TODO(), name, metaV1.GetOptions{})
	if err != nil {
		return nil, err
	}
	return detail, nil
}

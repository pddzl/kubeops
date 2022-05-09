package roleBinding

import (
	"context"
	rbacV1 "k8s.io/api/rbac/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/pddzl/kubeops/server/global"
)

func (r *RoleBindingService) GetRoleBindingDetail(namespace string, name string) (*rbacV1.RoleBinding, error) {
	detail, err := global.KOP_KUBERNETES.RbacV1().RoleBindings(namespace).Get(context.TODO(), name, metaV1.GetOptions{})
	if err != nil {
		return nil, err
	}

	return detail, nil
}

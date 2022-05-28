package role

import (
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/pddzl/kubeops/server/global"
)

func (r *RoleService) DeleteRole(namespace string, name string) error {
	return global.KOP_KUBERNETES.RbacV1().Roles(namespace).Delete(context.TODO(), name, metav1.DeleteOptions{})
}

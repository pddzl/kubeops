package roleBinding

import (
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/pddzl/kubeops/server/global"
)

func (r *RoleBindingService) DeleteRoleBinding(namespace string, name string) error {
	return global.KOP_KUBERNETES.RbacV1().RoleBindings(namespace).Delete(context.TODO(), name, metav1.DeleteOptions{})
}

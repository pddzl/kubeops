package clusterRoleBinding

import (
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/pddzl/kubeops/server/global"
)

func (c *ClusterRoleBindingService) DeleteClusterRoleBinding(name string) error {
	return global.KOP_KUBERNETES.RbacV1().ClusterRoleBindings().Delete(context.TODO(), name, metav1.DeleteOptions{})
}

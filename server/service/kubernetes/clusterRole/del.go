package clusterRole

import (
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/pddzl/kubeops/server/global"
)

func (c *ClusterRoleService) DeleteClusterRole(name string) error {
	return global.KOP_KUBERNETES.RbacV1().ClusterRoles().Delete(context.TODO(), name, metav1.DeleteOptions{})
}

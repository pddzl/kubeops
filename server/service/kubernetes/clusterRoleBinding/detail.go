package clusterRoleBinding

import (
	"context"
	rbacV1 "k8s.io/api/rbac/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/pddzl/kubeops/server/global"
)

func (c *ClusterRoleBindingService) GetClusterRoleBinding(name string) (*rbacV1.ClusterRoleBinding, error) {
	detail, err := global.KOP_KUBERNETES.RbacV1().ClusterRoleBindings().Get(context.TODO(), name, metaV1.GetOptions{})
	if err != nil {
		return nil, err
	}

	return detail, nil
}

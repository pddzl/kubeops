package clusterRole

import (
	"context"
	"github.com/pddzl/kubeops/server/global"
	rbacV1 "k8s.io/api/rbac/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (c *ClusterRoleService) GetClusterRoleDetail(name string) (*rbacV1.ClusterRole, error) {
	detail, err := global.KOP_KUBERNETES.RbacV1().ClusterRoles().Get(context.TODO(), name, metaV1.GetOptions{})
	if err != nil {
		return nil, err
	}

	return detail, nil
}

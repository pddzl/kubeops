package clusterRole

import (
	"context"
	"encoding/json"

	"github.com/pddzl/kubeops/server/global"
)

func (c *ClusterRoleService) GetClusterRoleRaw(name string) (interface{}, error) {
	// kubectl get --raw /apis/rbac.authorization.k8s.io/v1/clusterroles/ingress-nginx
	raw, err := global.KOP_KUBERNETES.RESTClient().Get().AbsPath("/apis/rbac.authorization.k8s.io/v1").Resource("clusterroles").Name(name).DoRaw(context.TODO())
	if err != nil {
		return nil, err
	}
	return json.RawMessage(raw), nil
}

package clusterRoleBinding

import (
	"context"
	"encoding/json"

	"github.com/pddzl/kubeops/server/global"
)

func (c *ClusterRoleBindingService) GetClusterRoleBindingRaw(name string) (interface{}, error) {
	// kubectl get --raw /apis/rbac.authorization.k8s.io/v1/clusterrolebindings/ingress-nginx
	raw, err := global.KOP_KUBERNETES.RESTClient().Get().AbsPath("/apis/rbac.authorization.k8s.io/v1").Resource("clusterrolebindings").Name(name).DoRaw(context.TODO())
	if err != nil {
		return nil, err
	}
	return json.RawMessage(raw), nil
}

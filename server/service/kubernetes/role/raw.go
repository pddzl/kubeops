package role

import (
	"context"
	"encoding/json"

	"github.com/pddzl/kubeops/server/global"
)

func (r *RoleService) GetRoleRaw(namespace string, name string) (interface{}, error) {
	// kubectl get --raw /apis/rbac.authorization.k8s.io/v1/namespaces/ingress-nginx/role/ingress-nginx
	raw, err := global.KOP_KUBERNETES.RESTClient().Get().AbsPath("/apis/rbac.authorization.k8s.io/v1").Resource("roles").Namespace(namespace).Name(name).DoRaw(context.TODO())
	if err != nil {
		return nil, err
	}
	return json.RawMessage(raw), nil
}

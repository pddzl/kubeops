package roleBinding

import (
	"context"
	"encoding/json"
	"github.com/pddzl/kubeops/server/global"
)

func (r *RoleBindingService) GetRoleBindingRaw(namespace string, name string) (interface{}, error) {
	// kubectl get --raw /apis/rbac.authorization.k8s.io/v1/namespaces/ingress-nginx/rolebindings/ingress-nginx
	raw, err := global.KOP_KUBERNETES.RESTClient().Get().AbsPath("/apis/rbac.authorization.k8s.io/v1").Resource("rolebindings").Namespace(namespace).Name(name).DoRaw(context.TODO())
	if err != nil {
		return nil, err
	}
	return json.RawMessage(raw), nil
}

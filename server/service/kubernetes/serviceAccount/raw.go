package serviceAccount

import (
	"context"
	"encoding/json"
	"github.com/pddzl/kubeops/server/global"
)

func (sa *ServiceAccountService) GetServiceAccountRaw(namespace string, name string) (raw interface{}, err error) {
	// kubectl get --raw '/api/v1/namespaces/test/serviceaccounts'
	req, err := global.KOP_KUBERNETES.RESTClient().Get().AbsPath("/api/v1").Resource("serviceaccounts").Name(name).Namespace(namespace).DoRaw(context.TODO())
	if err != nil {
		return nil, err
	}
	return json.RawMessage(req), nil
}

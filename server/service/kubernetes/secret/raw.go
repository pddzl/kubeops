package secret

import (
	"context"
	"encoding/json"

	"github.com/pddzl/kubeops/server/global"
)

func (s *SecretService) GetSecretRaw(namespace string, name string) (interface{}, error) {
	// kubectl get --raw '/api/v1/namespaces/test/secrets'
	raw, err := global.KOP_KUBERNETES.RESTClient().Get().AbsPath("/api/v1").Resource("secrets").Namespace(namespace).Name(name).DoRaw(context.TODO())
	if err != nil {
		return nil, err
	}
	return json.RawMessage(raw), nil
}

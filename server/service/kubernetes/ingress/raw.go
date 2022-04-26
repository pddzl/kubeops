package ingress

import (
	"context"
	"encoding/json"

	"github.com/pddzl/kubeops/server/global"
)

func (i *IngressService) GetReplicaSetRaw(namespace string, name string) (raw interface{}, err error) {
	// kubectl get --raw '/apis/networking.k8s.io/v1/namespaces/test/ingresses/pddzl-ingress'
	req, err := global.KOP_KUBERNETES.RESTClient().Get().AbsPath("/apis/networking.k8s.io/v1").Resource("ingresses").Name(name).Namespace(namespace).DoRaw(context.TODO())
	if err != nil {
		return nil, err
	}
	return json.RawMessage(req), nil
}

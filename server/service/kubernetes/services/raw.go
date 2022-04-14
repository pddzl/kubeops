package services

import (
	"context"
	"encoding/json"
	"github.com/pddzl/kubeops/server/global"
)

func (s *ServicesService) GetServicesRaw(namespace string, name string) (interface{}, error) {
	// kubectl get --raw '/api/v1/namespaces/kube-system/services/kube-dns'
	req, err := global.KOP_KUBERNETES.RESTClient().Get().AbsPath("/api/v1").Resource("services").Name(name).Namespace(namespace).DoRaw(context.TODO())
	if err != nil {
		return nil, err
	}
	return json.RawMessage(req), nil
}

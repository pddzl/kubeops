package pod

import (
	"context"
	"encoding/json"

	"github.com/pddzl/kubeops/server/global"
)

func (p *PodService) GetPodRaw(namespace string, name string) (raw interface{}, err error) {
	// kubectl get --raw '/api/v1/namespaces/test/pods'
	req, err := global.KOP_KUBERNETES.RESTClient().Get().AbsPath("/api/v1").Resource("pods").Name(name).Namespace(namespace).DoRaw(context.TODO())
	if err != nil {
		return nil, err
	}
	return json.RawMessage(req), nil
}

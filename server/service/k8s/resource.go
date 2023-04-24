package k8s

import (
	"context"
	"encoding/json"
	"github.com/pddzl/kubeops/server/global"
)

type ResourceService struct{}

func (rs *ResourceService) GetResourceRaw(name string, resource string) (raw interface{}, err error) {
	var req []byte
	switch resource {
	case "nodes":
		req, err = global.KOP_K8S_Client.RESTClient().Get().AbsPath("/api/v1").Resource("nodes").Name(name).DoRaw(context.TODO())
	}

	return json.RawMessage(req), nil
}

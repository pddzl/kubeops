package namespace

import (
	"context"
	"encoding/json"

	"github.com/pddzl/kubeops/server/global"
)

func (n *NamespaceService) GetNamespaceRaw(name string) (raw interface{}, err error) {
	//absPath := fmt.Sprintf("/api/v1/namespaces/%s", name)
	//_, err = global.KOP_KUBERNETES.RESTClient().Get().AbsPath(absPath).DoRaw(context.TODO())
	req, err := global.KOP_KUBERNETES.RESTClient().Get().AbsPath("/api/v1").Resource("namespaces").Name(name).DoRaw(context.TODO())
	if err != nil {
		return nil, err
	}
	return json.RawMessage(req), nil
}

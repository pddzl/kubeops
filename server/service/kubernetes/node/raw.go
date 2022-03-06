package node

import (
	"context"
	"encoding/json"

	"github.com/pddzl/kubeops/server/global"
)

func (n *NodeService) GetNodeRaw(name string) (raw interface{}, err error) {
	req, err := global.KOP_KUBERNETES.RESTClient().Get().AbsPath("/api/v1").Resource("nodes").Name(name).DoRaw(context.TODO())
	if err != nil {
		return nil, err
	}
	return json.RawMessage(req), nil
}

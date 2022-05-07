package replicaSet

import (
	"context"
	"encoding/json"

	"github.com/pddzl/kubeops/server/global"
)

func (r *ReplicaSetService) GetReplicaSetRaw(namespace string, name string) (interface{}, error) {
	// kubectl get --raw '/apis/apps/v1/namespaces/kube-system/replicasets/coredns-6d8c4cb4d'
	req, err := global.KOP_KUBERNETES.RESTClient().Get().AbsPath("/apis/apps/v1").Resource("replicasets").Name(name).Namespace(namespace).DoRaw(context.TODO())
	if err != nil {
		return nil, err
	}
	return json.RawMessage(req), nil
}

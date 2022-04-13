package daemonSet

import (
	"context"
	"encoding/json"

	"github.com/pddzl/kubeops/server/global"
)

func (d *DaemonSetService) GetDaemonSetRaw(namespace string, name string) (raw interface{}, err error) {
	// kubectl get --raw '/apis/apps/v1/namespaces/kube-system/daemonsets/kube-proxy'
	req, err := global.KOP_KUBERNETES.RESTClient().Get().AbsPath("/apis/apps/v1").Resource("daemonsets").Name(name).Namespace(namespace).DoRaw(context.TODO())
	if err != nil {
		return nil, err
	}
	return json.RawMessage(req), nil
}

package deployment

import (
	"context"
	"encoding/json"

	"github.com/pddzl/kubeops/server/global"
)

func (d *DeploymentService) GetDeploymentRaw(namespace string, name string) (raw interface{}, err error) {
	// kubectl get --raw '/apis/apps/v1/namespaces/kube-system/deployments/coredns'
	req, err := global.KOP_KUBERNETES.RESTClient().Get().AbsPath("/apis/apps/v1").Resource("deployments").Name(name).Namespace(namespace).DoRaw(context.TODO())
	if err != nil {
		return nil, err
	}
	return json.RawMessage(req), nil
}

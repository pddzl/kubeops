package configMap

import (
	"context"
	"encoding/json"

	"github.com/pddzl/kubeops/server/global"
)

func (cm *ConfigMapService) GetConfigMapRaw(namespace string, name string) (interface{}, error) {
	// kubectl get --raw '/api/v1/namespaces/test/configmaps'
	raw, err := global.KOP_KUBERNETES.RESTClient().Get().AbsPath("/api/v1").Resource("configmaps").Namespace(namespace).Name(name).DoRaw(context.TODO())
	if err != nil {
		return nil, err
	}
	return json.RawMessage(raw), nil
}

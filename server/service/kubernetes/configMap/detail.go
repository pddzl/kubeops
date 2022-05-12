package configMap

import (
	"context"
	coreV1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/pddzl/kubeops/server/global"
)

func (cm *ConfigMapService) GetConfigMapDetail(namespace string, name string) (*coreV1.ConfigMap, error) {
	detail, err := global.KOP_KUBERNETES.CoreV1().ConfigMaps(namespace).Get(context.TODO(), name, metaV1.GetOptions{})
	if err != nil {
		return nil, err
	}

	return detail, nil
}

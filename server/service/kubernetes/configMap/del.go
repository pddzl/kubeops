package configMap

import (
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/pddzl/kubeops/server/global"
)

func (cm *ConfigMapService) DeleteConfigMap(namespace string, name string) error {
	return global.KOP_KUBERNETES.CoreV1().ConfigMaps(namespace).Delete(context.TODO(), name, metav1.DeleteOptions{})
}

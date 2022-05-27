package secret

import (
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/pddzl/kubeops/server/global"
)

func (s *SecretService) DeleteSecret(namespace string, name string) error {
	return global.KOP_KUBERNETES.CoreV1().Secrets(namespace).Delete(context.TODO(), name, metav1.DeleteOptions{})
}

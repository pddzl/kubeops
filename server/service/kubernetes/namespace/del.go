package namespace

import (
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/pddzl/kubeops/server/global"
)

func (n *NamespaceService) DeleteNamespace(name string) error {
	return global.KOP_KUBERNETES.CoreV1().Namespaces().Delete(context.TODO(), name, metav1.DeleteOptions{})
}

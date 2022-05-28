package serviceAccount

import (
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/pddzl/kubeops/server/global"
)

func (sa *ServiceAccountService) DeleteServiceAccount(namespace string, name string) error {
	return global.KOP_KUBERNETES.CoreV1().ServiceAccounts(namespace).Delete(context.TODO(), name, metav1.DeleteOptions{})
}

package services

import (
	"context"

	"github.com/pddzl/kubeops/server/global"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (s *ServicesService) DeleteService(namespace string, name string) error {
	err := global.KOP_KUBERNETES.CoreV1().Services(namespace).Delete(context.TODO(), name, metav1.DeleteOptions{})

	return err
}

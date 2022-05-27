package pod

import (
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/pddzl/kubeops/server/global"
)

func (p *PodService) DeletePod(namespace string, name string) error {
	return global.KOP_KUBERNETES.CoreV1().Pods(namespace).Delete(context.TODO(), name, metav1.DeleteOptions{})
}

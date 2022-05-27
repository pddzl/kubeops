package daemonSet

import (
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/pddzl/kubeops/server/global"
)

func (d *DaemonSetService) DeleteDaemonSet(namespace string, name string) error {
	return global.KOP_KUBERNETES.AppsV1().DaemonSets(namespace).Delete(context.TODO(), name, metav1.DeleteOptions{})
}

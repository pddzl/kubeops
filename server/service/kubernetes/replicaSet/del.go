package replicaSet

import (
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/pddzl/kubeops/server/global"
)

func (r *ReplicaSetService) DeleteReplicaSet(namespace string, name string) error {
	return global.KOP_KUBERNETES.AppsV1().ReplicaSets(namespace).Delete(context.TODO(), name, metav1.DeleteOptions{})
}

package scale

import (
	"context"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/pddzl/kubeops/server/global"
)

type ScaleService struct{}

func (s *ScaleService) Scale(namespace string, name string, kind string, num int32) (count int32, err error) {
	switch kind {
	case "deployment":
		as, err := global.KOP_KUBERNETES.AppsV1().Deployments(namespace).GetScale(context.TODO(), name, metav1.GetOptions{})
		//global.KOP_KUBERNETES
		if err != nil {
			return num, fmt.Errorf("deployment getScale: %v", err)
		}

		as.Spec.Replicas = num

		_, err = global.KOP_KUBERNETES.AppsV1().Deployments(namespace).UpdateScale(context.TODO(), name, as, metav1.UpdateOptions{})
		if err != nil {
			return num, fmt.Errorf("deployment updateScale: %v", err)
		}
	case "replicaset":
		as, err := global.KOP_KUBERNETES.AppsV1().ReplicaSets(namespace).GetScale(context.TODO(), name, metav1.GetOptions{})
		if err != nil {
			return num, fmt.Errorf("replicaset getScale: %v", err)
		}

		as.Spec.Replicas = num

		_, err = global.KOP_KUBERNETES.AppsV1().ReplicaSets(namespace).UpdateScale(context.TODO(), name, as, metav1.UpdateOptions{})
		if err != nil {
			return num, fmt.Errorf("replicaset updateScale: %v", err)
		}
	}

	return num, err
}

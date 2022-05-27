package deployment

import (
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/pddzl/kubeops/server/global"
)

func (d *DeploymentService) DeleteDeployment(namespace string, name string) error {
	return global.KOP_KUBERNETES.AppsV1().Deployments(namespace).Delete(context.TODO(), name, metav1.DeleteOptions{})
}

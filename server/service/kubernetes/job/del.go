package job

import (
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/pddzl/kubeops/server/global"
)

func (j *JobService) DeleteJob(namespace string, name string) error {
	return global.KOP_KUBERNETES.BatchV1().Jobs(namespace).Delete(context.TODO(), name, metav1.DeleteOptions{})
}

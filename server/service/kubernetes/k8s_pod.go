package kubernetes

import (
	"context"
	"github.com/pddzl/kubeops/server/global"
	v1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
)

//type PodService struct{}

func getNodePods(nodeName string) (*v1.PodList, error) {
	fieldSelector, err := fields.ParseSelector("spec.nodeName=" + nodeName +
		",status.phase!=" + string(v1.PodSucceeded) +
		",status.phase!=" + string(v1.PodFailed))

	if err != nil {
		return nil, err
	}

	return global.KOP_KUBERNETES.CoreV1().Pods(v1.NamespaceAll).List(context.TODO(), metaV1.ListOptions{
		FieldSelector: fieldSelector.String(),
	})
}

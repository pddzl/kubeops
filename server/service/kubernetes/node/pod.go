package node

import (
	"context"
	"github.com/pddzl/kubeops/server/model/kubernetes/resource/common"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"

	"github.com/pddzl/kubeops/server/global"
	"github.com/pddzl/kubeops/server/model/common/request"
	"github.com/pddzl/kubeops/server/model/kubernetes/api"
)

func (n *NodeService) GetNodePods(name string, info request.PageInfo) ([]common.RelatedPod, int, error) {
	// 获取PodList原生数据
	podList, err := getNodePods(name)
	if err != nil {
		return nil, 0, err
	}

	// 处理podRaw数据
	var relatedPodList []common.RelatedPod
	for _, pod := range podList.Items {
		var relatedPod common.RelatedPod
		relatedPod.ObjectMeta = api.NewObjectMeta(pod.ObjectMeta)
		relatedPod.Status = string(pod.Status.Phase)
		relatedPod.NodeName = pod.Spec.NodeName
		// append
		relatedPodList = append(relatedPodList, relatedPod)
	}

	// 分页
	end := info.PageSize * info.Page
	offset := info.PageSize * (info.Page - 1)
	total := len(podList.Items)
	if total <= offset {
		return nil, total, nil
	}
	if total < end {
		return relatedPodList[offset:], total, nil
	} else {
		return relatedPodList[offset:end], total, nil
	}
}

func getNodePods(nodeName string) (*corev1.PodList, error) {
	fieldSelector, err := fields.ParseSelector("spec.nodeName=" + nodeName +
		",status.phase!=" + string(corev1.PodSucceeded) +
		",status.phase!=" + string(corev1.PodFailed))

	if err != nil {
		return nil, err
	}

	return global.KOP_KUBERNETES.CoreV1().Pods(corev1.NamespaceAll).List(context.TODO(), metav1.ListOptions{
		FieldSelector: fieldSelector.String(),
	})
}

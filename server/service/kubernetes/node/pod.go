package node

import (
	"context"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"

	"github.com/pddzl/kubeops/server/global"
	"github.com/pddzl/kubeops/server/model/common/request"
	"github.com/pddzl/kubeops/server/model/kubernetes/api"
	resourcePod "github.com/pddzl/kubeops/server/model/kubernetes/resource/pod"
)

func (n *NodeService) GetNodePods(name string, info request.PageInfo) ([]resourcePod.PodRefer, int, error) {
	// 获取PodList原生数据
	podListRaw, err := getNodePods(name)
	if err != nil {
		return nil, 0, err
	}

	// 分页
	var podList corev1.PodList
	end := info.PageSize * info.Page
	offset := info.PageSize * (info.Page - 1)
	total := len(podListRaw.Items)
	if total <= offset {
		return nil, total, nil
	}
	if total < end {
		podList.Items = podListRaw.Items[offset:]
	} else {
		podList.Items = podListRaw.Items[offset:end]
	}

	// 处理podRaw数据
	var podBriefList []resourcePod.PodRefer
	for _, podRaw := range podList.Items {
		var pod resourcePod.PodRefer
		pod.ObjectMeta = api.NewObjectMeta(podRaw.ObjectMeta)
		pod.Node = podRaw.Spec.NodeName
		pod.Status = string(podRaw.Status.Phase)
		// append
		podBriefList = append(podBriefList, pod)
	}

	return podBriefList, total, nil
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

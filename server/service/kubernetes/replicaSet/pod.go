package replicaSet

import (
	"context"
	"k8s.io/api/apps/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"

	"github.com/pddzl/kubeops/server/global"
	"github.com/pddzl/kubeops/server/model/common/request"
	"github.com/pddzl/kubeops/server/model/kubernetes/api"
	"github.com/pddzl/kubeops/server/model/kubernetes/resource/common"
	"github.com/pddzl/kubeops/server/model/kubernetes/resource/replicaSet"
)

func getReplicaSetPodInfo(replicaSet *v1.ReplicaSet) (*common.PodInfo, error) {
	labelSelector := labels.SelectorFromSet(replicaSet.Spec.Selector.MatchLabels)
	channels := &common.ResourceChannels{
		PodList: common.GetPodListChannelWithOptions(common.NewSameNamespaceQuery(replicaSet.Namespace),
			metaV1.ListOptions{
				LabelSelector: labelSelector.String(),
			}, 1),
	}

	podList := <-channels.PodList.List
	if err := <-channels.PodList.Error; err != nil {
		return nil, err
	}

	filterPod := common.FilterPodsByControllerRef(replicaSet, podList.Items)
	podInfo := common.GetPodInfo(replicaSet.Status.Replicas, replicaSet.Spec.Replicas, filterPod)
	return &podInfo, nil
}

func (r *ReplicaSetService) GetReplicaSetPods(namespace string, replicaSetName string, info request.PageInfo) ([]replicaSet.Pod, int, error) {
	end := info.PageSize * info.Page
	offset := info.PageSize * (info.Page - 1)
	var replicaSetPods []replicaSet.Pod

	rs, err := global.KOP_KUBERNETES.AppsV1().ReplicaSets(namespace).Get(context.TODO(), replicaSetName, metaV1.GetOptions{})
	if err != nil {
		return nil, 0, err
	}

	labelSelector := labels.SelectorFromSet(rs.Spec.Selector.MatchLabels)
	channels := &common.ResourceChannels{
		PodList: common.GetPodListChannelWithOptions(common.NewSameNamespaceQuery(namespace), metaV1.ListOptions{LabelSelector: labelSelector.String()}, 1),
	}

	podList := <-channels.PodList.List
	if err := <-channels.PodList.Error; err != nil {
		return nil, 0, err
	}

	// 处理replicaSet Pods
	for _, pod := range podList.Items {
		replicaSetPod := replicaSet.Pod{}
		replicaSetPod.ObjectMeta = api.NewObjectMeta(pod.ObjectMeta)
		replicaSetPod.TypeMeta = api.NewTypeMeta(api.ResourceKindPod)
		replicaSetPod.Status = string(pod.Status.Phase)
		replicaSetPod.NodeName = pod.Spec.NodeName
		// append
		replicaSetPods = append(replicaSetPods, replicaSetPod)
	}

	// 分页
	total := len(podList.Items)
	if total <= offset {
		return nil, total, nil
	}
	if total < end {
		return replicaSetPods[offset:], total, nil
	} else {
		return replicaSetPods[offset:end], total, nil
	}
}

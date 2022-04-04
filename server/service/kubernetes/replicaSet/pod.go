package replicaSet

import (
	"github.com/pddzl/kubeops/server/model/kubernetes/resource/common"
	"k8s.io/api/apps/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
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

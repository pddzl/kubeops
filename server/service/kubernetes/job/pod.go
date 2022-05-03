package job

import (
	"context"
	"github.com/pddzl/kubeops/server/global"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"

	"github.com/pddzl/kubeops/server/model/common/request"
	"github.com/pddzl/kubeops/server/model/kubernetes/api"
	"github.com/pddzl/kubeops/server/model/kubernetes/resource/common"
)

func (j *JobService) GetJobPods(namespace string, name string, info request.PageInfo) ([]common.RelatedPod, int, error) {
	// 获取job
	job, err := global.KOP_KUBERNETES.BatchV1().Jobs(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		return nil, 0, err
	}

	// 获取job的selector
	selector := labels.SelectorFromSet(job.Spec.Selector.MatchLabels)
	options := metav1.ListOptions{LabelSelector: selector.String()}

	// 获取job关联的pod
	podList, err := global.KOP_KUBERNETES.CoreV1().Pods(namespace).List(context.TODO(), options)
	if err != nil {
		return nil, 0, err
	}

	// 处理job Pods
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

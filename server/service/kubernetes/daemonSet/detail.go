package daemonSet

import (
	"context"
	"github.com/pddzl/kubeops/server/global"
	"github.com/pddzl/kubeops/server/model/kubernetes/api"
	resourceDaemonSet "github.com/pddzl/kubeops/server/model/kubernetes/resource/daemonSet"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (d *DaemonSetService) GetDaemonSetDetail(namespace string, name string) (*resourceDaemonSet.DeploymentDetail, error) {
	// 获取daemonSet原生数据
	daemonSet, err := global.KOP_KUBERNETES.AppsV1().DaemonSets(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}

	// 处理daemonSet数据
	var daemonSetDetail resourceDaemonSet.DeploymentDetail
	// metadata
	daemonSetDetail.ObjectMeta = api.NewObjectMeta(daemonSet.ObjectMeta)
	// spec
	daemonSetDetail.Spec.Selector = daemonSet.Spec.Selector.MatchLabels
	daemonSetDetail.Spec.UpdateStrategy.Type = string(daemonSet.Spec.UpdateStrategy.Type)
	daemonSetDetail.Spec.UpdateStrategy.RollingUpdate.MaxSurge = daemonSet.Spec.UpdateStrategy.RollingUpdate.MaxSurge
	daemonSetDetail.Spec.UpdateStrategy.RollingUpdate.MaxUnavailable = daemonSet.Spec.UpdateStrategy.RollingUpdate.MaxUnavailable
	daemonSetDetail.Spec.MinReadySeconds = daemonSet.Spec.MinReadySeconds
	daemonSetDetail.Spec.RevisionHistoryLimit = daemonSet.Spec.RevisionHistoryLimit
	// status
	daemonSetDetail.Status.CurrentNumberScheduled = daemonSet.Status.CurrentNumberScheduled
	daemonSetDetail.Status.NumberMisscheduled = daemonSet.Status.NumberMisscheduled
	daemonSetDetail.Status.DesiredNumberScheduled = daemonSet.Status.DesiredNumberScheduled
	daemonSetDetail.Status.NumberReady = daemonSet.Status.NumberReady
	daemonSetDetail.Status.UpdatedNumberScheduled = daemonSet.Status.UpdatedNumberScheduled
	daemonSetDetail.Status.NumberAvailable = daemonSet.Status.NumberAvailable
	daemonSetDetail.Status.NumberUnavailable = daemonSet.Status.NumberUnavailable
	// status -> conditions
	for _, condition := range daemonSet.Status.Conditions {
		var daemonSetCondition resourceDaemonSet.DaemonSetCondition
		daemonSetCondition.Type = string(condition.Type)
		daemonSetCondition.Status = string(condition.Status)
		daemonSetCondition.LastTransitionTime = condition.LastTransitionTime
		daemonSetCondition.Reason = condition.Reason
		daemonSetCondition.Message = condition.Message
		// append
		daemonSetDetail.Status.Conditions = append(daemonSetDetail.Status.Conditions, daemonSetCondition)
	}

	return &daemonSetDetail, nil
}

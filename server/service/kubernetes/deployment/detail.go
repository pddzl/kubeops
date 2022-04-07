package deployment

import (
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/pddzl/kubeops/server/global"
	"github.com/pddzl/kubeops/server/model/kubernetes/api"
	resourceDeployment "github.com/pddzl/kubeops/server/model/kubernetes/resource/deployment"
)

func (d *DeploymentService) GetDeploymentDetail(namespace string, deploymentName string) (*resourceDeployment.DeploymentDetail, error) {
	deployment, err := global.KOP_KUBERNETES.AppsV1().Deployments(namespace).Get(context.TODO(), deploymentName, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}

	// 处理deployment数据
	var deploymentDetail resourceDeployment.DeploymentDetail
	// metadata
	deploymentDetail.ObjectMeta = api.NewObjectMeta(deployment.ObjectMeta)
	// spec
	deploymentDetail.Spec.Replicas = deployment.Spec.Replicas
	deploymentDetail.Spec.Selector = deployment.Spec.Selector.MatchLabels
	deploymentDetail.Spec.Strategy.DeploymentStrategyType = string(deployment.Spec.Strategy.Type)
	deploymentDetail.Spec.Strategy.RollingUpdateDeployment.MaxUnavailable = deployment.Spec.Strategy.RollingUpdate.MaxUnavailable
	deploymentDetail.Spec.Strategy.RollingUpdateDeployment.MaxSurge = deployment.Spec.Strategy.RollingUpdate.MaxSurge
	deploymentDetail.Spec.MinReadySeconds = deployment.Spec.MinReadySeconds
	deploymentDetail.Spec.RevisionHistoryLimit = deployment.Spec.RevisionHistoryLimit
	// status
	deploymentDetail.Status.Replicas = deployment.Status.Replicas
	deploymentDetail.Status.UpdatedReplicas = deployment.Status.UpdatedReplicas
	deploymentDetail.Status.ReadyReplicas = deployment.Status.ReadyReplicas
	deploymentDetail.Status.AvailableReplicas = deployment.Status.AvailableReplicas
	deploymentDetail.Status.UnavailableReplicas = deployment.Status.UnavailableReplicas
	// status -> condition
	for _, dsc := range deployment.Status.Conditions {
		var condition resourceDeployment.DeploymentCondition
		condition.Type = string(dsc.Type)
		condition.Status = string(dsc.Status)
		condition.LastTransitionTime = dsc.LastTransitionTime
		condition.LastUpdateTime = dsc.LastUpdateTime
		condition.Reason = dsc.Reason
		condition.Message = dsc.Message
		// append
		deploymentDetail.Status.Conditions = append(deploymentDetail.Status.Conditions, condition)
	}

	return &deploymentDetail, nil
}

package deployment

import (
	"context"
	"fmt"
	v1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/pddzl/kubeops/server/global"
	"github.com/pddzl/kubeops/server/model/common/request"
	"github.com/pddzl/kubeops/server/model/kubernetes/resource/deployment"
)

func (d *DeploymentService) GetDeploymentList(namespace string, info request.PageInfo) ([]deployment.DeploymentBrief, int, error) {
	end := info.PageSize * info.Page
	offset := info.PageSize * (info.Page - 1)
	var deploymentBriefList []deployment.DeploymentBrief
	var deploymentList v1.DeploymentList

	deployments, err := global.KOP_KUBERNETES.AppsV1().Deployments(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, 0, err
	}

	// 分页
	total := len(deployments.Items)
	if total <= offset {
		return nil, total, nil
	}
	if total < end {
		deploymentList.Items = deployments.Items[offset:]
	} else {
		deploymentList.Items = deployments.Items[offset:end]
	}

	// 处理deployments数据
	for _, dm := range deployments.Items {
		var deploymentBrief deployment.DeploymentBrief
		deploymentBrief.Name = dm.Name
		deploymentBrief.NameSpace = dm.Namespace
		deploymentBrief.Pods = fmt.Sprintf("%d / %d", dm.Status.AvailableReplicas, dm.Status.Replicas)
		deploymentBrief.CreationTimestamp = dm.CreationTimestamp
		// append
		deploymentBriefList = append(deploymentBriefList, deploymentBrief)
	}

	return deploymentBriefList, total, nil
}

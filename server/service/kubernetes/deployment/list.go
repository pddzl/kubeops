package deployment

import (
	"context"
	"fmt"
	appsV1 "k8s.io/api/apps/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/pddzl/kubeops/server/global"
	"github.com/pddzl/kubeops/server/model/common/request"
	resourceDeployment "github.com/pddzl/kubeops/server/model/kubernetes/resource/deployment"
)

func (d *DeploymentService) GetDeploymentList(namespace string, info request.PageInfo) ([]resourceDeployment.DeploymentBrief, int, error) {
	// 获取deployment list
	list, err := global.KOP_KUBERNETES.AppsV1().Deployments(namespace).List(context.TODO(), metaV1.ListOptions{})
	if err != nil {
		return nil, 0, err
	}

	var deploymentList appsV1.DeploymentList
	// 分页
	end := info.PageSize * info.Page
	offset := info.PageSize * (info.Page - 1)
	total := len(list.Items)
	if total <= offset {
		return nil, total, nil
	}
	if total < end {
		deploymentList.Items = list.Items[offset:]
	} else {
		deploymentList.Items = list.Items[offset:end]
	}

	var deploymentBriefList []resourceDeployment.DeploymentBrief
	// 处理list数据
	for _, dm := range list.Items {
		var deploymentBrief resourceDeployment.DeploymentBrief
		deploymentBrief.Name = dm.Name
		deploymentBrief.NameSpace = dm.Namespace
		deploymentBrief.Pods = fmt.Sprintf("%d / %d", dm.Status.AvailableReplicas, dm.Status.Replicas)
		deploymentBrief.CreationTimestamp = dm.CreationTimestamp
		// append
		deploymentBriefList = append(deploymentBriefList, deploymentBrief)
	}

	return deploymentBriefList, total, nil
}

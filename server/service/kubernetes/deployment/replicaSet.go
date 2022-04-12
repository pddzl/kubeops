package deployment

import (
	"context"
	"fmt"
	"github.com/pddzl/kubeops/server/global"
	deployment2 "github.com/pddzl/kubeops/server/model/kubernetes/resource/deployment"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (d *DeploymentService) GetNewReplicaSet(namespace string, deploymentName string) (*deployment2.NewReplicaSet, error) {
	deployment, err := global.KOP_KUBERNETES.AppsV1().Deployments(namespace).Get(context.TODO(), deploymentName, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}

	// 获取deployment的selector
	selector, err := metav1.LabelSelectorAsSelector(deployment.Spec.Selector)
	if err != nil {
		return nil, err
	}
	options := metav1.ListOptions{LabelSelector: selector.String()}

	// 获取replicaSet
	replicaSetList, err := global.KOP_KUBERNETES.AppsV1().ReplicaSets(namespace).List(context.TODO(), options)
	if err != nil {
		return nil, err
	}

	var newReplicaSet deployment2.NewReplicaSet
	newReplicaSet.Name = replicaSetList.Items[0].Name
	newReplicaSet.NameSpace = replicaSetList.Items[0].Namespace
	newReplicaSet.Labels = replicaSetList.Items[0].Labels
	newReplicaSet.CreationTimestamp = replicaSetList.Items[0].CreationTimestamp
	newReplicaSet.Replicas = fmt.Sprintf("%d / %d", replicaSetList.Items[0].Status.AvailableReplicas, replicaSetList.Items[0].Status.Replicas)

	return &newReplicaSet, nil
}

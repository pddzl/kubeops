package node

import (
	"context"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"regexp"

	"github.com/pddzl/kubeops/server/global"
	"github.com/pddzl/kubeops/server/model/common/request"
	resourceNode "github.com/pddzl/kubeops/server/model/kubernetes/resource/node"
)

func (n *NodeService) GetNodeList(info *request.PageInfo) ([]resourceNode.NodeBrief, int, error) {
	// 获取node list
	list, err := global.KOP_KUBERNETES.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, 0, err
	}

	var nodeList corev1.NodeList
	// 分页
	end := info.PageSize * info.Page
	offset := info.PageSize * (info.Page - 1)
	total := len(list.Items)
	if total <= offset {
		return nil, total, nil
	}
	if total < end {
		nodeList.Items = list.Items[offset:]
	} else {
		nodeList.Items = list.Items[offset:end]
	}

	var nodeBriefList []resourceNode.NodeBrief
	// 处理list数据
	roleRe, _ := regexp.Compile("node-role.kubernetes.io/(.*)")
	for _, node := range nodeList.Items {
		var nodeBrief resourceNode.NodeBrief
		nodeBrief.Name = node.Name
		nodeBrief.InternalIP = node.Status.Addresses[0].Address
		for label := range node.ObjectMeta.Labels {
			role := roleRe.FindStringSubmatch(label)
			if len(role) == 2 {
				nodeBrief.Roles = append(nodeBrief.Roles, role[1])
			}
		}
		for _, condition := range node.Status.Conditions {
			if condition.Type == "Ready" {
				nodeBrief.Status = string(condition.Status)
			}
		}
		nodeBrief.Cpu = node.Status.Capacity.Cpu().String()
		nodeBrief.Memory = node.Status.Capacity.Memory().String()
		nodeBrief.CreationTimestamp = node.CreationTimestamp
		nodeBriefList = append(nodeBriefList, nodeBrief)
	}

	return nodeBriefList, total, nil
}

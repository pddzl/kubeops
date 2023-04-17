package k8s

import (
	"context"
	"github.com/pddzl/kubeops/server/global"
	"github.com/pddzl/kubeops/server/model/common/request"
	k8sResp "github.com/pddzl/kubeops/server/model/k8s/response"
	coreV1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"regexp"
)

type NodeService struct{}

func (n *NodeService) GetNodes(info request.PageInfo) ([]k8sResp.NodeBrief, int, error) {
	// 获取node list
	list, err := global.KOP_K8S_Client.CoreV1().Nodes().List(context.TODO(), metaV1.ListOptions{})
	if err != nil {
		return nil, 0, err
	}

	var nodeList coreV1.NodeList
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

	var nodeBriefList []k8sResp.NodeBrief
	// 处理list数据
	roleRe, _ := regexp.Compile("node-role.kubernetes.io/(.*)")
	for _, node := range nodeList.Items {
		var nodeBrief k8sResp.NodeBrief
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
		nodeBrief.CreationTimestamp = node.CreationTimestamp.Time
		nodeBriefList = append(nodeBriefList, nodeBrief)
	}

	return nodeBriefList, total, nil
}

package kubernetes

import (
	"context"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/pddzl/kubeops/server/global"
	"github.com/pddzl/kubeops/server/model/common/request"
)

type NodeService struct{}

func (n *NodeService) GetNodeList(info request.PageInfo) (list interface{}, total int, err error) {
	end := info.PageSize * info.Page
	offset := info.PageSize * (info.Page - 1)
	nodes, err := global.KOP_KUBERNETES.CoreV1().Nodes().List(context.TODO(), metaV1.ListOptions{})
	if err != nil {
		return nil, 0, err
	}
	total = len(nodes.Items)
	if total <= offset {
		return nil, total, nil
	}
	if total < end {
		list = nodes.Items[offset:]
	} else {
		list = nodes.Items[offset : end-1]
	}
	return list, total, nil
}

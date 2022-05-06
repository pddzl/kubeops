package namespace

import (
	"context"
	coreV1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/pddzl/kubeops/server/global"
	"github.com/pddzl/kubeops/server/model/common/request"
	resourceNamespace "github.com/pddzl/kubeops/server/model/kubernetes/resource/namespace"
)

func (n *NamespaceService) GetNamespaceList(info request.PageInfo) ([]resourceNamespace.NamespaceBrief, int, error) {
	// 获取namespace list
	list, err := global.KOP_KUBERNETES.CoreV1().Namespaces().List(context.TODO(), metaV1.ListOptions{})
	if err != nil {
		return nil, 0, err
	}

	var namespaceList coreV1.NamespaceList
	// 分页
	end := info.PageSize * info.Page
	offset := info.PageSize * (info.Page - 1)
	total := len(list.Items)
	if total <= offset {
		return nil, total, nil
	}
	if total < end {
		namespaceList.Items = list.Items[offset:]
	} else {
		namespaceList.Items = list.Items[offset:end]
	}

	var namespaceBriefList []resourceNamespace.NamespaceBrief
	// 处理namespace原始数据
	for _, ns := range namespaceList.Items {
		var namespaceBrief resourceNamespace.NamespaceBrief
		namespaceBrief.Name = ns.Name
		namespaceBrief.Labels = ns.Labels
		namespaceBrief.CreationTimestamp = ns.CreationTimestamp
		namespaceBrief.Status = string(ns.Status.Phase)
		// append
		namespaceBriefList = append(namespaceBriefList, namespaceBrief)
	}

	return namespaceBriefList, total, nil
}

func (n *NamespaceService) GetNamespaceOnlyName() ([]string, error) {
	// 获取namespace list
	list, err := global.KOP_KUBERNETES.CoreV1().Namespaces().List(context.TODO(), metaV1.ListOptions{})
	if err != nil {
		return nil, err
	}

	var nameList []string
	// 处理namespace原始数据
	for _, ns := range list.Items {
		// append
		nameList = append(nameList, ns.Name)
	}
	return nameList, nil
}

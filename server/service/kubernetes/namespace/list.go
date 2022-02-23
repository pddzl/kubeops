package namespace

import (
	"context"
	v1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/pddzl/kubeops/server/global"
	"github.com/pddzl/kubeops/server/model/common/request"
	"github.com/pddzl/kubeops/server/model/kubernetes"
)

func (n *NamespaceService) GetNamespaceList(info request.PageInfo) (list interface{}, total int, err error) {
	end := info.PageSize * info.Page
	offset := info.PageSize * (info.Page - 1)
	var namespaceBriefList []kubernetes.NameSpaceBrief
	var namespaceList v1.NamespaceList

	opts := metaV1.ListOptions{}
	namespaces, err := global.KOP_KUBERNETES.CoreV1().Namespaces().List(context.TODO(), opts)
	if err != nil {
		return nil, 0, err
	}

	// 分页
	total = len(namespaces.Items)
	if total <= offset {
		return nil, total, nil
	}
	if total < end {
		namespaceList.Items = namespaces.Items[offset:]
	} else {
		namespaceList.Items = namespaces.Items[offset:end]
	}

	// 处理namespace原始数据
	for _, ns := range namespaceList.Items {
		var namespaceBrief kubernetes.NameSpaceBrief
		namespaceBrief.Name = ns.Name
		namespaceBrief.Labels = ns.Labels
		namespaceBrief.CreationTimestamp = ns.CreationTimestamp
		namespaceBrief.Status = string(ns.Status.Phase)
		// append
		namespaceBriefList = append(namespaceBriefList, namespaceBrief)
	}

	return namespaceBriefList, total, nil
}

func (n *NamespaceService) GetNamespaceOnlyName() (list interface{}, err error) {
	var nameList []string
	opts := metaV1.ListOptions{}
	namespaces, err := global.KOP_KUBERNETES.CoreV1().Namespaces().List(context.TODO(), opts)
	if err != nil {
		return nil, err
	}

	// 处理namespace原始数据
	for _, ns := range namespaces.Items {
		// append
		nameList = append(nameList, ns.Name)
	}
	return nameList, nil
}

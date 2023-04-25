package k8s

import (
	"context"
	"github.com/pddzl/kubeops/server/global"
	"github.com/pddzl/kubeops/server/model/common/request"
	k8sResponse "github.com/pddzl/kubeops/server/model/k8s/response"
	coreV1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type NamespaceService struct{}

// GetNamespaceList 获取namepsace list
func (nss *NamespaceService) GetNamespaceList(info request.PageInfo) ([]k8sResponse.NamespaceBrief, int, error) {
	// 获取namespace list
	list, err := global.KOP_K8S_Client.CoreV1().Namespaces().List(context.TODO(), metaV1.ListOptions{})
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

	var namespaceBriefList []k8sResponse.NamespaceBrief
	// 处理namespace原始数据
	for _, ns := range namespaceList.Items {
		var namespaceBrief k8sResponse.NamespaceBrief
		namespaceBrief.Name = ns.Name
		namespaceBrief.Labels = ns.Labels
		namespaceBrief.CreationTimestamp = ns.CreationTimestamp
		namespaceBrief.Status = string(ns.Status.Phase)
		// append
		namespaceBriefList = append(namespaceBriefList, namespaceBrief)
	}

	return namespaceBriefList, total, nil
}

// GetNamespaceDetail 获取某个namespace详情
func (nss *NamespaceService) GetNamespaceDetail(name string) (*resourceNamespace.NamespaceDetail, error) {
	// 获取namespace原生数据
	namespace, err := global.KOP_K8S_Client.CoreV1().Namespaces().Get(context.TODO(), name, metaV1.GetOptions{})
	if err != nil {
		return nil, err
	}

	// 处理namespace数据
	var namespaceDetail resourceNamespace.NamespaceDetail
	// metadata
	namespaceDetail.Metadata = api.NewObjectMeta(namespace.ObjectMeta)
	// status
	namespaceDetail.Status = string(namespace.Status.Phase)

	// resourceQuotaList
	resourceQuotaList, err := getResourceQuotas(name)
	if err != nil {
		return nil, err
	}
	namespaceDetail.ResourceQuotaList = resourceQuotaList

	// resourceLimits
	resourceLimits, err := getLimitRanges(*namespace)
	if err != nil {
		return nil, err
	}
	namespaceDetail.ResourceLimits = resourceLimits

	return &namespaceDetail, nil
}

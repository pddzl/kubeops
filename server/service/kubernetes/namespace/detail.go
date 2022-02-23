package namespace

import (
	"context"
	v1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/pddzl/kubeops/server/global"
	"github.com/pddzl/kubeops/server/model/kubernetes"
	"github.com/pddzl/kubeops/server/model/kubernetes/api"
	"github.com/pddzl/kubeops/server/model/kubernetes/resource"
)

func (n *NamespaceService) GetNamespaceDetail(name string) (detail interface{}, err error) {
	var namespaceDetail kubernetes.NameSpaceDetail
	namespace, err := global.KOP_KUBERNETES.CoreV1().Namespaces().Get(context.TODO(), name, metaV1.GetOptions{})
	if err != nil {
		return nil, err
	}
	// metadata
	namespaceDetail.Metadata.Name = namespace.Name
	namespaceDetail.Metadata.Labels = namespace.Labels
	namespaceDetail.Metadata.CreationTimestamp = namespace.CreationTimestamp
	namespaceDetail.Metadata.Uid = string(namespace.UID)

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
	return namespaceDetail, nil
}

func getResourceQuotas(namespace string) ([]resource.ResourceQuotaDetail, error) {
	list, err := global.KOP_KUBERNETES.CoreV1().ResourceQuotas(namespace).List(context.TODO(), api.ListEverything)

	result := make([]resource.ResourceQuotaDetail, 0)

	for _, item := range list.Items {
		detail := resource.ToResourceQuotaDetail(&item)
		result = append(result, *detail)
	}

	return result, err
}

func getLimitRanges(namespace v1.Namespace) ([]resource.LimitRangeItem, error) {
	list, err := global.KOP_KUBERNETES.CoreV1().LimitRanges(namespace.Name).List(context.TODO(), api.ListEverything)
	if err != nil {
		return nil, err
	}

	resourceLimits := make([]resource.LimitRangeItem, 0)
	for _, item := range list.Items {
		list := resource.ToLimitRanges(&item)
		resourceLimits = append(resourceLimits, list...)
	}

	return resourceLimits, nil
}

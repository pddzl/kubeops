package namespace

import (
	"context"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/pddzl/kubeops/server/global"
	"github.com/pddzl/kubeops/server/model/kubernetes/api"
	"github.com/pddzl/kubeops/server/model/kubernetes/resource"
	resourceNamespace "github.com/pddzl/kubeops/server/model/kubernetes/resource/namespace"
)

func (n *NamespaceService) GetNamespaceDetail(name string) (*resourceNamespace.NameSpaceDetail, error) {
	// 获取namespace原生数据
	namespace, err := global.KOP_KUBERNETES.CoreV1().Namespaces().Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}

	// 处理namespace数据
	var namespaceDetail resourceNamespace.NameSpaceDetail
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

func getResourceQuotas(namespace string) ([]resource.ResourceQuotaDetail, error) {
	list, err := global.KOP_KUBERNETES.CoreV1().ResourceQuotas(namespace).List(context.TODO(), api.ListEverything)

	result := make([]resource.ResourceQuotaDetail, 0)

	for _, item := range list.Items {
		detail := resource.ToResourceQuotaDetail(&item)
		result = append(result, *detail)
	}

	return result, err
}

func getLimitRanges(namespace corev1.Namespace) ([]resource.LimitRangeDetail, error) {
	list, err := global.KOP_KUBERNETES.CoreV1().LimitRanges(namespace.Name).List(context.TODO(), api.ListEverything)
	if err != nil {
		return nil, err
	}

	resourceLimits := make([]resource.LimitRangeDetail, 0)
	for _, item := range list.Items {
		list := resource.ToLimitRanges(&item)
		resourceLimits = append(resourceLimits, list)
	}

	return resourceLimits, nil
}

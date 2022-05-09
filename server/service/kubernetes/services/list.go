package services

import (
	"context"
	coreV1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/pddzl/kubeops/server/global"
	"github.com/pddzl/kubeops/server/model/common/request"
	resourceServices "github.com/pddzl/kubeops/server/model/kubernetes/resource/services"
)

func (s *ServicesService) GetServicesList(namespace string, info request.PageInfo) ([]resourceServices.ServicesBrief, int, error) {
	// 获取service list
	list, err := global.KOP_KUBERNETES.CoreV1().Services(namespace).List(context.TODO(), metaV1.ListOptions{})
	if err != nil {
		return nil, 0, err
	}

	var serviceList coreV1.ServiceList
	// 分页
	end := info.PageSize * info.Page
	offset := info.PageSize * (info.Page - 1)
	total := len(list.Items)
	if total <= offset {
		return nil, total, nil
	}
	if total < end {
		serviceList.Items = list.Items[offset:]
	} else {
		serviceList.Items = list.Items[offset:end]
	}

	var serviceBriefList []resourceServices.ServicesBrief
	// 处理list数据
	for _, slr := range serviceList.Items {
		var servicesBrief resourceServices.ServicesBrief
		servicesBrief.Name = slr.Name
		servicesBrief.NameSpace = slr.Namespace
		servicesBrief.ClusterIP = slr.Spec.ClusterIP
		servicesBrief.Type = string(slr.Spec.Type)
		servicesBrief.CreationTimestamp = slr.CreationTimestamp
		// append
		serviceBriefList = append(serviceBriefList, servicesBrief)
	}

	return serviceBriefList, total, nil
}

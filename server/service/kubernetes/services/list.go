package services

import (
	"context"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/pddzl/kubeops/server/global"
	"github.com/pddzl/kubeops/server/model/common/request"
	"github.com/pddzl/kubeops/server/model/kubernetes/resource/services"
)

func (s *ServicesService) GetServicesList(namespace string, info request.PageInfo) ([]services.ServicesBrief, int, error) {
	// 获取services原生数据
	servicesListRaw, err := global.KOP_KUBERNETES.CoreV1().Services(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, 0, err
	}

	end := info.PageSize * info.Page
	offset := info.PageSize * (info.Page - 1)
	var replicaSetBriefList []services.ServicesBrief
	var replicaSetList v1.ServiceList

	// 分页
	total := len(servicesListRaw.Items)
	if total <= offset {
		return nil, total, nil
	}
	if total < end {
		replicaSetList.Items = servicesListRaw.Items[offset:]
	} else {
		replicaSetList.Items = servicesListRaw.Items[offset:end]
	}

	// 处理servicesListRaw数据
	for _, slr := range servicesListRaw.Items {
		var servicesBrief services.ServicesBrief
		servicesBrief.Name = slr.Name
		servicesBrief.NameSpace = slr.Namespace
		servicesBrief.ClusterIP = slr.Spec.ClusterIP
		servicesBrief.Type = string(slr.Spec.Type)
		servicesBrief.CreationTimestamp = slr.CreationTimestamp
		// append
		replicaSetBriefList = append(replicaSetBriefList, servicesBrief)
	}

	return replicaSetBriefList, total, nil
}

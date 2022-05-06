package serviceAccount

import (
	"context"
	coreV1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/pddzl/kubeops/server/global"
	"github.com/pddzl/kubeops/server/model/common/request"
	resourceServiceAccount "github.com/pddzl/kubeops/server/model/kubernetes/resource/serviceAccount"
)

func (sa *ServiceAccountService) GetServiceAccountList(namespace string, info request.PageInfo) ([]resourceServiceAccount.ServiceAccountBrief, int, error) {
	// 获取serviceAccount List
	list, err := global.KOP_KUBERNETES.CoreV1().ServiceAccounts(namespace).List(context.TODO(), metaV1.ListOptions{})
	if err != nil {
		return nil, 0, err
	}

	var serviceAccountList coreV1.ServiceAccountList
	// 分页
	end := info.PageSize * info.Page
	offset := info.PageSize * (info.Page - 1)
	total := len(list.Items)
	if total <= offset {
		return nil, total, nil
	}
	if total < end {
		serviceAccountList.Items = list.Items[offset:]
	} else {
		serviceAccountList.Items = list.Items[offset:end]
	}

	// 处理list数据
	var serviceAccountBriefList []resourceServiceAccount.ServiceAccountBrief
	for _, serviceAccount := range serviceAccountList.Items {
		var serviceAccountBrief resourceServiceAccount.ServiceAccountBrief
		serviceAccountBrief.Name = serviceAccount.Name
		serviceAccountBrief.Namespace = serviceAccount.Namespace
		serviceAccountBrief.CreationTimestamp = serviceAccount.CreationTimestamp
		// append
		serviceAccountBriefList = append(serviceAccountBriefList, serviceAccountBrief)
	}

	return serviceAccountBriefList, total, nil
}

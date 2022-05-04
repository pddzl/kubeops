package serviceAccount

import (
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/pddzl/kubeops/server/global"
	"github.com/pddzl/kubeops/server/model/common/request"
	resourceServiceAccount "github.com/pddzl/kubeops/server/model/kubernetes/resource/serviceAccount"
)

func (sa *ServiceAccountService) GetServiceAccountList(namespace string, info request.PageInfo) ([]resourceServiceAccount.ServiceAccountBrief, int, error) {
	// 获取serviceAccount List
	serviceAccountList, err := global.KOP_KUBERNETES.CoreV1().ServiceAccounts(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, 0, err
	}

	// 处理serviceAccount List数据
	var serviceAccountBriefList []resourceServiceAccount.ServiceAccountBrief
	for _, serviceAccount := range serviceAccountList.Items {
		var serviceAccountBrief resourceServiceAccount.ServiceAccountBrief
		serviceAccountBrief.Name = serviceAccount.Name
		serviceAccountBrief.Namespace = serviceAccount.Namespace
		serviceAccountBrief.CreationTimestamp = serviceAccount.CreationTimestamp
		// append
		serviceAccountBriefList = append(serviceAccountBriefList, serviceAccountBrief)
	}

	// 分页
	end := info.PageSize * info.Page
	offset := info.PageSize * (info.Page - 1)
	total := len(serviceAccountBriefList)
	if total <= offset {
		return nil, total, nil
	}
	if total < end {
		return serviceAccountBriefList[offset:], total, nil
	} else {
		return serviceAccountBriefList[offset:end], total, nil
	}
}

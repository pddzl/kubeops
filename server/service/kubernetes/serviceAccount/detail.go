package serviceAccount

import (
	"context"
	"github.com/pddzl/kubeops/server/global"
	coreV1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (sa *ServiceAccountService) GetServiceAccountDetail(namespace string, name string) (*coreV1.ServiceAccount, error) {
	// 获取原生数据
	detail, err := global.KOP_KUBERNETES.CoreV1().ServiceAccounts(namespace).Get(context.TODO(), name, metaV1.GetOptions{})
	if err != nil {
		return nil, err
	}

	return detail, nil
}

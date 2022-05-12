package secret

import (
	"context"
	coreV1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/pddzl/kubeops/server/global"
	"github.com/pddzl/kubeops/server/model/common/request"
	resourceSecret "github.com/pddzl/kubeops/server/model/kubernetes/resource/secret"
)

func (s *SecretService) GetSecretList(namespace string, info request.PageInfo) ([]resourceSecret.SecretBrief, int, error) {
	// 获取secret list
	list, err := global.KOP_KUBERNETES.CoreV1().Secrets(namespace).List(context.TODO(), metaV1.ListOptions{})
	if err != nil {
		return nil, 0, err
	}

	var secretList coreV1.SecretList
	// 分页
	end := info.PageSize * info.Page
	offset := info.PageSize * (info.Page - 1)
	total := len(list.Items)
	if total <= offset {
		return nil, total, nil
	}
	if total < end {
		secretList.Items = list.Items[offset:]
	} else {
		secretList.Items = list.Items[offset:end]
	}

	var secretBriefList []resourceSecret.SecretBrief
	for _, secret := range secretList.Items {
		var secretBrief resourceSecret.SecretBrief
		secretBrief.Name = secret.Name
		secretBrief.Namespace = secret.Namespace
		secretBrief.Type = string(secret.Type)
		secretBrief.CreationTimestamp = secret.CreationTimestamp
		// append
		secretBriefList = append(secretBriefList, secretBrief)
	}

	return secretBriefList, total, nil
}

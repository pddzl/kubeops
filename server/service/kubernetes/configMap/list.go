package configMap

import (
	"context"
	coreV1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/pddzl/kubeops/server/global"
	"github.com/pddzl/kubeops/server/model/common/request"
	resourceConfigMap "github.com/pddzl/kubeops/server/model/kubernetes/resource/configMap"
)

func (cm *ConfigMapService) GetConfigMapList(namespace string, info request.PageInfo) ([]resourceConfigMap.ConfigMapBrief, int, error) {
	// 获取configMap list
	list, err := global.KOP_KUBERNETES.CoreV1().ConfigMaps(namespace).List(context.TODO(), metaV1.ListOptions{})
	if err != nil {
		return nil, 0, err
	}

	var configMapList coreV1.ConfigMapList
	// 分页
	end := info.PageSize * info.Page
	offset := info.PageSize * (info.Page - 1)
	total := len(list.Items)
	if total <= offset {
		return nil, total, nil
	}
	if total < end {
		configMapList.Items = list.Items[offset:]
	} else {
		configMapList.Items = list.Items[offset:end]
	}

	var configMapBriefList []resourceConfigMap.ConfigMapBrief
	// 处理configMap数据
	for _, cm := range configMapList.Items {
		var configMapBrief resourceConfigMap.ConfigMapBrief
		configMapBrief.Name = cm.Name
		configMapBrief.Namespace = cm.Namespace
		configMapBrief.CreationTimestamp = cm.CreationTimestamp
		// append
		configMapBriefList = append(configMapBriefList, configMapBrief)
	}

	return configMapBriefList, total, nil
}

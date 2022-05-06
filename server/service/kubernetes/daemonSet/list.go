package daemonSet

import (
	"context"
	"fmt"
	appsV1 "k8s.io/api/apps/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/pddzl/kubeops/server/global"
	"github.com/pddzl/kubeops/server/model/common/request"
	resourceDaemonSet "github.com/pddzl/kubeops/server/model/kubernetes/resource/daemonSet"
)

func (d *DaemonSetService) GetDaemonSetList(namespace string, info request.PageInfo) ([]resourceDaemonSet.DaemonSetBrief, int, error) {
	// 获取daemonSet list数据
	daemonSets, err := global.KOP_KUBERNETES.AppsV1().DaemonSets(namespace).List(context.TODO(), metaV1.ListOptions{})
	if err != nil {
		return nil, 0, err
	}

	var daemonSetList appsV1.DaemonSetList
	// 分页
	end := info.PageSize * info.Page
	offset := info.PageSize * (info.Page - 1)
	total := len(daemonSets.Items)
	if total <= offset {
		return nil, total, nil
	}
	if total < end {
		daemonSetList.Items = daemonSets.Items[offset:]
	} else {
		daemonSetList.Items = daemonSets.Items[offset:end]
	}

	var daemonSetBriefList []resourceDaemonSet.DaemonSetBrief
	// 处理daemonSets数据
	for _, ds := range daemonSets.Items {
		var daemonSetBrief resourceDaemonSet.DaemonSetBrief
		daemonSetBrief.Name = ds.Name
		daemonSetBrief.NameSpace = ds.Namespace
		daemonSetBrief.Pods = fmt.Sprintf("%d / %d", ds.Status.NumberAvailable, ds.Status.DesiredNumberScheduled)
		daemonSetBrief.CreationTimestamp = ds.CreationTimestamp
		// append
		daemonSetBriefList = append(daemonSetBriefList, daemonSetBrief)
	}

	return daemonSetBriefList, 0, nil
}

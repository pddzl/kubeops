package daemonSet

import (
	"context"
	"fmt"
	v1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/pddzl/kubeops/server/global"
	"github.com/pddzl/kubeops/server/model/common/request"
	"github.com/pddzl/kubeops/server/model/kubernetes/resource/daemonSet"
)

func (d *DaemonSetService) GetDaemonSetList(namespace string, info request.PageInfo) ([]daemonSet.DaemonSetBrief, int, error) {
	// 获取原始DaemonSetList数据
	daemonSets, err := global.KOP_KUBERNETES.AppsV1().DaemonSets(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, 0, err
	}

	// 分页
	end := info.PageSize * info.Page
	offset := info.PageSize * (info.Page - 1)
	var daemonSetBriefList []daemonSet.DaemonSetBrief
	var daemonSetList v1.DaemonSetList
	total := len(daemonSets.Items)
	if total <= offset {
		return nil, total, nil
	}
	if total < end {
		daemonSetList.Items = daemonSets.Items[offset:]
	} else {
		daemonSetList.Items = daemonSets.Items[offset:end]
	}

	// 处理daemonSets数据
	for _, ds := range daemonSets.Items {
		var daemonSetBrief daemonSet.DaemonSetBrief
		daemonSetBrief.Name = ds.Name
		daemonSetBrief.NameSpace = ds.Namespace
		daemonSetBrief.Pods = fmt.Sprintf("%d / %d", ds.Status.NumberAvailable, ds.Status.DesiredNumberScheduled)
		daemonSetBrief.CreationTimestamp = ds.CreationTimestamp
		// append
		daemonSetBriefList = append(daemonSetBriefList, daemonSetBrief)
	}

	return daemonSetBriefList, 0, nil
}

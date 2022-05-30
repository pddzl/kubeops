package pod

import (
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/pddzl/kubeops/server/global"
	"github.com/pddzl/kubeops/server/model/kubernetes/resource/common"
)

func (p *PodService) GetPodStatus() ([]common.ChartData, error) {
	// 获取pod list
	list, err := global.KOP_KUBERNETES.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	// 初始化
	chartDataList := make([]common.ChartData, 2)
	chartDataList[0].Value = 0
	chartDataList[0].Name = "正常"
	chartDataList[1].Value = 0
	chartDataList[1].Name = "异常"

	for _, pod := range list.Items {
		if string(pod.Status.Phase) == "Running" || string(pod.Status.Phase) == "Succeeded" {
			chartDataList[0].Value += 1
		} else {
			chartDataList[1].Value += 1
		}
	}

	return chartDataList, err
}

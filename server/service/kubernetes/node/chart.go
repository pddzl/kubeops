package node

import (
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"regexp"

	"github.com/pddzl/kubeops/server/global"
	"github.com/pddzl/kubeops/server/model/kubernetes/resource/common"
	"github.com/pddzl/kubeops/server/utils"
)

func (n *NodeService) GetNodeType() ([]common.ChartData, error) {
	list, err := global.KOP_KUBERNETES.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	// 初始化
	chartDataList := make([]common.ChartData, 2)
	chartDataList[0].Value = 0
	chartDataList[0].Name = "Master"
	chartDataList[1].Value = 0
	chartDataList[1].Name = "Slave"
	roleRe, _ := regexp.Compile("node-role.kubernetes.io/(.*)")
	for _, node := range list.Items {
		for label := range node.ObjectMeta.Labels {
			roles := roleRe.FindStringSubmatch(label)
			if len(roles) == 0 {
				continue
			}
			if ok := utils.Contains(roles, "master"); ok {
				chartDataList[0].Value += 1
				continue
			}
			if ok := utils.Contains(roles, "worker"); ok {
				chartDataList[1].Value += 1
			}
		}
	}

	return chartDataList, err
}

func (n *NodeService) GetNodeStatus() ([]common.ChartData, error) {
	list, err := global.KOP_KUBERNETES.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	// 初始化
	chartDataList := make([]common.ChartData, 2)
	chartDataList[0].Value = 0
	chartDataList[0].Name = "正常"
	chartDataList[1].Value = 0
	chartDataList[1].Name = "异常"

	for _, node := range list.Items {
		for _, condition := range node.Status.Conditions {
			if condition.Type == "Ready" {
				if condition.Status == "True" {
					chartDataList[0].Value += 1
				} else {
					chartDataList[1].Value += 1
				}
			}
		}
	}

	return chartDataList, err
}

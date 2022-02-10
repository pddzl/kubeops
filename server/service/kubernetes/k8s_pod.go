package kubernetes

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"io"
	v1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"

	"github.com/pddzl/kubeops/server/global"
	"github.com/pddzl/kubeops/server/model/common/request"
)

type PodService struct{}

func getNodePods(nodeName string) (*v1.PodList, error) {
	fieldSelector, err := fields.ParseSelector("spec.nodeName=" + nodeName +
		",status.phase!=" + string(v1.PodSucceeded) +
		",status.phase!=" + string(v1.PodFailed))

	if err != nil {
		return nil, err
	}

	return global.KOP_KUBERNETES.CoreV1().Pods(v1.NamespaceAll).List(context.TODO(), metaV1.ListOptions{
		FieldSelector: fieldSelector.String(),
	})
}

func (p *PodService) GetPodLog(namespace string, pod string, lines int64) (log string, err error) {
	fmt.Println(lines)
	opts := v1.PodLogOptions{
		//TailLines: &lines,
	}
	req := global.KOP_KUBERNETES.CoreV1().Pods(namespace).GetLogs(pod, &opts)
	podLogs, err := req.Stream(context.TODO())
	if err != nil {
		return "", err
	}
	defer podLogs.Close()

	buf := new(bytes.Buffer)
	_, err = io.Copy(buf, podLogs)
	if err != nil {
		return "", errors.New("error in copy pod log")
	}
	global.KOP_LOG.Info("here1", zap.Any("pod", buf))
	log = buf.String()
	global.KOP_LOG.Info("here2", zap.Any("pod", log))
	return log, nil
}

func (p *PodService) GetPodList(namespace string, info request.PageInfo) (list []v1.Pod, total int, err error) {
	end := info.PageSize * info.Page
	offset := info.PageSize * (info.Page - 1)
	podList, err := global.KOP_KUBERNETES.CoreV1().Pods(namespace).List(context.TODO(), metaV1.ListOptions{})
	if err != nil {
		return nil, 0, err
	}
	total = len(podList.Items)
	if total <= offset {
		return nil, total, nil
	}
	if total < end {
		list = podList.Items[offset:]
	} else {
		list = podList.Items[offset:end]
	}
	return list, total, nil
}

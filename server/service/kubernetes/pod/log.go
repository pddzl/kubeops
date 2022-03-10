package pod

import (
	"bytes"
	"context"
	"errors"
	"io"
	v1 "k8s.io/api/core/v1"

	"github.com/pddzl/kubeops/server/global"
)

func (p *PodService) GetPodLog(namespace string, container string, pod string, lines int64) (log string, err error) {
	if lines <= 0 {
		lines = 50
	}
	opts := v1.PodLogOptions{
		TailLines: &lines,
	}
	if container != "" {
		opts.Container = container
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
	log = buf.String()
	return log, nil
}

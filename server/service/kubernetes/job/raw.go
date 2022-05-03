package job

import (
	"context"
	"encoding/json"

	"github.com/pddzl/kubeops/server/global"
)

func (j *JobService) GetJobRaw(namespace string, name string) (interface{}, error) {
	// kubectl get --raw '/apis/batch/v1/namespaces/ingress-nginx/jobs/ingress-nginx-admission-create'
	req, err := global.KOP_KUBERNETES.RESTClient().Get().AbsPath("/apis/batch/v1").Resource("jobs").Name(name).Namespace(namespace).DoRaw(context.TODO())
	if err != nil {
		return nil, err
	}
	return json.RawMessage(req), nil
}

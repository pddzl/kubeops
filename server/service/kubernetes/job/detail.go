package job

import (
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/pddzl/kubeops/server/global"
	"github.com/pddzl/kubeops/server/model/kubernetes/api"
	resourceJob "github.com/pddzl/kubeops/server/model/kubernetes/resource/job"
)

func (j *JobService) GetJobDetail(namespace string, name string) (*resourceJob.JobDetail, error) {
	// 获取job原生数据
	job, err := global.KOP_KUBERNETES.BatchV1().Jobs(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}

	// 处理job数据
	var jobDetail resourceJob.JobDetail
	// metadata
	jobDetail.Metadata = api.NewObjectMeta(job.ObjectMeta)
	// spec
	jobDetail.Spec.Parallelism = job.Spec.Parallelism
	jobDetail.Spec.Completions = job.Spec.Completions
	jobDetail.Spec.BackoffLimit = job.Spec.BackoffLimit
	jobDetail.Spec.Selector = job.Spec.Selector.MatchLabels
	// status
	jobDetail.Status.Conditions = job.Status.Conditions

	return &jobDetail, nil
}

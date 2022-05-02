package job

import (
	"context"
	"fmt"
	resourceJob "github.com/pddzl/kubeops/server/model/kubernetes/resource/job"
	batchv1 "k8s.io/api/batch/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/pddzl/kubeops/server/global"
	"github.com/pddzl/kubeops/server/model/common/request"
)

func (j *JobService) GetJobList(namespace string, info request.PageInfo) ([]resourceJob.JobBrief, int, error) {
	// 获取job原生数据
	jobs, err := global.KOP_KUBERNETES.BatchV1().Jobs(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, 0, err
	}

	end := info.PageSize * info.Page
	offset := info.PageSize * (info.Page - 1)
	var jobList batchv1.JobList

	// 分页
	total := len(jobs.Items)
	if total <= offset {
		return nil, total, nil
	}
	if total < end {
		jobList.Items = jobs.Items[offset:]
	} else {
		jobList.Items = jobs.Items[offset:end]
	}

	// 处理jobs数据
	var jobBriefList []resourceJob.JobBrief
	for _, job := range jobList.Items {
		var jobBrief resourceJob.JobBrief
		jobBrief.Name = job.Name
		jobBrief.Namespace = job.Namespace
		var completed uint
		jobBrief.CreationTimestamp = job.CreationTimestamp
		for _, condition := range job.Status.Conditions {
			if condition.Type == "Complete" {
				completed += 1
			}
		}
		jobBrief.Completed = fmt.Sprintf("%d / %d", completed, len(job.Spec.Template.Spec.Containers))
		jobBriefList = append(jobBriefList, jobBrief)
	}

	return jobBriefList, total, nil
}

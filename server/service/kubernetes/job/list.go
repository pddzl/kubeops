package job

import (
	"context"
	"fmt"
	resourceJob "github.com/pddzl/kubeops/server/model/kubernetes/resource/job"
	batchV1 "k8s.io/api/batch/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/pddzl/kubeops/server/global"
	"github.com/pddzl/kubeops/server/model/common/request"
)

func (j *JobService) GetJobList(namespace string, info request.PageInfo) ([]resourceJob.JobBrief, int, error) {
	// 获取job list
	list, err := global.KOP_KUBERNETES.BatchV1().Jobs(namespace).List(context.TODO(), metaV1.ListOptions{})
	if err != nil {
		return nil, 0, err
	}

	var jobList batchV1.JobList
	// 分页
	end := info.PageSize * info.Page
	offset := info.PageSize * (info.Page - 1)
	total := len(list.Items)
	if total <= offset {
		return nil, total, nil
	}
	if total < end {
		jobList.Items = list.Items[offset:]
	} else {
		jobList.Items = list.Items[offset:end]
	}

	var jobBriefList []resourceJob.JobBrief
	// 处理jobs数据
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

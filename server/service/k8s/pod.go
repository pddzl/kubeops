package k8s

import (
	"context"
	"github.com/pddzl/kubeops/server/global"
	modelK8s "github.com/pddzl/kubeops/server/model/k8s"
	coreV1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type PodService struct{}

// GetPods 获取集群pods
func (ps *PodService) GetPods(namespace string, page int, pageSize int) ([]modelK8s.PodBrief, int, error) {
	// 获取pod list
	list, err := global.KOP_K8S_Client.CoreV1().Pods(namespace).List(context.TODO(), metaV1.ListOptions{})
	if err != nil {
		return nil, 0, err
	}

	var podList coreV1.PodList
	// 分页
	end := pageSize * page
	offset := pageSize * (page - 1)
	total := len(list.Items)
	if total <= offset {
		return nil, total, nil
	}
	if total < end {
		podList.Items = list.Items[offset:]
	} else {
		podList.Items = list.Items[offset:end]
	}

	// 处理list数据
	var podBriefList []modelK8s.PodBrief
	for _, pod := range podList.Items {
		var podBrief modelK8s.PodBrief
		podBrief.Name = pod.Name
		podBrief.Namespace = pod.Namespace
		podBrief.Node = pod.Spec.NodeName
		podBrief.Status = string(pod.Status.Phase)
		podBrief.CreationTimestamp = pod.CreationTimestamp.Time
		// append
		podBriefList = append(podBriefList, podBrief)
	}
	return podBriefList, total, nil
}

// GetPodDetail 获取指定pod详情
func (ps *PodService) GetPodDetail(namespace string, name string) (*modelK8s.PodDetail, error) {
	podDetailRaw, err := global.KOP_K8S_Client.CoreV1().Pods(namespace).Get(context.TODO(), name, metaV1.GetOptions{})
	if err != nil {
		return nil, nil
	}

	var podDetail modelK8s.PodDetail

	// metadata
	podDetail.MetaData = modelK8s.NewObjectMeta(podDetailRaw.ObjectMeta)

	// spec
	podDetail.Spec.NodeName = podDetailRaw.Spec.NodeName
	podDetail.Spec.RestartPolicy = string(podDetailRaw.Spec.RestartPolicy)
	podDetail.Spec.ServiceAccountName = podDetailRaw.Spec.ServiceAccountName

	// spec -> initContainer | container
	podDetail.Spec.InitContainers = extractContainerInfo(podDetailRaw.Spec.InitContainers, podDetailRaw)
	podDetail.Spec.Containers = extractContainerInfo(podDetailRaw.Spec.Containers, podDetailRaw)

	// status
	podDetail.Status.Phase = string(podDetailRaw.Status.Phase)
	podDetail.Status.PodIP = podDetailRaw.Status.PodIP
	podDetail.Status.QOSClass = string(podDetailRaw.Status.QOSClass)
	podDetail.Status.Conditions = podDetailRaw.Status.Conditions

	return &podDetail, err
}

func extractContainerInfo(containerList []coreV1.Container, pod *coreV1.Pod) (containers []modelK8s.Container) {
	for _, value := range containerList {
		var container modelK8s.Container
		container.Name = value.Name
		container.Image = value.Image
		container.Command = value.Command
		container.Args = value.Args
		container.SecurityContext = value.SecurityContext
		container.LivenessProbe = value.LivenessProbe
		container.ReadinessProbe = value.ReadinessProbe
		container.StartupProbe = value.StartupProbe
		// volume
		for _, volume := range value.VolumeMounts {
			var volumeMount modelK8s.VolumeMount
			volumeMount.Name = volume.Name
			volumeMount.ReadOnly = volume.ReadOnly
			volumeMount.MountPath = volume.MountPath
			for _, specV := range pod.Spec.Volumes {
				if volume.Name == specV.Name {
					if specV.VolumeSource.ConfigMap != nil {
						volumeMount.VolumeType = "configMap"
					} else if specV.VolumeSource.Secret != nil {
						volumeMount.VolumeType = "secret"
					} else if specV.VolumeSource.EmptyDir != nil {
						volumeMount.VolumeType = "emptyDir"
					} else if specV.VolumeSource.Projected != nil {
						volumeMount.VolumeType = "projected"
					} else if specV.VolumeSource.HostPath != nil {
						volumeMount.VolumeType = "hostPath"
					} // ...待续
				}
			}
			container.VolumeMounts = append(container.VolumeMounts, volumeMount)
		}
		// status
		for _, status := range pod.Status.ContainerStatuses {
			if status.Name == value.Name {
				container.Status = &status
			}
		}
		containers = append(containers, container)
	}
	return
}

package pod

import (
	"context"
	"encoding/base64"
	"fmt"
	v1 "k8s.io/api/core/v1"
	res "k8s.io/apimachinery/pkg/api/resource"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"math"
	"strconv"

	"github.com/pddzl/kubeops/server/global"
	"github.com/pddzl/kubeops/server/model/kubernetes"
)

// 获取pod详情

func (p *PodService) GetPodDetail(namespace string, name string) (info interface{}, err error) {
	var podDetail kubernetes.PodDetail

	channels := &kubernetes.ResourceChannels{
		ConfigMapList: kubernetes.GetConfigMapListChannel(kubernetes.NewSameNamespaceQuery(namespace), 1),
		SecretList:    kubernetes.GetSecretListChannel(kubernetes.NewSameNamespaceQuery(namespace), 1),
	}

	pod, err := global.KOP_KUBERNETES.CoreV1().Pods(namespace).Get(context.TODO(), name, metaV1.GetOptions{})
	if err != nil {
		return "", nil
	}

	// configMap
	configMapList := <-channels.ConfigMapList.List
	err = <-channels.ConfigMapList.Error
	if err != nil {
		return "", err
	}

	// secret
	secretList := <-channels.SecretList.List
	err = <-channels.SecretList.Error
	if err != nil {
		return "", err
	}

	// metaData
	podDetail.MetaData.Name = pod.Name
	podDetail.MetaData.Namespace = pod.Namespace
	podDetail.MetaData.Uid = string(pod.UID)
	podDetail.MetaData.CreateTimestamp = pod.CreationTimestamp
	podDetail.MetaData.Labels = pod.Labels
	podDetail.MetaData.OwnerReferences.Controller = *pod.OwnerReferences[0].Controller
	podDetail.MetaData.OwnerReferences.Name = pod.OwnerReferences[0].Name
	podDetail.MetaData.OwnerReferences.Kind = pod.OwnerReferences[0].Kind
	// ownerReferences
	podDetail.ResourceInfo.Phase = string(pod.Status.Phase)
	podDetail.ResourceInfo.Node = pod.Spec.NodeName
	podDetail.ResourceInfo.IP = pod.Status.HostIP
	podDetail.ResourceInfo.QosClass = string(pod.Status.QOSClass)
	podDetail.ResourceInfo.RestartPolicy = string(pod.Spec.RestartPolicy)
	podDetail.ResourceInfo.Restarts = pod.Status.ContainerStatuses[0].RestartCount
	podDetail.ResourceInfo.ServiceAccount = pod.Spec.ServiceAccountName
	// Conditions
	for _, condition := range pod.Status.Conditions {
		var podCondition kubernetes.Conditions
		podCondition.Type = string(condition.Type)
		podCondition.Status = string(condition.Status)
		podCondition.LastProbeTime = condition.LastProbeTime
		podCondition.LastTransitionTime = condition.LastTransitionTime
		// append
		podDetail.Conditions = append(podDetail.Conditions, podCondition)
	}
	// Containers
	podDetail.Containers = extractContainerInfo(pod.Spec.Containers, pod, configMapList, secretList)
	// initContainers
	podDetail.InitContainers = extractContainerInfo(pod.Spec.InitContainers, pod, configMapList, secretList)
	return podDetail, nil
}

func extractContainerInfo(containerList []v1.Container, pod *v1.Pod, configMaps *v1.ConfigMapList, secrets *v1.SecretList) []kubernetes.Container {
	containers := make([]kubernetes.Container, 0)
	for _, container := range containerList {
		vars := make([]kubernetes.EnvVar, 0)
		for _, envVar := range container.Env {
			variable := kubernetes.EnvVar{
				Name:      envVar.Name,
				Value:     envVar.Value,
				ValueFrom: envVar.ValueFrom,
			}
			if variable.ValueFrom != nil {
				variable.Value = evalValueFrom(variable.ValueFrom, &container, pod,
					configMaps, secrets)
			}
			vars = append(vars, variable)
		}
		vars = append(vars, evalEnvFrom(container, configMaps, secrets)...)

		volumeMounts := extractContainerMounts(container, pod)

		containers = append(containers, kubernetes.Container{
			Name:            container.Name,
			Image:           container.Image,
			Env:             vars,
			Commands:        container.Command,
			Args:            container.Args,
			VolumeMounts:    volumeMounts,
			SecurityContext: container.SecurityContext,
			Status:          extractContainerStatus(pod, &container),
			LivenessProbe:   container.LivenessProbe,
			ReadinessProbe:  container.ReadinessProbe,
			StartupProbe:    container.StartupProbe,
		})
	}
	return containers
}

func extractContainerMounts(container v1.Container, pod *v1.Pod) []kubernetes.VolumeMount {
	volumeMounts := make([]kubernetes.VolumeMount, 0)
	for _, aVolumeMount := range container.VolumeMounts {
		volumeMount := kubernetes.VolumeMount{
			Name:      aVolumeMount.Name,
			ReadOnly:  aVolumeMount.ReadOnly,
			MountPath: aVolumeMount.MountPath,
			SubPath:   aVolumeMount.SubPath,
			Volume:    getVolume(pod.Spec.Volumes, aVolumeMount.Name),
		}
		volumeMounts = append(volumeMounts, volumeMount)
	}
	return volumeMounts
}

func getVolume(volumes []v1.Volume, volumeName string) v1.Volume {
	for _, volume := range volumes {
		if volume.Name == volumeName {
			// yes, this is exponential, but N is VERY small, so the malloc for creating a named dictionary would probably take longer
			return volume
		}
	}
	return v1.Volume{}
}

// evalValueFrom evaluates environment value from given source. For more details check:
// https://github.com/kubernetes/kubernetes/blob/d82e51edc5f02bff39661203c9b503d054c3493b/pkg/kubectl/describe.go#L1056
func evalValueFrom(src *v1.EnvVarSource, container *v1.Container, pod *v1.Pod, configMaps *v1.ConfigMapList, secrets *v1.SecretList) string {
	switch {
	case src.ConfigMapKeyRef != nil:
		name := src.ConfigMapKeyRef.LocalObjectReference.Name
		for _, configMap := range configMaps.Items {
			if configMap.ObjectMeta.Name == name {
				return configMap.Data[src.ConfigMapKeyRef.Key]
			}
		}
	case src.SecretKeyRef != nil:
		name := src.SecretKeyRef.LocalObjectReference.Name
		for _, secret := range secrets.Items {
			if secret.ObjectMeta.Name == name {
				return base64.StdEncoding.EncodeToString([]byte(
					secret.Data[src.SecretKeyRef.Key]))
			}
		}
	case src.ResourceFieldRef != nil:
		valueFrom, err := extractContainerResourceValue(src.ResourceFieldRef, container)
		if err != nil {
			valueFrom = ""
		}
		resource := src.ResourceFieldRef.Resource
		if valueFrom == "0" && (resource == "limits.cpu" || resource == "limits.memory") {
			valueFrom = "node allocatable"
		}
		return valueFrom
	case src.FieldRef != nil:
		gv, err := schema.ParseGroupVersion(src.FieldRef.APIVersion)
		if err != nil {
			global.KOP_LOG.Error(err.Error())
			return ""
		}
		gvk := gv.WithKind("Pod")
		internalFieldPath, _, err := runtime.NewScheme().ConvertFieldLabel(gvk, src.FieldRef.FieldPath, "")
		if err != nil {
			global.KOP_LOG.Error(err.Error())
			return ""
		}
		valueFrom, err := ExtractFieldPathAsString(pod, internalFieldPath)
		if err != nil {
			global.KOP_LOG.Error(err.Error())
			return ""
		}
		return valueFrom
	}
	return ""
}

// extractContainerResourceValue extracts the value of a resource in an already known container.
func extractContainerResourceValue(fs *v1.ResourceFieldSelector, container *v1.Container) (string, error) {
	divisor := res.Quantity{}
	if divisor.Cmp(fs.Divisor) == 0 {
		divisor = res.MustParse("1")
	} else {
		divisor = fs.Divisor
	}

	switch fs.Resource {
	case "limits.cpu":
		return strconv.FormatInt(int64(math.Ceil(float64(container.Resources.Limits.
			Cpu().MilliValue())/float64(divisor.MilliValue()))), 10), nil
	case "limits.memory":
		return strconv.FormatInt(int64(math.Ceil(float64(container.Resources.Limits.
			Memory().Value())/float64(divisor.Value()))), 10), nil
	case "requests.cpu":
		return strconv.FormatInt(int64(math.Ceil(float64(container.Resources.Requests.
			Cpu().MilliValue())/float64(divisor.MilliValue()))), 10), nil
	case "requests.memory":
		return strconv.FormatInt(int64(math.Ceil(float64(container.Resources.Requests.
			Memory().Value())/float64(divisor.Value()))), 10), nil
	}

	return "", fmt.Errorf("unsupported container resource : %v", fs.Resource)
}

func extractContainerStatus(pod *v1.Pod, container *v1.Container) *v1.ContainerStatus {
	for _, status := range pod.Status.ContainerStatuses {
		if status.Name == container.Name {
			return &status
		}
	}

	return nil
}

func evalEnvFrom(container v1.Container, configMaps *v1.ConfigMapList, secrets *v1.SecretList) []kubernetes.EnvVar {
	vars := make([]kubernetes.EnvVar, 0)
	for _, envFromVar := range container.EnvFrom {
		switch {
		case envFromVar.ConfigMapRef != nil:
			name := envFromVar.ConfigMapRef.LocalObjectReference.Name
			for _, configMap := range configMaps.Items {
				if configMap.ObjectMeta.Name == name {
					for key, value := range configMap.Data {
						valueFrom := &v1.EnvVarSource{
							ConfigMapKeyRef: &v1.ConfigMapKeySelector{
								LocalObjectReference: v1.LocalObjectReference{
									Name: name,
								},
								Key: key,
							},
						}
						variable := kubernetes.EnvVar{
							Name:      envFromVar.Prefix + key,
							Value:     value,
							ValueFrom: valueFrom,
						}
						vars = append(vars, variable)
					}
					break
				}
			}
		case envFromVar.SecretRef != nil:
			name := envFromVar.SecretRef.LocalObjectReference.Name
			for _, secret := range secrets.Items {
				if secret.ObjectMeta.Name == name {
					for key, value := range secret.Data {
						valueFrom := &v1.EnvVarSource{
							SecretKeyRef: &v1.SecretKeySelector{
								LocalObjectReference: v1.LocalObjectReference{
									Name: name,
								},
								Key: key,
							},
						}
						variable := kubernetes.EnvVar{
							Name:      envFromVar.Prefix + key,
							Value:     base64.StdEncoding.EncodeToString(value),
							ValueFrom: valueFrom,
						}
						vars = append(vars, variable)
					}
					break
				}
			}
		}
	}
	return vars
}

package pod

import (
	"context"
	"encoding/base64"
	"fmt"
	corev1 "k8s.io/api/core/v1"
	res "k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"math"
	"strconv"

	"github.com/pddzl/kubeops/server/global"
	"github.com/pddzl/kubeops/server/model/kubernetes/api"
	resourcePod "github.com/pddzl/kubeops/server/model/kubernetes/resource/pod"
)

// 获取pod详情

func (p *PodService) GetPodDetail(namespace string, name string) (*resourcePod.PodDetail, error) {
	// 获取原生pod数据
	podRaw, err := global.KOP_KUBERNETES.CoreV1().Pods(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		return nil, nil
	}

	// 获取 configMap
	configMapList, err := global.KOP_KUBERNETES.CoreV1().ConfigMaps(namespace).List(context.TODO(), api.ListEverything)
	if err != nil {
		return nil, err
	}

	// 获取secret
	secretList, err := global.KOP_KUBERNETES.CoreV1().Secrets(namespace).List(context.TODO(), api.ListEverything)
	if err != nil {
		return nil, err
	}

	// 处理原生pod数据
	var podDetail resourcePod.PodDetail
	// metadata
	podDetail.Metadata.Name = podRaw.Name
	podDetail.Metadata.Namespace = podRaw.Namespace
	podDetail.Metadata.Labels = podRaw.Labels
	podDetail.Metadata.CreationTimestamp = podRaw.CreationTimestamp
	podDetail.Metadata.UID = podRaw.UID
	podDetail.Metadata.OwnerReferences = podRaw.OwnerReferences
	// spec
	podDetail.Spec.RestartPolicy = string(podRaw.Spec.RestartPolicy)
	podDetail.Spec.NodeName = podRaw.Spec.NodeName
	podDetail.Spec.ServiceAccountName = podRaw.Spec.ServiceAccountName
	// spec -> Containers
	podDetail.Spec.Containers = extractContainerInfo(podRaw.Spec.Containers, podRaw, configMapList, secretList)
	// initContainers
	if len(podRaw.Spec.InitContainers) > 0 {
		podDetail.Spec.InitContainers = extractContainerInfo(podRaw.Spec.InitContainers, podRaw, configMapList, secretList)
	}
	// status
	podDetail.Status.Phase = string(podRaw.Status.Phase)
	podDetail.Status.PodIP = podRaw.Status.PodIP
	podDetail.Status.QOSClass = string(podRaw.Status.QOSClass)
	podDetail.Status.Conditions = podRaw.Status.Conditions

	return &podDetail, nil
}

func extractContainerInfo(containerList []corev1.Container, pod *corev1.Pod, configMaps *corev1.ConfigMapList, secrets *corev1.SecretList) []resourcePod.Container {
	containers := make([]resourcePod.Container, 0)
	for _, container := range containerList {
		vars := make([]resourcePod.EnvVar, 0)
		for _, envVar := range container.Env {
			variable := resourcePod.EnvVar{
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

		containers = append(containers, resourcePod.Container{
			Name:            container.Name,
			Image:           container.Image,
			Env:             vars,
			Command:         container.Command,
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

func extractContainerMounts(container corev1.Container, podDeatil *corev1.Pod) []resourcePod.VolumeMount {
	volumeMounts := make([]resourcePod.VolumeMount, 0)
	for _, aVolumeMount := range container.VolumeMounts {
		volumeMount := resourcePod.VolumeMount{
			Name:      aVolumeMount.Name,
			ReadOnly:  aVolumeMount.ReadOnly,
			MountPath: aVolumeMount.MountPath,
			SubPath:   aVolumeMount.SubPath,
			Volume:    getVolume(podDeatil.Spec.Volumes, aVolumeMount.Name),
		}
		volumeMounts = append(volumeMounts, volumeMount)
	}
	return volumeMounts
}

func getVolume(volumes []corev1.Volume, volumeName string) corev1.Volume {
	for _, volume := range volumes {
		if volume.Name == volumeName {
			// yes, this is exponential, but N is VERY small, so the malloc for creating a named dictionary would probably take longer
			return volume
		}
	}
	return corev1.Volume{}
}

// evalValueFrom evaluates environment value from given source. For more details check:
// https://github.com/kubernetes/kubernetes/blob/d82e51edc5f02bff39661203c9b503d054c3493b/pkg/kubectl/describe.go#L1056
func evalValueFrom(src *corev1.EnvVarSource, container *corev1.Container, pod *corev1.Pod, configMaps *corev1.ConfigMapList, secrets *corev1.SecretList) string {
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
func extractContainerResourceValue(fs *corev1.ResourceFieldSelector, container *corev1.Container) (string, error) {
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

func extractContainerStatus(pod *corev1.Pod, container *corev1.Container) *corev1.ContainerStatus {
	for _, status := range pod.Status.ContainerStatuses {
		if status.Name == container.Name {
			return &status
		}
	}

	return nil
}

func evalEnvFrom(container corev1.Container, configMaps *corev1.ConfigMapList, secrets *corev1.SecretList) []resourcePod.EnvVar {
	vars := make([]resourcePod.EnvVar, 0)
	for _, envFromVar := range container.EnvFrom {
		switch {
		case envFromVar.ConfigMapRef != nil:
			name := envFromVar.ConfigMapRef.LocalObjectReference.Name
			for _, configMap := range configMaps.Items {
				if configMap.ObjectMeta.Name == name {
					for key, value := range configMap.Data {
						valueFrom := &corev1.EnvVarSource{
							ConfigMapKeyRef: &corev1.ConfigMapKeySelector{
								LocalObjectReference: corev1.LocalObjectReference{
									Name: name,
								},
								Key: key,
							},
						}
						variable := resourcePod.EnvVar{
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
						valueFrom := &corev1.EnvVarSource{
							SecretKeyRef: &corev1.SecretKeySelector{
								LocalObjectReference: corev1.LocalObjectReference{
									Name: name,
								},
								Key: key,
							},
						}
						variable := resourcePod.EnvVar{
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

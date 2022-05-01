package pod

import (
	"github.com/pddzl/kubeops/server/model/kubernetes/api"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type PodDetail struct {
	Metadata metadata  `json:"metadata,omitempty"`
	Spec     podSpec   `json:"spec,omitempty"`
	Status   podStatus `json:"status,omitempty"`
}

type metadata struct {
	api.ObjectMeta
	OwnerReferences []metav1.OwnerReference `json:"ownerReferences,omitempty"`
}

type podSpec struct {
	InitContainers     []Container `json:"initContainers,omitempty"`
	Containers         []Container `json:"containers"`
	RestartPolicy      string      `json:"restartPolicy,omitempty"`
	ServiceAccountName string      `json:"serviceAccountName,omitempty"`
	NodeName           string      `json:"nodeName,omitempty"`
}

type podStatus struct {
	Phase      string                `json:"phase,omitempty"`
	Conditions []corev1.PodCondition `json:"conditions,omitempty"`
	PodIP      string                `json:"podIP,omitempty"`
	QOSClass   string                `json:"qosClass,omitempty"`
}

type Container struct {
	Name            string                  `json:"name"`
	Image           string                  `json:"image,omitempty"`
	Command         []string                `json:"command,omitempty"`
	Args            []string                `json:"args,omitempty"`
	Env             []EnvVar                `json:"env,omitempty"`
	VolumeMounts    []VolumeMount           `json:"volumeMounts,omitempty"`
	LivenessProbe   *corev1.Probe           `json:"livenessProbe,omitempty"`
	ReadinessProbe  *corev1.Probe           `json:"readinessProbe,omitempty"`
	StartupProbe    *corev1.Probe           `json:"startupProbe,omitempty"`
	SecurityContext *corev1.SecurityContext `json:"securityContext,omitempty"`
	Status          *corev1.ContainerStatus `json:"status"`
}

type EnvVar struct {
	Name      string               `json:"name"`
	Value     string               `json:"value"`
	ValueFrom *corev1.EnvVarSource `json:"valueFrom"`
}

type VolumeMount struct {
	Name      string        `json:"name"`
	ReadOnly  bool          `json:"readOnly,omitempty"`
	MountPath string        `json:"mountPath"`
	SubPath   string        `json:"subPath,omitempty"`
	Volume    corev1.Volume `json:"volume"`
}

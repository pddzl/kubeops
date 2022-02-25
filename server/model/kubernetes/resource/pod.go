package resource

import (
	v1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type PodDetail struct {
	MetaData        metadata                `json:"metadata"`
	ResourceInfo    resourceInfo            `json:"resource_info"`
	Conditions      []Conditions            `json:"conditions"`
	OwnerReferences []metaV1.OwnerReference `json:"ownerReferences,omitempty"`
	Containers      []Container             `json:"containers"`
	InitContainers  []Container             `json:"initContainers,omitempty"`
}

type metadata struct {
	Name              string            `json:"name,omitempty"`
	Namespace         string            `json:"namespace,omitempty"`
	Uid               string            `json:"uid,omitempty"`
	CreationTimestamp metaV1.Time       `json:"creationTimestamp,omitempty"`
	Labels            map[string]string `json:"labels,omitempty"`
}

type resourceInfo struct {
	Phase          string `json:"phase,omitempty"`
	Node           string `json:"node,omitempty"`
	IP             string `json:"ip,omitempty"`
	QosClass       string `json:"qosClass,omitempty"`
	RestartPolicy  string `json:"restartPolicy,omitempty"`
	Restarts       int32  `json:"restarts"`
	ServiceAccount string `json:"serviceAccount,omitempty"`
}

type Conditions struct {
	Type               string      `json:"type"`
	Status             string      `json:"status"`
	LastProbeTime      metaV1.Time `json:"lastProbeTime"`
	LastTransitionTime metaV1.Time `json:"lastTransitionTime"`
}

// Container represents a docker/rkt/etc. container that lives in a pod.
type Container struct {
	// Name of the container.
	Name string `json:"name"`

	// Image URI of the container.
	Image string `json:"image"`

	// List of environment variables.
	Env []EnvVar `json:"env"`

	// Commands of the container
	Commands []string `json:"commands"`

	// Command arguments
	Args []string `json:"args"`

	// Information about mounted volumes
	VolumeMounts []VolumeMount `json:"volumeMounts"`

	// Security configuration that will be applied to a container.
	SecurityContext *v1.SecurityContext `json:"securityContext"`

	// Status of a pod container
	Status *v1.ContainerStatus `json:"status"`

	// Probes
	LivenessProbe  *v1.Probe `json:"livenessProbe"`
	ReadinessProbe *v1.Probe `json:"readinessProbe"`
	StartupProbe   *v1.Probe `json:"startupProbe"`
}

// EnvVar represents an environment variable of a container.
type EnvVar struct {
	// Name of the variable.
	Name string `json:"name"`

	// Value of the variable. May be empty if value from is defined.
	Value string `json:"value"`

	// Defined for derived variables. If non-null, the value is get from the reference.
	// Note that this is an API struct. This is intentional, as EnvVarSources are plain struct
	// references.
	ValueFrom *v1.EnvVarSource `json:"valueFrom"`
}

type VolumeMount struct {
	// Name of the variable.
	Name string `json:"name"`

	// Is the volume read only ?
	ReadOnly bool `json:"readOnly"`

	// Path within the container at which the volume should be mounted. Must not contain ':'.
	MountPath string `json:"mountPath"`

	// Path within the volume from which the container's volume should be mounted. Defaults to "" (volume's root).
	SubPath string `json:"subPath"`

	// Information about the Volume itself
	Volume v1.Volume `json:"volume"`
}

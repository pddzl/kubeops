package kubernetes

import (
	v1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type NodeDetail struct {
	ObjectMeta             objectMeta             `json:"objectMeta"`
	NodeAllocatedResources NodeAllocatedResources `json:"allocatedResources"`
	PodCIDR                string                 `json:"podCIDR"`
	Unschedulable          bool                   `json:"unschedulable"`
	NodeInfo               v1.NodeSystemInfo      `json:"nodeInfo"`
	Conditions             []Condition            `json:"conditions"`
	ContainerImages        []string               `json:"containerImages"`
	PodList                []Pod                  `json:"podList"`
	Taints                 []v1.Taint             `json:"taints,omitempty"`
	Addresses              []Addresses            `json:"addresses,omitempty"`
}

type objectMeta struct {
	Name              string            `json:"name,omitempty"`
	Labels            map[string]string `json:"labels,omitempty"`
	Annotations       map[string]string `json:"annotations,omitempty"`
	CreationTimestamp metaV1.Time       `json:"creationTimestamp,omitempty"`
	UID               string            `json:"uid,omitempty"`
}

type NodeAllocatedResources struct {
	// CPURequests is number of allocated milicores.
	CPURequests int64 `json:"cpuRequests"`

	// CPURequestsFraction is a fraction of CPU, that is allocated.
	CPURequestsFraction float64 `json:"cpuRequestsFraction"`

	// CPULimits is defined CPU limit.
	CPULimits int64 `json:"cpuLimits"`

	// CPULimitsFraction is a fraction of defined CPU limit, can be over 100%, i.e.
	// overcommitted.
	CPULimitsFraction float64 `json:"cpuLimitsFraction"`

	// CPUCapacity is specified node CPU capacity in milicores.
	CPUCapacity int64 `json:"cpuCapacity"`

	// MemoryRequests is a fraction of memory, that is allocated.
	MemoryRequests int64 `json:"memoryRequests"`

	// MemoryRequestsFraction is a fraction of memory, that is allocated.
	MemoryRequestsFraction float64 `json:"memoryRequestsFraction"`

	// MemoryLimits is defined memory limit.
	MemoryLimits int64 `json:"memoryLimits"`

	// MemoryLimitsFraction is a fraction of defined memory limit, can be over 100%, i.e.
	// overcommitted.
	MemoryLimitsFraction float64 `json:"memoryLimitsFraction"`

	// MemoryCapacity is specified node memory capacity in bytes.
	MemoryCapacity int64 `json:"memoryCapacity"`

	// AllocatedPods in number of currently allocated pods on the node.
	AllocatedPods int `json:"allocatedPods"`

	// PodCapacity is maximum number of pods, that can be allocated on the node.
	PodCapacity int64 `json:"podCapacity"`

	// PodFraction is a fraction of pods, that can be allocated on given node.
	PodFraction float64 `json:"podFraction"`
}

type Condition struct {
	// Type of a condition.
	Type string `json:"type"`
	// Status of a condition.
	Status string `json:"status"`
	// Last probe time of a condition.
	LastProbeTime metaV1.Time `json:"lastProbeTime"`
	// Last transition time of a condition.
	LastTransitionTime metaV1.Time `json:"lastTransitionTime"`
	// Reason of a condition.
	Reason string `json:"reason"`
	// Message of a condition.
	Message string `json:"message"`
}

type Pod struct {
	Name              string      `json:"name"`
	Namespace         string      `json:"namespace"`
	Image             string      `json:"image"`
	Node              string      `json:"node"`
	Resource          resource    `json:"resource"`
	Status            string      `json:"status"`
	CreationTimestamp metaV1.Time `json:"creationTimestamp"`
}

type resource struct {
	CpuLimit       string `json:"cpuLimit"`
	MemoryLimit    string `json:"memoryLimit"`
	CpuRequests    string `json:"cpuRequests"`
	MemoryRequests string `json:"memoryRequests"`
}

type Addresses struct {
	Type    string `json:"type"`
	Address string `json:"address"`
}

package node

import (
	"github.com/pddzl/kubeops/server/model/kubernetes/api"
	"github.com/pddzl/kubeops/server/model/kubernetes/resource/pod"
	corev1 "k8s.io/api/core/v1"
)

type NodeDetail struct {
	Metadata               api.ObjectMeta         `json:"metadata"`
	Spec                   nodeSpec               `json:"spec,omitempty"`
	Status                 NodeStatus             `json:"status,omitempty"`
	NodeAllocatedResources NodeAllocatedResources `json:"allocatedResources"`
	Pods                   []pod.PodBrief         `json:"pods"`
}

type nodeSpec struct {
	PodCIDR       string         `json:"podCIDR,omitempty"`
	Unschedulable bool           `json:"unschedulable,omitempty"`
	Taints        []corev1.Taint `json:"taints,omitempty"`
}

type NodeStatus struct {
	Phase           string                     `json:"phase,omitempty"`
	Conditions      []corev1.NodeCondition     `json:"conditions,omitempty"`
	Addresses       []corev1.NodeAddress       `json:"addresses,omitempty"`
	DaemonEndpoints corev1.NodeDaemonEndpoints `json:"Port"` // kubelet服务端口
	NodeInfo        corev1.NodeSystemInfo      `json:"nodeInfo,omitempty"`
	Images          []corev1.ContainerImage    `json:"images,omitempty"`
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

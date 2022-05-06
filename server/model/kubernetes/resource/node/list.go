package node

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

type NodeBrief struct {
	Name              string      `json:"name,omitempty"`
	InternalIP        string      `json:"internalIP,omitempty"`
	Roles             []string    `json:"role,omitempty"`
	Status            string      `json:"status,omitempty"`
	Cpu               string      `json:"cpu,omitempty"`
	Memory            string      `json:"memory,omitempty"`
	CreationTimestamp metav1.Time `json:"creationTimestamp,omitempty"`
}

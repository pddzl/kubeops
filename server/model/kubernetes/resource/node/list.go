package node

type NodeBrief struct {
	Name       string   `json:"name,omitempty"`
	InternalIP string   `json:"internalIP,omitempty"`
	Roles      []string `json:"roles,omitempty"`
	Status     string   `json:"status,omitempty"`
	Cpu        string   `json:"cpu,omitempty"`
	Memory     string   `json:"memory,omitempty"`
}

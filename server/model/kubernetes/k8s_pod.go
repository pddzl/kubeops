package kubernetes

import "time"

type PodDetail struct {
	MetaData     metadata     `json:"metadata"`
	ResourceInfo resourceInfo `json:"resource_info"`
	Conditions   []Conditions `json:"conditions"`
}

type metadata struct {
	Name              string            `json:"name"`
	Namespace         string            `json:"namespace"`
	Uid               string            `json:"uid"`
	CreationTimestamp time.Time         `json:"creationTimestamp"`
	Labels            map[string]string `json:"labels"`
	OwnerReferences   ownerReferences   `json:"ownerReferences"`
}

type ownerReferences struct {
	Controller bool   `json:"controller"`
	Name       string `json:"name"`
	Kind       string `json:"kind"`
}

type resourceInfo struct {
	Phase          string `json:"phase"`
	Node           string `json:"node"`
	IP             string `json:"ip"`
	QosClass       string `json:"qosClass"`
	RestartPolicy  string `json:"restartPolicy"`
	Restarts       int32  `json:"restarts"`
	ServiceAccount string `json:"serviceAccount"`
}

type Conditions struct {
	Type               string    `json:"type"`
	Status             string    `json:"status"`
	LastProbeTime      time.Time `json:"lastProbeTime"`
	LastTransitionTime time.Time `json:"lastTransitionTime"`
}

package k8s

import "time"

type RelatedPod struct {
	ObjectMeta ObjectMeta `json:"metadata"`
	Status     string     `json:"status"`
	NodeName   string     `json:"nodeName"`
}

type PodBrief struct {
	Name              string    `json:"name,omitempty"`
	Namespace         string    `json:"namespace,omitempty"`
	Status            string    `json:"status,omitempty"`
	Node              string    `json:"node,omitempty"`
	CreationTimestamp time.Time `json:"creationTimestamp,omitempty"`
}

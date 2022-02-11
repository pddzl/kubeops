package kubernetes

import "time"

type NameSpaceBrief struct {
	Name              string            `json:"name"`
	Labels            map[string]string `json:"labels"`
	CreationTimestamp time.Time         `json:"createTimestamp"`
	Status            string            `json:"status"`
}

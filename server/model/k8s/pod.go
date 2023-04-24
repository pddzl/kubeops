package k8s

type RelatedPod struct {
	ObjectMeta ObjectMeta `json:"metadata"`
	Status     string     `json:"status"`
	NodeName   string     `json:"nodeName"`
}

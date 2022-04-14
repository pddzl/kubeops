package services

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

type ServicesBrief struct {
	Name              string      `json:"name,omitempty"`
	NameSpace         string      `json:"namespace,omitempty"`
	ClusterIP         string      `json:"clusterIP,omitempty"`
	Type              string      `json:"type,omitempty"`
	CreationTimestamp metav1.Time `json:"creationTimestamp,omitempty"`
}

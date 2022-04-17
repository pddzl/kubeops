package services

import (
	"github.com/pddzl/kubeops/server/model/kubernetes/api"
	"k8s.io/apimachinery/pkg/util/intstr"
)

type ServiceDetail struct {
	ObjectMeta api.ObjectMeta `json:"metadata,omitempty"`
	Spec       serviceSpec    `json:"spec,omitempty"`
}

type serviceSpec struct {
	Ports           []ServicePort     `json:"ports,omitempty"`
	Selector        map[string]string `json:"selector,omitempty"`
	ClusterIP       string            `json:"clusterIP,omitempty"`
	ClusterIPs      []string          `json:"clusterIPs,omitempty"`
	Type            string            `json:"type,omitempty"`
	ExternalIPs     []string          `json:"externalIPs,omitempty"`
	SessionAffinity string            `json:"sessionAffinity,omitempty"`
}

type ServicePort struct {
	Name        string             `json:"name,omitempty"`
	Protocol    string             `json:"protocol,omitempty"`
	AppProtocol *string            `json:"appProtocol,omitempty"`
	Port        int32              `json:"port" protobuf:"varint,3,opt,name=port"`
	TargetPort  intstr.IntOrString `json:"targetPort,omitempty"`
	NodePort    int32              `json:"nodePort,omitempty"`
}

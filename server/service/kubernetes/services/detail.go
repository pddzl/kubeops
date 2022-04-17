package services

import (
	"context"
	"github.com/pddzl/kubeops/server/global"
	"github.com/pddzl/kubeops/server/model/kubernetes/api"
	resourceServices "github.com/pddzl/kubeops/server/model/kubernetes/resource/services"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (s *ServicesService) GetServiceDetail(namespace string, name string) (*resourceServices.ServiceDetail, error) {
	// 获取service原生数据
	services, err := global.KOP_KUBERNETES.CoreV1().Services(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}

	// 处理service数据
	var serviceDetail resourceServices.ServiceDetail
	// metadata
	serviceDetail.ObjectMeta = api.NewObjectMeta(services.ObjectMeta)
	// spec
	for _, port := range services.Spec.Ports {
		var servicePort resourceServices.ServicePort
		servicePort.Name = port.Name
		servicePort.Protocol = string(port.Protocol)
		servicePort.AppProtocol = port.AppProtocol
		servicePort.Port = port.Port
		servicePort.TargetPort = port.TargetPort
		servicePort.NodePort = port.NodePort
		// append
		serviceDetail.Spec.Ports = append(serviceDetail.Spec.Ports, servicePort)
	}
	serviceDetail.Spec.Selector = services.Spec.Selector
	serviceDetail.Spec.ClusterIP = services.Spec.ClusterIP
	serviceDetail.Spec.ClusterIPs = services.Spec.ClusterIPs
	serviceDetail.Spec.Type = string(services.Spec.Type)
	serviceDetail.Spec.ExternalIPs = services.Spec.ExternalIPs
	serviceDetail.Spec.SessionAffinity = string(services.Spec.SessionAffinity)

	return &serviceDetail, nil
}

package ingress

import (
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/pddzl/kubeops/server/global"
	"github.com/pddzl/kubeops/server/model/kubernetes/api"
	resourceIngress "github.com/pddzl/kubeops/server/model/kubernetes/resource/ingress"
)

func (i *IngressService) GetIngressDetail(namespace string, name string) (*resourceIngress.IngressDetail, error) {
	ingress, err := global.KOP_KUBERNETES.NetworkingV1().Ingresses(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}

	// 处理ingress数据
	var ingressDetail resourceIngress.IngressDetail
	// metadata
	ingressDetail.ObjectMeta = api.NewObjectMeta(ingress.ObjectMeta)
	// spec
	ingressDetail.Spec.IngressClassName = *ingress.Spec.IngressClassName
	for _, rule := range ingress.Spec.Rules {
		var ingressRule resourceIngress.IngressRule
		ingressRule.Host = rule.Host
		for _, path := range rule.IngressRuleValue.HTTP.Paths {
			var httpIngressPath resourceIngress.HTTPIngressPath
			httpIngressPath.Path = path.Path
			httpIngressPath.PathType = string(*path.PathType)
			httpIngressPath.Backend.Service.Name = path.Backend.Service.Name
			httpIngressPath.Backend.Service.Port.Name = path.Backend.Service.Port.Name
			httpIngressPath.Backend.Service.Port.Number = path.Backend.Service.Port.Number
			ingressRule.HTTP.Paths = append(ingressRule.HTTP.Paths, httpIngressPath)
		}
		ingressDetail.Spec.Rules = append(ingressDetail.Spec.Rules, ingressRule)
	}
	// status
	for _, ig := range ingress.Status.LoadBalancer.Ingress {
		if ig.Hostname != "" {
			ingressDetail.Status.EndPoints = append(ingressDetail.Status.EndPoints, ig.Hostname)
		} else if ig.IP != "" {
			ingressDetail.Status.EndPoints = append(ingressDetail.Status.EndPoints, ig.IP)
		}
	}

	return &ingressDetail, nil
}

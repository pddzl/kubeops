package ingress

import (
	"context"
	"k8s.io/api/networking/v1"

	"github.com/pddzl/kubeops/server/global"
	"github.com/pddzl/kubeops/server/model/common/request"
	"github.com/pddzl/kubeops/server/model/kubernetes/api"
	resourceIngress "github.com/pddzl/kubeops/server/model/kubernetes/resource/ingress"
)

func (i *IngressService) GetIngressList(namespace string, info request.PageInfo) ([]resourceIngress.IngressBrief, int, error) {
	// 获取ingress原生数据
	ingressListRaw, err := global.KOP_KUBERNETES.NetworkingV1().Ingresses(namespace).List(context.TODO(), api.ListEverything)
	if err != nil {
		return nil, 0, err
	}

	end := info.PageSize * info.Page
	offset := info.PageSize * (info.Page - 1)
	var IngressBriefList []resourceIngress.IngressBrief
	var ingressList v1.IngressList

	// 分页
	total := len(ingressListRaw.Items)
	if total <= offset {
		return nil, total, nil
	}
	if total < end {
		ingressList.Items = ingressListRaw.Items[offset:]
	} else {
		ingressList.Items = ingressListRaw.Items[offset:end]
	}

	// 处理ingressListRaw数据
	for _, ingress := range ingressListRaw.Items {
		var ingressBrief resourceIngress.IngressBrief
		ingressBrief.Name = ingress.Name
		ingressBrief.NameSpace = ingress.Namespace
		ingressBrief.CreationTimestamp = ingress.CreationTimestamp
		// endPoint
		for _, ig := range ingress.Status.LoadBalancer.Ingress {
			if ig.Hostname != "" {
				ingressBrief.EndPoints = append(ingressBrief.EndPoints, ig.Hostname)
			} else if ig.IP != "" {
				ingressBrief.EndPoints = append(ingressBrief.EndPoints, ig.IP)
			}
		}
		// hosts
		set := make(map[string]struct{})
		hosts := make([]string, 0)
		for _, rule := range ingress.Spec.Rules {
			if _, exists := set[rule.Host]; !exists && len(rule.Host) > 0 {
				hosts = append(hosts, rule.Host)
			}
			set[rule.Host] = struct{}{}
		}
		ingressBrief.Hosts = hosts
		// append
		IngressBriefList = append(IngressBriefList, ingressBrief)
	}

	return IngressBriefList, total, nil
}

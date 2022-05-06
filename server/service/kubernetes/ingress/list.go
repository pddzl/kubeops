package ingress

import (
	"context"
	"k8s.io/api/networking/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/pddzl/kubeops/server/global"
	"github.com/pddzl/kubeops/server/model/common/request"
	resourceIngress "github.com/pddzl/kubeops/server/model/kubernetes/resource/ingress"
)

func (i *IngressService) GetIngressList(namespace string, info request.PageInfo) ([]resourceIngress.IngressBrief, int, error) {
	// 获取ingress list
	list, err := global.KOP_KUBERNETES.NetworkingV1().Ingresses(namespace).List(context.TODO(), metaV1.ListOptions{})
	if err != nil {
		return nil, 0, err
	}

	var ingressList v1.IngressList
	// 分页
	end := info.PageSize * info.Page
	offset := info.PageSize * (info.Page - 1)
	total := len(list.Items)
	if total <= offset {
		return nil, total, nil
	}
	if total < end {
		ingressList.Items = list.Items[offset:]
	} else {
		ingressList.Items = list.Items[offset:end]
	}

	// 处理list数据
	var IngressBriefList []resourceIngress.IngressBrief
	for _, ingress := range list.Items {
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

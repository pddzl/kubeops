package node

import (
	"context"
	"github.com/pddzl/kubeops/server/global"
	"github.com/pddzl/kubeops/server/model/kubernetes"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
)

func (n *NodeService) GetNodeDetail(nodeName string) (*kubernetes.NodeDetail, error) {
	var nodeDetail kubernetes.NodeDetail
	node, err := global.KOP_KUBERNETES.CoreV1().Nodes().Get(context.TODO(), nodeName, metaV1.GetOptions{})
	if err != nil {
		return nil, err
	}

	// ObjectMeta
	nodeDetail.ObjectMeta.Name = node.ObjectMeta.Name
	nodeDetail.ObjectMeta.Labels = node.ObjectMeta.Labels
	nodeDetail.ObjectMeta.Annotations = node.ObjectMeta.Annotations
	nodeDetail.ObjectMeta.CreationTimestamp = node.ObjectMeta.CreationTimestamp
	nodeDetail.ObjectMeta.UID = string(node.ObjectMeta.UID)

	// PodList
	podList, err := getNodePods(nodeName)
	if err != nil {
		return nil, err
	}
	for _, podRaw := range podList.Items {
		var pod kubernetes.Pod
		pod.Name = podRaw.Name
		pod.Namespace = podRaw.Namespace
		pod.Image = podRaw.Spec.Containers[0].Image
		pod.Node = podRaw.Spec.NodeName
		pod.Status = string(podRaw.Status.Phase)
		pod.CreationTimestamp = podRaw.CreationTimestamp
		// append
		nodeDetail.PodList = append(nodeDetail.PodList, pod)
	}

	// NodeAllocatedResources
	allocatedResources, err := getNodeAllocatedResources(node, podList)
	nodeDetail.NodeAllocatedResources = allocatedResources

	// PodCIDR
	nodeDetail.PodCIDR = node.Spec.PodCIDR

	// Unschedulable
	nodeDetail.Unschedulable = node.Spec.Unschedulable

	// NodeInfo
	nodeDetail.NodeInfo = node.Status.NodeInfo

	// Conditions
	for _, conditionRaw := range node.Status.Conditions {
		var condition kubernetes.Condition
		condition.Type = string(conditionRaw.Type)
		condition.Status = string(conditionRaw.Status)
		condition.LastProbeTime = conditionRaw.LastHeartbeatTime
		condition.LastTransitionTime = conditionRaw.LastTransitionTime
		condition.Reason = conditionRaw.Reason
		condition.Message = conditionRaw.Message
		// append
		nodeDetail.Conditions = append(nodeDetail.Conditions, condition)
	}

	// ContainerImages
	for _, image := range node.Status.Images {
		for _, name := range image.Names {
			// append
			nodeDetail.ContainerImages = append(nodeDetail.ContainerImages, name)
		}
	}

	// Addresses
	for _, addr := range node.Status.Addresses {
		var address kubernetes.Addresses
		address.Type = string(addr.Type)
		address.Address = addr.Address
		// append
		nodeDetail.Addresses = append(nodeDetail.Addresses, address)
	}

	// Taints
	nodeDetail.Taints = node.Spec.Taints

	return &nodeDetail, err
}

func getNodePods(nodeName string) (*v1.PodList, error) {
	fieldSelector, err := fields.ParseSelector("spec.nodeName=" + nodeName +
		",status.phase!=" + string(v1.PodSucceeded) +
		",status.phase!=" + string(v1.PodFailed))

	if err != nil {
		return nil, err
	}

	return global.KOP_KUBERNETES.CoreV1().Pods(v1.NamespaceAll).List(context.TODO(), metaV1.ListOptions{
		FieldSelector: fieldSelector.String(),
	})
}

func getNodeAllocatedResources(node *v1.Node, podList *v1.PodList) (kubernetes.NodeAllocatedResources, error) {
	reqs, limits := map[v1.ResourceName]resource.Quantity{}, map[v1.ResourceName]resource.Quantity{}

	for _, pod := range podList.Items {
		podReqs, podLimits, err := PodRequestsAndLimits(&pod)
		if err != nil {
			return kubernetes.NodeAllocatedResources{}, err
		}
		for podReqName, podReqValue := range podReqs {
			if value, ok := reqs[podReqName]; !ok {
				reqs[podReqName] = podReqValue.DeepCopy()
			} else {
				value.Add(podReqValue)
				reqs[podReqName] = value
			}
		}
		for podLimitName, podLimitValue := range podLimits {
			if value, ok := limits[podLimitName]; !ok {
				limits[podLimitName] = podLimitValue.DeepCopy()
			} else {
				value.Add(podLimitValue)
				limits[podLimitName] = value
			}
		}
	}

	cpuRequests, cpuLimits, memoryRequests, memoryLimits := reqs[v1.ResourceCPU],
		limits[v1.ResourceCPU], reqs[v1.ResourceMemory], limits[v1.ResourceMemory]

	var cpuRequestsFraction, cpuLimitsFraction float64 = 0, 0
	if capacity := float64(node.Status.Allocatable.Cpu().MilliValue()); capacity > 0 {
		cpuRequestsFraction = float64(cpuRequests.MilliValue()) / capacity * 100
		cpuLimitsFraction = float64(cpuLimits.MilliValue()) / capacity * 100
	}

	var memoryRequestsFraction, memoryLimitsFraction float64 = 0, 0
	if capacity := float64(node.Status.Allocatable.Memory().MilliValue()); capacity > 0 {
		memoryRequestsFraction = float64(memoryRequests.MilliValue()) / capacity * 100
		memoryLimitsFraction = float64(memoryLimits.MilliValue()) / capacity * 100
	}

	var podFraction float64 = 0
	var podCapacity int64 = node.Status.Capacity.Pods().Value()
	if podCapacity > 0 {
		podFraction = float64(len(podList.Items)) / float64(podCapacity) * 100
	}

	return kubernetes.NodeAllocatedResources{
		CPURequests:            cpuRequests.MilliValue(),
		CPURequestsFraction:    cpuRequestsFraction,
		CPULimits:              cpuLimits.MilliValue(),
		CPULimitsFraction:      cpuLimitsFraction,
		CPUCapacity:            node.Status.Allocatable.Cpu().MilliValue(),
		MemoryRequests:         memoryRequests.Value(),
		MemoryRequestsFraction: memoryRequestsFraction,
		MemoryLimits:           memoryLimits.Value(),
		MemoryLimitsFraction:   memoryLimitsFraction,
		MemoryCapacity:         node.Status.Allocatable.Memory().Value(),
		AllocatedPods:          len(podList.Items),
		PodCapacity:            podCapacity,
		PodFraction:            podFraction,
	}, nil
}

// PodRequestsAndLimits returns a dictionary of all defined resources summed up for all
// containers of the pod. If pod overhead is non-nil, the pod overhead is added to the
// total container resource requests and to the total container limits which have a
// non-zero quantity.
func PodRequestsAndLimits(pod *v1.Pod) (reqs, limits v1.ResourceList, err error) {
	reqs, limits = v1.ResourceList{}, v1.ResourceList{}
	for _, container := range pod.Spec.Containers {
		addResourceList(reqs, container.Resources.Requests)
		addResourceList(limits, container.Resources.Limits)
	}
	// init containers define the minimum of any resource
	for _, container := range pod.Spec.InitContainers {
		maxResourceList(reqs, container.Resources.Requests)
		maxResourceList(limits, container.Resources.Limits)
	}

	// Add overhead for running a pod to the sum of requests and to non-zero limits:
	if pod.Spec.Overhead != nil {
		addResourceList(reqs, pod.Spec.Overhead)

		for name, quantity := range pod.Spec.Overhead {
			if value, ok := limits[name]; ok && !value.IsZero() {
				value.Add(quantity)
				limits[name] = value
			}
		}
	}
	return
}

// addResourceList adds the resources in newList to list
func addResourceList(list, new v1.ResourceList) {
	for name, quantity := range new {
		if value, ok := list[name]; !ok {
			list[name] = quantity.DeepCopy()
		} else {
			value.Add(quantity)
			list[name] = value
		}
	}
}

// maxResourceList sets list to the greater of list/newList for every resource
// either list
func maxResourceList(list, new v1.ResourceList) {
	for name, quantity := range new {
		if value, ok := list[name]; !ok {
			list[name] = quantity.DeepCopy()
			continue
		} else {
			if quantity.Cmp(value) > 0 {
				list[name] = quantity.DeepCopy()
			}
		}
	}
}

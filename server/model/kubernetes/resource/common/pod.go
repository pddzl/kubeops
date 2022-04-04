package common

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// FilterPodsByControllerRef returns a subset of pods controlled by given controller resource, excluding deployments.
func FilterPodsByControllerRef(owner metav1.Object, allPods []v1.Pod) []v1.Pod {
	var matchingPods []v1.Pod
	for _, pod := range allPods {
		if metav1.IsControlledBy(&pod, owner) {
			matchingPods = append(matchingPods, pod)
		}
	}
	return matchingPods
}

package resource

import (
	"github.com/pddzl/kubeops/server/model/kubernetes/api"
	"github.com/pddzl/kubeops/server/model/kubernetes/resource/common"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ReplicaSetBrief struct {
	Name              string      `json:"name"`
	NameSpace         string      `json:"namespace"`
	Pods              string      `json:"pods"`
	CreationTimestamp metaV1.Time `json:"creationTimestamp,omitempty"`
}

type ReplicaSetDetail struct {
	ObjectMeta api.ObjectMeta        `json:"objectMeta"`
	TypeMeta   api.TypeMeta          `json:"typeMeta"`
	Selector   *metaV1.LabelSelector `json:"selector"`
	Pods       common.PodInfo        `json:"podInfo"`
}

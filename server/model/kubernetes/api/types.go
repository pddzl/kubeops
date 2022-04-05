package api

import (
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/types"
)

// ObjectMeta is metadata about an instance of a resource.
type ObjectMeta struct {
	// Name is unique within a namespace. Name is primarily intended for creation
	// idempotence and configuration definition.
	Name string `json:"name,omitempty"`

	// Namespace defines the space within which name must be unique. An empty namespace is
	// equivalent to the "default" namespace, but "default" is the canonical representation.
	// Not all objects are required to be scoped to a namespace - the value of this field for
	// those objects will be empty.
	Namespace string `json:"namespace,omitempty"`

	// Labels are key value pairs that may be used to scope and select individual resources.
	// Label keys are of the form:
	//     label-key ::= prefixed-name | name
	//     prefixed-name ::= prefix '/' name
	//     prefix ::= DNS_SUBDOMAIN
	//     name ::= DNS_LABEL
	// The prefix is optional.  If the prefix is not specified, the key is assumed to be private
	// to the user.  Other system components that wish to use labels must specify a prefix.
	Labels map[string]string `json:"labels,omitempty"`

	// Annotations are unstructured key value data stored with a resource that may be set by
	// external tooling. They are not queryable and should be preserved when modifying
	// objects.  Annotation keys have the same formatting restrictions as Label keys. See the
	// comments on Labels for details.
	Annotations map[string]string `json:"annotations,omitempty"`

	// CreationTimestamp is a timestamp representing the server time when this object was
	// created. It is not guaranteed to be set in happens-before order across separate operations.
	// Clients may not set this value. It is represented in RFC3339 form and is in UTC.
	CreationTimestamp v1.Time `json:"creationTimestamp,omitempty"`

	// UID is a type that holds unique ID values, including UUIDs.  Because we
	// don't ONLY use UUIDs, this is an alias to string.  Being a type captures
	// intent and helps make sure that UIDs and names do not get conflated.
	UID types.UID `json:"uid,omitempty"`
}

// TypeMeta describes an individual object in an API response or request with strings representing
// the type of the object.
type TypeMeta struct {
	// Kind is a string value representing the REST resource this object represents.
	// Servers may infer this from the endpoint the client submits requests to.
	// In smalllettercase.
	// More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#types-kinds
	Kind ResourceKind `json:"kind,omitempty"`

	// Scalable represents whether or not an object is scalable.
	Scalable bool `json:"scalable,omitempty"`

	// Restartable represents whether or not an object is restartable.
	Restartable bool `json:"restartable,omitempty"`
}

// ResourceKind is an unique name for each resource. It can used for API discovery and generic
// code that does things based on the kind. For example, there may be a generic "deleter"
// that based on resource kind, name and namespace deletes it.
type ResourceKind string

// ListMeta describes list of objects, i.e. holds information about pagination options set for the list.
type ListMeta struct {
	// Total number of items on the list. Used for pagination.
	TotalItems int `json:"totalItems"`
}

// NewObjectMeta returns internal endpoint name for the given service properties, e.g.,
// NewObjectMeta creates a new instance of ObjectMeta struct based on K8s object meta.
func NewObjectMeta(k8SObjectMeta metaV1.ObjectMeta) ObjectMeta {
	return ObjectMeta{
		Name:              k8SObjectMeta.Name,
		Namespace:         k8SObjectMeta.Namespace,
		Labels:            k8SObjectMeta.Labels,
		CreationTimestamp: k8SObjectMeta.CreationTimestamp,
		Annotations:       k8SObjectMeta.Annotations,
		UID:               k8SObjectMeta.UID,
	}
}

// NewTypeMeta creates new type mete for the resource kind.
func NewTypeMeta(kind ResourceKind) TypeMeta {
	return TypeMeta{
		Kind:        kind,
		Scalable:    kind.Scalable(),
		Restartable: kind.Restartable(),
	}
}

// Scalable method return whether ResourceKind is scalable.
func (k ResourceKind) Scalable() bool {
	scalable := []ResourceKind{
		ResourceKindDeployment,
		ResourceKindReplicaSet,
		ResourceKindReplicationController,
		ResourceKindStatefulSet,
	}

	for _, kind := range scalable {
		if k == kind {
			return true
		}
	}

	return false
}

// Restartable method return whether ResourceKind is restartable.
func (k ResourceKind) Restartable() bool {
	restartable := []ResourceKind{
		ResourceKindDeployment,
	}

	for _, kind := range restartable {
		if k == kind {
			return true
		}
	}

	return false
}

// List of all resource kinds supported by the UI.
const (
	ResourceKindConfigMap                = "configmap"
	ResourceKindDaemonSet                = "daemonset"
	ResourceKindDeployment               = "deployment"
	ResourceKindEvent                    = "event"
	ResourceKindHorizontalPodAutoscaler  = "horizontalpodautoscaler"
	ResourceKindIngress                  = "ingress"
	ResourceKindServiceAccount           = "serviceaccount"
	ResourceKindJob                      = "job"
	ResourceKindCronJob                  = "cronjob"
	ResourceKindLimitRange               = "limitrange"
	ResourceKindNamespace                = "namespace"
	ResourceKindNode                     = "node"
	ResourceKindPersistentVolumeClaim    = "persistentvolumeclaim"
	ResourceKindPersistentVolume         = "persistentvolume"
	ResourceKindCustomResourceDefinition = "customresourcedefinition"
	ResourceKindPod                      = "pod"
	ResourceKindReplicaSet               = "replicaset"
	ResourceKindReplicationController    = "replicationcontroller"
	ResourceKindResourceQuota            = "resourcequota"
	ResourceKindSecret                   = "secret"
	ResourceKindService                  = "service"
	ResourceKindStatefulSet              = "statefulset"
	ResourceKindStorageClass             = "storageclass"
	ResourceKindClusterRole              = "clusterrole"
	ResourceKindClusterRoleBinding       = "clusterrolebinding"
	ResourceKindRole                     = "role"
	ResourceKindRoleBinding              = "rolebinding"
	ResourceKindPlugin                   = "plugin"
	ResourceKindEndpoint                 = "endpoint"
	ResourceKindNetworkPolicy            = "networkpolicy"
)

// ListEverything is a list options used to list all resources without any filtering.
var ListEverything = metaV1.ListOptions{
	LabelSelector: labels.Everything().String(),
	FieldSelector: fields.Everything().String(),
}

// IsSelectorMatching returns true when an object with the given selector targets the same
// Resources (or subset) that the target object with the given selector.
func IsSelectorMatching(srcSelector map[string]string, targetObjectLabels map[string]string) bool {
	// If service has no selectors, then assume it targets different resource.
	if len(srcSelector) == 0 {
		return false
	}
	for label, value := range srcSelector {
		if rsValue, ok := targetObjectLabels[label]; !ok || rsValue != value {
			return false
		}
	}
	return true
}

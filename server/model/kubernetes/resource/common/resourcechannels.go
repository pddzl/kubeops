package common

import (
	"context"
	"github.com/pddzl/kubeops/server/global"
	"github.com/pddzl/kubeops/server/model/kubernetes/api"
	v1 "k8s.io/api/core/v1"
)

// ResourceChannels struct holds channels to resource lists. Each list channel is paired with
// an error channel which *must* be read sequentially: first read the list channel and then
// the error channel.
//
// This struct can be used when there are multiple clients that want to process, e.g., a
// list of pods. With this helper, the list can be read only once from the backend and
// distributed asynchronously to clients that need it.
//
// When a channel is nil, it means that no resource list is available for getting.
//
// Each channel pair can be read up to N times. N is specified upon creation of the channels.
type ResourceChannels struct {
	// List and error channels to Replication Controllers.
	//ReplicationControllerList ReplicationControllerListChannel

	// List and error channels to Replica Sets.
	//ReplicaSetList ReplicaSetListChannel

	// List and error channels to Deployments.
	//DeploymentList DeploymentListChannel

	// List and error channels to Daemon Sets.
	//DaemonSetList DaemonSetListChannel

	// List and error channels to Jobs.
	//JobList JobListChannel

	// List and error channels to Cron Jobs.
	//CronJobList CronJobListChannel

	// List and error channels to Services.
	//ServiceList ServiceListChannel

	// List and error channels to Endpoints.
	//EndpointList EndpointListChannel

	// List and error channels to Ingresses.
	//IngressList IngressListChannel

	// List and error channels to Pods.
	PodList PodListChannel

	// List and error channels to Events.
	//EventList EventListChannel

	// List and error channels to LimitRanges.
	//LimitRangeList LimitRangeListChannel

	// List and error channels to Nodes.
	//NodeList NodeListChannel

	// List and error channels to Namespaces.
	//NamespaceList NamespaceListChannel

	// List and error channels to StatefulSets.
	//StatefulSetList StatefulSetListChannel

	// List and error channels to ConfigMaps.
	ConfigMapList ConfigMapListChannel

	// List and error channels to Secrets.
	SecretList SecretListChannel

	// List and error channels to PersistentVolumes
	//PersistentVolumeList PersistentVolumeListChannel

	// List and error channels to PersistentVolumeClaims
	//PersistentVolumeClaimList PersistentVolumeClaimListChannel

	// List and error channels to ResourceQuotas
	//ResourceQuotaList ResourceQuotaListChannel

	// List and error channels to HorizontalPodAutoscalers
	//HorizontalPodAutoscalerList HorizontalPodAutoscalerListChannel

	// List and error channels to StorageClasses
	//StorageClassList StorageClassListChannel

	// List and error channels to Roles
	//RoleList RoleListChannel

	// List and error channels to ClusterRoles
	//ClusterRoleList ClusterRoleListChannel

	// List and error channels to RoleBindings
	//RoleBindingList RoleBindingListChannel

	// List and error channels to ClusterRoleBindings
	//ClusterRoleBindingList ClusterRoleBindingListChannel
}

// ConfigMapListChannel is a list and error channels to ConfigMaps.
type ConfigMapListChannel struct {
	List  chan *v1.ConfigMapList
	Error chan error
}

// GetConfigMapListChannel returns a pair of channels to a ConfigMap list and errors that both must be read
// numReads times.
func GetConfigMapListChannel(nsQuery *NamespaceQuery, numReads int) ConfigMapListChannel {

	channel := ConfigMapListChannel{
		List:  make(chan *v1.ConfigMapList, numReads),
		Error: make(chan error, numReads),
	}

	go func() {
		list, err := global.KOP_KUBERNETES.CoreV1().ConfigMaps(nsQuery.ToRequestParam()).List(context.TODO(), api.ListEverything)
		var filteredItems []v1.ConfigMap
		for _, item := range list.Items {
			if nsQuery.Matches(item.ObjectMeta.Namespace) {
				filteredItems = append(filteredItems, item)
			}
		}
		list.Items = filteredItems
		for i := 0; i < numReads; i++ {
			channel.List <- list
			channel.Error <- err
		}
	}()

	return channel
}

// SecretListChannel is a list and error channels to Secrets.
type SecretListChannel struct {
	List  chan *v1.SecretList
	Error chan error
}

// GetSecretListChannel returns a pair of channels to a Secret list and errors that
// both must be read numReads times.
func GetSecretListChannel(nsQuery *NamespaceQuery, numReads int) SecretListChannel {

	channel := SecretListChannel{
		List:  make(chan *v1.SecretList, numReads),
		Error: make(chan error, numReads),
	}

	go func() {
		list, err := global.KOP_KUBERNETES.CoreV1().Secrets(nsQuery.ToRequestParam()).List(context.TODO(), api.ListEverything)
		var filteredItems []v1.Secret
		for _, item := range list.Items {
			if nsQuery.Matches(item.ObjectMeta.Namespace) {
				filteredItems = append(filteredItems, item)
			}
		}
		list.Items = filteredItems
		for i := 0; i < numReads; i++ {
			channel.List <- list
			channel.Error <- err
		}
	}()

	return channel
}

// PodListChannel is a list and error channels to Pods.
type PodListChannel struct {
	List  chan *v1.PodList
	Error chan error
}

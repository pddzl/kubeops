package resource

import (
	"context"
	"fmt"
	"io"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/yaml"
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/dynamic"
	ctrl "sigs.k8s.io/controller-runtime"
	"strings"
)

type ResourceService struct{}

func (r *ResourceService) CreateDynamicResource(content string) error {
	reader := strings.NewReader(content)
	d := yaml.NewYAMLOrJSONDecoder(reader, 4096)

	data := &unstructured.Unstructured{}
	if err := d.Decode(data); err != nil {
		if err == io.EOF {
			return fmt.Errorf("io EOF")
		}
		return err
	}

	version := data.GetAPIVersion()
	kind := data.GetKind()

	gv, err := schema.ParseGroupVersion(version)
	if err != nil {
		gv = schema.GroupVersion{Version: version}
	}

	cfg := ctrl.GetConfigOrDie()
	discoveryClient, err := discovery.NewDiscoveryClientForConfig(cfg)
	if err != nil {
		return err
	}

	apiResourceList, err := discoveryClient.ServerResourcesForGroupVersion(version)
	if err != nil {
		return err
	}
	apiResources := apiResourceList.APIResources
	var resource *metaV1.APIResource
	for _, apiResource := range apiResources {
		if apiResource.Kind == kind && !strings.Contains(apiResource.Name, "/") {
			resource = &apiResource
			break
		}
	}
	if resource == nil {
		return fmt.Errorf("unknown resource kind: %s", kind)
	}

	dynamicClient, err := dynamic.NewForConfig(cfg)
	if err != nil {
		return err
	}

	groupVersionResource := schema.GroupVersionResource{Group: gv.Group, Version: gv.Version, Resource: resource.Name}

	namespace := data.GetNamespace()
	if resource.Namespaced {
		_, err = dynamicClient.Resource(groupVersionResource).Namespace(namespace).Create(context.TODO(), data, metaV1.CreateOptions{})
	} else {
		_, err = dynamicClient.Resource(groupVersionResource).Create(context.TODO(), data, metaV1.CreateOptions{})
	}

	return err
}

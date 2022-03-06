package initialize

import (
	"github.com/pddzl/kubeops/server/global"
	"go.uber.org/zap"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func Kubernetes() *kubernetes.Clientset {
	var err error
	var config *rest.Config

	if global.KOP_CONFIG.Kubernetes.InCluster {
		config, err = rest.InClusterConfig()
	} else {
		config, err = clientcmd.BuildConfigFromFlags("", global.KOP_CONFIG.Kubernetes.Kubeconfig)
	}
	if err != nil {
		global.KOP_LOG.Error("k8s config", zap.Error(err))
		return nil
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		global.KOP_LOG.Error("k8s clientset", zap.Error(err))
		return nil
	}

	return clientset
}

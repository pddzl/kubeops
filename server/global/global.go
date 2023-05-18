package global

import (
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"

	"github.com/pddzl/kubeops/server/config"
)

var (
	KOP_VP                  *viper.Viper
	KOP_CONFIG              config.Server
	KOP_LOG                 *zap.Logger
	KOP_DB                  *gorm.DB
	KOP_REDIS               *redis.Client
	KOP_Concurrency_Control = &singleflight.Group{}
	KOP_K8S_Client          *kubernetes.Clientset
	KOP_K8S_CONFIG          *rest.Config
)

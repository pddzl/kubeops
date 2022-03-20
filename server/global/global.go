package global

import (
	"github.com/go-redis/redis/v8"
	"github.com/songzhibin97/gkit/cache/local_cache"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"

	"github.com/pddzl/kubeops/server/config"
	"github.com/pddzl/kubeops/server/utils/timer"
)

var (
	KOP_DB     *gorm.DB
	KOP_DBList map[string]*gorm.DB
	KOP_REDIS  *redis.Client
	KOP_CONFIG config.Server
	KOP_VP     *viper.Viper
	// KOP_LOG    *oplogging.Logger
	KOP_LOG                 *zap.Logger
	KOP_Timer               timer.Timer = timer.NewTimerTask()
	KOP_Concurrency_Control             = &singleflight.Group{}
	KOP_KUBERNETES          *kubernetes.Clientset
	KUBERNETES_CONFIG       *rest.Config

	BlackCache local_cache.Cache
)

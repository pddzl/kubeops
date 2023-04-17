package initialize

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"

	"github.com/pddzl/kubeops/server/global"
)

func Redis() {
	redisCfg := global.KOP_CONFIG.Redis
	client := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", global.KOP_CONFIG.Redis.Host, global.KOP_CONFIG.Redis.Port),
		//Password: redisCfg.Password, // no password set
		DB: redisCfg.DB, // use default DB
	})
	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		global.KOP_LOG.Error("redis connect ping failed, err:", zap.Error(err))
	} else {
		global.KOP_LOG.Info("redis connect ping response:", zap.String("pong", pong))
		global.KOP_REDIS = client
	}
}

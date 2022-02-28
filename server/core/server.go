package core

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/pddzl/kubeops/server/global"
	"github.com/pddzl/kubeops/server/initialize"
	"github.com/pddzl/kubeops/server/service/system"
)

func RunServer() {
	if global.KOP_CONFIG.System.UseMultipoint || global.KOP_CONFIG.System.UseRedis {
		// 初始化redis服务
		initialize.Redis()
	}

	// 从db加载jwt数据
	if global.KOP_DB != nil {
		system.LoadAll()
	}

	// Create context that listens for the interrupt signal from the OS.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	address := fmt.Sprintf(":%d", global.KOP_CONFIG.System.Addr)
	router := initialize.Routers()
	router.Static("/form-generator", "./resource/page")
	srv := &http.Server{
		Addr:           address,
		Handler:        router,
		ReadTimeout:    300 * time.Second,
		WriteTimeout:   300 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			global.KOP_LOG.Info(fmt.Sprintf("listen: %s", err))
		}
	}()

	// Listen for the interrupt signal.
	<-ctx.Done()

	// Restore default behavior on the interrupt signal and notify user of shutdown.
	stop()
	global.KOP_LOG.Info("shutting down gracefully, press Ctrl+C again to force")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		global.KOP_LOG.Error("Server forced to shutdown: ", zap.Error(err))
		os.Exit(1)
	}
	global.KOP_LOG.Info("Server exiting")
}

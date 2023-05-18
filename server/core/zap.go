package core

import (
	"fmt"
	"github.com/pddzl/kubeops/server/core/internal"
	"github.com/pddzl/kubeops/server/global"
	"github.com/pddzl/kubeops/server/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

func Zap() (logger *zap.Logger) {
	// 判断是否有Director文件夹
	if ok, _ := utils.PathExists(global.KOP_CONFIG.Zap.Director); !ok {
		fmt.Printf("create %v directory\n", global.KOP_CONFIG.Zap.Director)
		_ = os.Mkdir(global.KOP_CONFIG.Zap.Director, os.ModePerm)
	}

	cores := internal.Zap.GetZapCores()
	logger = zap.New(zapcore.NewTee(cores...))

	if global.KOP_CONFIG.Zap.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}
	return logger
}

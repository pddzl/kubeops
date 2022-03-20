package main

import (
	"github.com/pddzl/kubeops/server/core"
	"github.com/pddzl/kubeops/server/global"
	"github.com/pddzl/kubeops/server/initialize"
	"go.uber.org/zap"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

// @title Swagger Example API
// @version 0.0.1
// @description This is a sample Server pets
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name x-token
// @BasePath /
func main() {
	global.KOP_VP = core.Viper() // 初始化Viper
	global.KOP_LOG = core.Zap()  // 初始化zap日志库
	zap.ReplaceGlobals(global.KOP_LOG)
	global.KOP_DB = initialize.Gorm()                                         // gorm连接数据库
	global.KOP_KUBERNETES, global.KUBERNETES_CONFIG = initialize.Kubernetes() // 初始化kubernetes api
	initialize.Timer()
	initialize.DBList()
	if global.KOP_DB != nil {
		//initialize.RegisterTables(global.KOP_DB) // 初始化表
		// 程序结束前关闭数据库链接
		db, _ := global.KOP_DB.DB()
		defer db.Close()
	}
	core.RunServer()
}

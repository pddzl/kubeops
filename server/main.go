package main

import (
	"go.uber.org/zap"
	"os"

	"github.com/pddzl/kubeops/server/core"
	"github.com/pddzl/kubeops/server/global"
	"github.com/pddzl/kubeops/server/initialize"
)

func main() {
	global.KOP_VP = core.Viper() // 初始化viper
	global.KOP_LOG = core.Zap()  // 初始化zap日志
	zap.ReplaceGlobals(global.KOP_LOG)
	global.KOP_DB = initialize.Gorm() // gorm连接数据库
	if global.KOP_DB == nil {
		global.KOP_LOG.Error("mysql连接失败，退出程序")
		os.Exit(127)
	} else {
		initialize.RegisterTables(global.KOP_DB) // 初始化表
		// 程序结束前关闭数据库链接
		db, _ := global.KOP_DB.DB()
		defer db.Close()
	}

	global.KOP_K8S_Client, global.KOP_K8S_CONFIG = initialize.K8sConnect()
	if global.KOP_K8S_Client == nil {
		global.KOP_LOG.Error("k8s连接失败，退出程序")
		os.Exit(128)
	}

	core.RunServer()
}

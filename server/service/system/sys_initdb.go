package system

import (
	"database/sql"
	"fmt"

	adapter "github.com/casbin/gorm-adapter/v3"
	"github.com/pddzl/kubeops/server/global"
	"github.com/pddzl/kubeops/server/model/example"
	"github.com/pddzl/kubeops/server/model/system"
	"github.com/pddzl/kubeops/server/model/system/request"
)

type InitDBService struct{}

// InitDB 创建数据库并初始化 总入口
func (initDBService *InitDBService) InitDB(conf request.InitDB) error {
	switch conf.DBType {
	case "mysql":
		return initDBService.initMsqlDB(conf)
	default:
		return initDBService.initMsqlDB(conf)
	}
}

// initTables 初始化表
func (initDBService *InitDBService) initTables() error {
	return global.KOP_DB.AutoMigrate(
		system.SysApi{},
		system.SysUser{},
		system.SysBaseMenu{},
		system.SysAuthority{},
		system.JwtBlacklist{},
		system.SysOperationRecord{},
		system.SysBaseMenuParameter{},
		adapter.CasbinRule{},

		example.ExaFileUploadAndDownload{},
	)
}

// createDatabase 创建数据库(mysql)
func (initDBService *InitDBService) createDatabase(dsn string, driver string, createSql string) error {
	db, err := sql.Open(driver, dsn)
	if err != nil {
		return err
	}
	defer func(db *sql.DB) {
		err = db.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(db)
	if err = db.Ping(); err != nil {
		return err
	}
	_, err = db.Exec(createSql)
	return err
}

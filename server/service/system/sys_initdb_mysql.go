package system

import (
	"fmt"
	"github.com/pddzl/kubeops/server/config"
	"github.com/pddzl/kubeops/server/global"
	model "github.com/pddzl/kubeops/server/model/system"
	"github.com/pddzl/kubeops/server/model/system/request"
	"github.com/pddzl/kubeops/server/source/example"
	"github.com/pddzl/kubeops/server/source/system"
	"github.com/pddzl/kubeops/server/utils"
	uuid "github.com/satori/go.uuid"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// writeMysqlConfig mysql回写配置
func (initDBService *InitDBService) writeMysqlConfig(mysql config.Mysql) error {
	global.KOP_CONFIG.Mysql = mysql
	cs := utils.StructToMap(global.KOP_CONFIG)
	for k, v := range cs {
		global.KOP_VP.Set(k, v)
	}
	global.KOP_VP.Set("jwt.signing-key", uuid.NewV4().String())
	return global.KOP_VP.WriteConfig()
}

// initMsqlDB 创建数据库并初始化 mysql
func (initDBService *InitDBService) initMsqlDB(conf request.InitDB) error {
	dsn := conf.MysqlEmptyDsn()
	createSql := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS `%s` DEFAULT CHARACTER SET utf8mb4 DEFAULT COLLATE utf8mb4_general_ci;", conf.DBName)
	if err := initDBService.createDatabase(dsn, "mysql", createSql); err != nil {
		return err
	} // 创建数据库

	mysqlConfig := conf.ToMysqlConfig()
	if mysqlConfig.Dbname == "" {
		return nil
	} // 如果没有数据库名, 则跳出初始化数据

	if db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       mysqlConfig.Dsn(), // DSN data source name
		DefaultStringSize:         191,               // string 类型字段的默认长度
		SkipInitializeWithVersion: true,              // 根据版本自动配置
	}), &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true}); err != nil {
		return nil
	} else {
		global.KOP_DB = db
	}

	if err := initDBService.initTables(); err != nil {
		global.KOP_DB = nil
		return err
	}

	if err := initDBService.initMysqlData(); err != nil {
		global.KOP_DB = nil
		return err
	}

	if err := initDBService.writeMysqlConfig(mysqlConfig); err != nil {
		return err
	}

	//global.KOP_CONFIG.AutoCode.Root, _ = filepath.Abs("..")
	return nil
}

// initData mysql 初始化数据
func (initDBService *InitDBService) initMysqlData() error {
	return model.MysqlDataInitialize(
		system.Api,
		system.User,
		system.Casbin,
		system.BaseMenu,
		system.Authority,
		system.UserAuthority,
		system.DataAuthorities,
		system.AuthoritiesMenus,
		system.ViewAuthorityMenuMysql,
		example.FileMysql,
	)
}

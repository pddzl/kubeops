package example

import (
	"github.com/pddzl/kubeops/server/global"
	"github.com/pddzl/kubeops/server/model/example"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

var FileMysql = new(fileMysql)

type fileMysql struct{}

func (f *fileMysql) TableName() string {
	return "exa_file_upload_and_downloads"
}

func (f *fileMysql) Initialize() error {
	entities := []example.ExaFileUploadAndDownload{
		{Name: "doge.png", Url: "http://127.0.0.1:8888/uploads/file/doge.png", Tag: "png", Key: "doge.png"},
	}
	if err := global.KOP_DB.Create(&entities).Error; err != nil {
		return errors.Wrap(err, f.TableName()+"表数据初始化失败!")
	}
	return nil
}

func (f *fileMysql) CheckDataExist() bool {
	if errors.Is(global.KOP_DB.Where("`name` = ? AND `key` = ?", "doge.png", "doge.png").First(&example.ExaFileUploadAndDownload{}).Error, gorm.ErrRecordNotFound) {
		return false
	}
	return true
}

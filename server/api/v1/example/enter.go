package example

import (
	"github.com/pddzl/kubeops/server/service"
)

type ApiGroup struct {
	FileUploadAndDownloadApi
}

type FileUploadAndDownloadApi struct{}

var (
	fileUploadAndDownloadService = service.ServiceGroupApp.ExampleServiceGroup.FileUploadAndDownloadService
)

package response

import (
	"github.com/pddzl/kubeops/server/model/example"
)

type ExaFileResponse struct {
	File example.ExaFileUploadAndDownload `json:"file"`
}

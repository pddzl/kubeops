package response

import "github.com/pddzl/kubeops/server/config"

type SysConfigResponse struct {
	Config config.Server `json:"config"`
}

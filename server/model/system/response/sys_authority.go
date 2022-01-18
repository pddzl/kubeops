package response

import "github.com/pddzl/kubeops/server/model/system"

type SysAuthorityResponse struct {
	Authority system.SysAuthority `json:"authority"`
}

type SysAuthorityCopyResponse struct {
	Authority      system.SysAuthority `json:"authority"`
	OldAuthorityId string              `json:"oldAuthorityId" validate:"required"` // 旧角色ID
}

package request

type ClusterRoleCommon struct {
	ClusterRole string `json:"clusterRole" validate:"required"`
}

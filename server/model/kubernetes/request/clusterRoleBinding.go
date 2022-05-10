package request

type ClusterRoleBindingCommon struct {
	ClusterRoleBinding string `json:"cluster_role_binding" validated:"required"`
}

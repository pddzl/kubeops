package request

type RoleBindingCommon struct {
	Namespace   string `json:"namespace" validate:"required"`
	RoleBinding string `json:"roleBinding" validate:"required"`
}

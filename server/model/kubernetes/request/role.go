package request

type RoleCommon struct {
	Namespace string `json:"namespace" validate:"required"`
	Role      string `json:"role" validate:"required"`
}

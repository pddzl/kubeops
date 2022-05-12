package request

type SecretCommon struct {
	Namespace string `json:"namespace" validate:"required"`
	Secret    string `json:"secret" validate:"required"`
}

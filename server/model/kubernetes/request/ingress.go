package request

type IngressCommon struct {
	NameSpace string `json:"namespace" validate:"required"`
	Ingress   string `json:"ingress" validate:"required"`
}

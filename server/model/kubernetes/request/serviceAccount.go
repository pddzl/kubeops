package request

type ServiceAccountCommon struct {
	NameSpace      string `json:"namespace" validate:"required"`
	ServiceAccount string `json:"serviceAccount" validated:"required"`
}

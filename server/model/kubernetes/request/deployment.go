package request

type DeploymentCommon struct {
	NameSpace  string `json:"namespace" validate:"required"`
	Deployment string `json:"deployment" validate:"required"`
}

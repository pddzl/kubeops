package request

type ConfigMapCommon struct {
	Namespace string `json:"namespace" validate:"required"`
	ConfigMap string `json:"configMap" validate:"required"`
}

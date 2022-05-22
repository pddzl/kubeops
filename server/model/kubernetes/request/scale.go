package request

type ScaleCommon struct {
	Namespace string `json:"namespace" validate:"required"`
	Name      string `json:"name" validate:"required"`
	Kind      string `json:"kind" validate:"oneof=replicaSet deployment"`
	Num       int32  `json:"num" validate:"min=0,max=50"`
}

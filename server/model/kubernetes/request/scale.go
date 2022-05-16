package request

type ScaleCommon struct {
	Namespace string `json:"namespace" validated:"required"`
	Name      string `json:"name" validated:"required"`
	Kind      string `json:"kind" validated:"oneof=replicaset deployment"`
	Num       int32  `json:"num" validated:"min=0,max=50"`
}

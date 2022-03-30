package request

type ReplicaSetInRaw struct {
	NameSpace  string `json:"namespace" validate:"required"`
	ReplicaSet string `json:"replicaSet" validate:"required"`
}

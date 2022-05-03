package request

type JobCommon struct {
	NameSpace string `json:"namespace" validate:"required"`
	Job       string `json:"job" validate:"required"`
}

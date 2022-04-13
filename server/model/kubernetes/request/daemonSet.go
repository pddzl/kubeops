package request

type DaemonSetCommon struct {
	NameSpace string `json:"namespace" validate:"required"`
	DaemonSet string `json:"daemonSet" validate:"required"`
}

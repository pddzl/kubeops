package resource

import (
	api "k8s.io/api/core/v1"
)

func ToLimitRanges(rawLimitRange *api.LimitRange) (limitRangeDetail LimitRangeDetail) {
	limitRangeDetail.Name = rawLimitRange.Name
	limitRangeDetail.LimitRangeItem = rawLimitRange.Spec.Limits
	return
}

type LimitRangeDetail struct {
	Name           string               `json:"name"`
	LimitRangeItem []api.LimitRangeItem `json:"limits"`
}

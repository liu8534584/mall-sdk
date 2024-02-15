package buyin_doukeCrowdMatch_response

import (
	"doudian.com/open/sdk_golang/core"
)

type BuyinDoukeCrowdMatchResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *BuyinDoukeCrowdMatchData `json:"data"`
}
type BuyinDoukeCrowdMatchData struct {
	// 是否为潜客
	IsPotentialCustomer bool `json:"is_potential_customer"`
	// 是否为流失用户
	IsLostCustomer bool `json:"is_lost_customer"`
}

package buyin_doukeRewardOrders_response

import (
	"doudian.com/open/sdk_golang/core"
)

type BuyinDoukeRewardOrdersResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *BuyinDoukeRewardOrdersData `json:"data"`
}
type RewardOrdersItem struct {
	//跟单参数
	ExternalInfo string `json:"external_info"`
	// 成交日期
	PayDate string `json:"pay_date"`
	// 推广位信息
	PromotionInfo string `json:"promotion_info"`
	// 奖励订单ID
	RewardOrderId string `json:"reward_order_id"`
	// 成交商品ID
	ProductId string `json:"product_id"`
	// 成交商品名称
	ProductName string `json:"product_name"`
	// 成交金额
	PayAmount int64 `json:"pay_amount"`
	// 奖励类型：拉新/召回
	RewardType string `json:"reward_type"`
}
type BuyinDoukeRewardOrdersData struct {
	// 奖励订单信息
	RewardOrders []RewardOrdersItem `json:"reward_orders"`
	// 查询的奖励订单总数
	Total int64 `json:"total"`
}

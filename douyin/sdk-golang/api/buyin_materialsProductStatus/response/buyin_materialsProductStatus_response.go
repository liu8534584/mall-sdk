package buyin_materialsProductStatus_response

import (
	"doudian.com/open/sdk_golang/core"
)

type BuyinMaterialsProductStatusResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *BuyinMaterialsProductStatusData `json:"data"`
}
type DataItem struct {
	// 商品 url
	ProductUrl string `json:"product_url"`
	// 上下架状态 -2:彻底删除，-1:创建成功，提审失败（废弃），0:在线，1:下线，2:删除
	Status int16 `json:"status"`
	// 是否加入精选联盟  true：加入精选联盟，false：未加入精选联盟
	JoinAlliance bool `json:"join_alliance"`
	// 推广状态  1:终止推广  2.开启推广 3.商家主动关闭 4.商家修改关闭  6.平台治理关闭
	PromotionStatus int64 `json:"promotion_status"`
	// 商品是否可分销
	CanShare bool `json:"can_share"`
}
type BuyinMaterialsProductStatusData struct {
	// 返回数据
	Data []DataItem `json:"data"`
}

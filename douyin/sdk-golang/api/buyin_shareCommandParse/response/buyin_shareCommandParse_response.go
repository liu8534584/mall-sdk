package buyin_shareCommandParse_response

import (
	"doudian.com/open/sdk_golang/core"
)

type BuyinShareCommandParseResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *BuyinShareCommandParseData `json:"data"`
}
type ProductInfo struct {
	// 商品ID
	ProductId int64 `json:"product_id"`
	// 商品名称
	Title string `json:"title"`
	// 商品售价（单位为分）
	Price int64 `json:"price"`
	// 普通佣金比例（乘100，例如10%为10）
	CosRatio float64 `json:"cos_ratio"`
	// 普通佣金金额（单位为分）
	CosFee int64 `json:"cos_fee"`
	// 商品链接
	DetailUrl string `json:"detail_url"`
	// 是否可推广（加入联盟且是可推广状态）
	Promotable bool `json:"promotable"`
}
type BuyinShareCommandParseData struct {
	// 商品信息
	ProductInfo ProductInfo `json:"product_info"`
}

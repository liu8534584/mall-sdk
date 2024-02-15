package buyin_productSkus_response

import (
	"doudian.com/open/sdk_golang/core"
)

type BuyinProductSkusResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *BuyinProductSkusData `json:"data"`
}
type SpecItemsItem struct {
	// sku item id
	Id string `json:"id"`
	// sku item 名字
	Name string `json:"name"`
}
type SpecsItem struct {
	// sku item 名字
	Name string `json:"name"`
	// sku item 信息
	SpecItems []SpecItemsItem `json:"spec_items"`
}
type PicturesItem struct {
	// 小图
	LittlePicture string `json:"little_picture"`
	// 大图
	BigPicture string `json:"big_picture"`
}
type SkusItem struct {
	// 价格（单位：分）
	EffectivePrice int64 `json:"effective_price"`
	// 库存
	StockNum int64 `json:"stock_num"`
}
type BuyinProductSkusData struct {
	// 商品 sku 所有信息
	Specs []SpecsItem `json:"specs"`
	// 图片
	Pictures map[string]PicturesItem `json:"pictures"`
	// sku 信息
	Skus map[string]SkusItem `json:"skus"`
}

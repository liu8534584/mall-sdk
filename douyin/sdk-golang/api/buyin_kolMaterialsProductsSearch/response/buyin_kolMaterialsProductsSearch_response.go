package buyin_kolMaterialsProductsSearch_response

import (
	"doudian.com/open/sdk_golang/core"
)

type BuyinKolMaterialsProductsSearchResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *BuyinKolMaterialsProductsSearchData `json:"data"`
}
type BuyinKolMaterialsProductsSearchData struct {
	// 商品总数
	Total int64 `json:"total"`
	// 商品列表
	Products []ProductsItem `json:"products"`
}
type ProductsItem struct {
	// 商品id
	ProductId int64 `json:"product_id"`
	// 商品标题
	Title string `json:"title"`
	// 商品售价（单位为分）
	Price int64 `json:"price"`
	// 普通佣金比例（乘100，例如10%为10）
	CosRatio float64 `json:"cos_ratio"`
	// 普通佣金金额（单位为分）
	CosFee int64 `json:"cos_fee"`
	// 商品一级类目
	FirstCid int64 `json:"first_cid"`
	// 商品二级类目
	SecondCid int64 `json:"second_cid"`
	// 商品三级类目
	ThirdCid int64 `json:"third_cid"`
	// 是否有库存
	InStock bool `json:"in_stock"`
	// 历史总销量
	Sales int64 `json:"sales"`
	// 商品主图
	Cover string `json:"cover"`
	// 商品链接（普通计划链接）
	DetailUrl string `json:"detail_url"`
	// 店铺ID
	ShopId int64 `json:"shop_id"`
	// 店铺名称
	ShopName string `json:"shop_name"`
	// 券后价（单位：分）
	CouponPrice int64 `json:"coupon_price"`
	// 是否可分销
	Sharable bool `json:"sharable"`
	// 达人佣金比例
	KolCosRatio float64 `json:"kol_cos_ratio"`
	// 达人佣金金额（单位为分）
	KolCosFee int64 `json:"kol_cos_fee"`
}

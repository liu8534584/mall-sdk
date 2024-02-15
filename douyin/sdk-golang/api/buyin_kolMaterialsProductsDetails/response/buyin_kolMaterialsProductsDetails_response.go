package buyin_kolMaterialsProductsDetails_response

import (
	"doudian.com/open/sdk_golang/core"
)

type BuyinKolMaterialsProductsDetailsResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *BuyinKolMaterialsProductsDetailsData `json:"data"`
}
type ServiceScore struct {
	// 文本
	Text string `json:"text"`
	// 得分
	Score string `json:"score"`
	// 等级（1:高 2:中 3:低）
	Level int16 `json:"level"`
}
type ShopTotalScore struct {
	// 商家体验分
	ShopScore ShopScore `json:"shop_score"`
	// 商品体验分
	ProductScore ProductScore `json:"product_score"`
	// 物流体验分
	LogisticsScore LogisticsScore `json:"logistics_score"`
	// 商家服务分
	ServiceScore ServiceScore `json:"service_score"`
}
type ProductsItem struct {
	// 券后价
	CouponPrice int64 `json:"coupon_price"`
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
	// 销量
	Sales int64 `json:"sales"`
	// 商品主图
	Cover string `json:"cover"`
	// 商品轮播图
	Imgs []string `json:"imgs"`
	// 商品链接（普通计划链接）
	DetailUrl string `json:"detail_url"`
	// 店铺ID
	ShopId int64 `json:"shop_id"`
	// 店铺名称
	ShopName string `json:"shop_name"`
	// 商品活动审核状态
	Status int32 `json:"status"`
	// 活动审核原因
	Reason string `json:"reason"`
	// 达人佣金比例
	KolCosRatio float64 `json:"kol_cos_ratio"`
	// 达人佣金金额（单位为分）
	KolCosFee int64 `json:"kol_cos_fee"`
	// 商品评分（5分制，保留一位）
	CommentScore float64 `json:"comment_score"`
	// 商品评价数目
	CommentNum int64 `json:"comment_num"`
	// 近30天商品总订单量
	OrderNum int64 `json:"order_num"`
	// 近30天商品总浏览量
	ViewNum int64 `json:"view_num"`
	// 近30天推广总达人数
	KolNum int64 `json:"kol_num"`
	// 近30天推广达人数、浏览量、和订单量明细
	DailyStatistics []DailyStatisticsItem `json:"daily_statistics"`
	// 专属团长功能活动对应的限制类目名称
	CategoryName string `json:"category_name"`
	// 专属团长功能活动对应的限制类目名称id
	CategoryId int64 `json:"category_id"`
	// 是否提供安心购服务
	IsAssured bool `json:"is_assured"`
	// 商品物流说明
	LogisticsInfo string `json:"logistics_info"`
	// 是否具有短视频随心推资质
	HasSxt bool `json:"has_sxt"`
	// 商家得分
	ShopTotalScore ShopTotalScore `json:"shop_total_score"`
	// 是否可分销
	Sharable bool `json:"sharable"`
	// 商品参与活动id。0:未参加活动；1: 超值购
	ActivityId int `json:"activity_id"`
}
type BuyinKolMaterialsProductsDetailsData struct {
	// 商品总数
	Total int64 `json:"total"`
	// 商品列表
	Products []ProductsItem `json:"products"`
}
type DailyStatisticsItem struct {
	// 日期
	Date string `json:"date"`
	// 当日商品订单量明细
	OrderNum int64 `json:"order_num"`
	// 当日商品浏览量明细
	ViewNum int64 `json:"view_num"`
	// 当日推广达人数明细
	KolNum int64 `json:"kol_num"`
}
type ShopScore struct {
	// 文本
	Text string `json:"text"`
	// 得分
	Score string `json:"score"`
	// 等级（1:高 2:中 3:低）
	Level int16 `json:"level"`
}
type ProductScore struct {
	// 文本
	Text string `json:"text"`
	// 得分
	Score string `json:"score"`
	// 等级（1:高 2:中 3:低）
	Level int16 `json:"level"`
}
type LogisticsScore struct {
	// 文本
	Text string `json:"text"`
	// 得分
	Score string `json:"score"`
	// 等级（1:高 2:中 3:低）
	Level int16 `json:"level"`
}

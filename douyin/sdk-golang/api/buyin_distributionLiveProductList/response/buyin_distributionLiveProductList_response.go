package buyin_distributionLiveProductList_response

import (
	"doudian.com/open/sdk_golang/core"
)

type BuyinDistributionLiveProductListResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *BuyinDistributionLiveProductListData `json:"data"`
}
type GivenSkusItem struct {
	// 商品ID
	ProductId int64 `json:"product_id"`
	// SKU ID
	SkuId int64 `json:"sku_id"`
	// sku主图
	Cover string `json:"cover"`
}
type GivenProductsItem struct {
	// 商品ID
	ProductId int64 `json:"product_id"`
	// 商品名称
	ProductName string `json:"product_name"`
	// 商品主图
	Cover string `json:"cover"`
	// SKU列表
	GivenSkus []GivenSkusItem `json:"given_skus"`
}
type LotteryProductsItem struct {
	// 福袋商品
	ProductInfo *ProductInfo `json:"product_info"`
	// 福袋id
	LotteryActivityId int64 `json:"lottery_activity_id"`
	// 福袋活动开始时间
	LotteryProductStartTime int64 `json:"lottery_product_start_time"`
	// 福袋活动结束时间/开奖时间
	LotteryProductEndTime int64 `json:"lottery_product_end_time"`
}
type AvailableCouponsItem struct {
	// 优惠券类型：1 平台券；2 店铺券；3 达人券
	CouponType int32 `json:"coupon_type"`
	// 券类型描述
	TypeDesc string `json:"type_desc"`
	// 券描述
	DiscountDesc string `json:"discount_desc"`
}
type ProductsItem struct {
	// 商品在列表序号
	ListNum int32 `json:"list_num"`
	// 商品ID
	ProductId int64 `json:"product_id"`
	// 商品名称
	Title string `json:"title"`
	// 商品价格
	Price int64 `json:"price"`
	// 一级类目
	FirstCid int64 `json:"first_cid"`
	// 一级类目名称
	FirstCname string `json:"first_cname"`
	// 二级类目
	SecondCid int64 `json:"second_cid"`
	// 二级类目名称
	SecondCname string `json:"second_cname"`
	// 三级类目
	ThirdCid int64 `json:"third_cid"`
	// 三级类目名称
	ThirdCname string `json:"third_cname"`
	// 行业类目ID
	CategoryId int64 `json:"category_id"`
	// 直播间价格
	LiveRoomPrice int64 `json:"live_room_price"`
	// 是否在讲解中
	IsExplaining bool `json:"is_explaining"`
	// 优惠券列表
	AvailableCoupons []AvailableCouponsItem `json:"available_coupons"`
	// 渠道推广费率（X10000）
	AdsPromotionRatio *AdsPromotionRatio `json:"ads_promotion_ratio"`
	// 赠品信息
	GivenProducts []GivenProductsItem `json:"given_products"`
	// 商品主图
	Cover string `json:"cover"`
}
type FlashProductsItem struct {
	// 商品在列表序号
	ListNum int32 `json:"list_num"`
	// 商品ID
	ProductId int64 `json:"product_id"`
	// 商品名称
	Title string `json:"title"`
	// 商品价格
	Price int64 `json:"price"`
	// 一级类目
	FirstCid int64 `json:"first_cid"`
	// 一级类目名称
	FirstCname string `json:"first_cname"`
	// 二级类目
	SecondCid int64 `json:"second_cid"`
	// 二级类目名称
	SecondCname string `json:"second_cname"`
	// 三级类目
	ThirdCid int64 `json:"third_cid"`
	// 三级类目名称
	ThirdCname string `json:"third_cname"`
	// 行业类目ID
	CategoryId int64 `json:"category_id"`
	// 直播间价格
	LiveRoomPrice int64 `json:"live_room_price"`
	// 是否在讲解中
	IsExplaining bool `json:"is_explaining"`
	// 优惠券列表
	AvailableCoupons []AvailableCouponsItem `json:"available_coupons"`
	// 渠道推广费率（X10000）
	AdsPromotionRatio *AdsPromotionRatio `json:"ads_promotion_ratio"`
	// 赠品信息
	GivenProducts []GivenProductsItem `json:"given_products"`
	// 主图
	Cover string `json:"cover"`
}
type ProductInfo struct {
	// 商品序号
	ListNum int32 `json:"list_num"`
	// 商品id
	ProductId int64 `json:"product_id"`
	// 商品名称
	Title string `json:"title"`
	// 商品价格
	Price int64 `json:"price"`
	// 一级类目
	FirstCid int64 `json:"first_cid"`
	// 一级类目名称
	FirstCname string `json:"first_cname"`
	// 二级类目
	SecondCid int64 `json:"second_cid"`
	// 二级类目名称
	SecondCname string `json:"second_cname"`
	// 三级类目
	ThirdCid int64 `json:"third_cid"`
	// 三级类目名称
	ThirdCname string `json:"third_cname"`
	// 行业类目ID
	CategoryId int64 `json:"category_id"`
	// 直播间价格
	LiveRoomPrice int64 `json:"live_room_price"`
	// 是否在讲解中
	IsExplaining bool `json:"is_explaining"`
	// 优惠券
	AvailableCoupons []AvailableCouponsItem `json:"available_coupons"`
	// 渠道推广费率（X10000）
	AdsPromotionRatio *AdsPromotionRatio `json:"ads_promotion_ratio"`
	// 赠品信息
	GivenProducts []GivenProductsItem `json:"given_products"`
	// 主图
	Cover string `json:"cover"`
}
type BuyinDistributionLiveProductListData struct {
	// 直播间创建时间
	CreateTime int64 `json:"create_time"`
	// 在线人数
	OnlineNum int64 `json:"online_num"`
	// 是否支持安心购
	IsRestAssured bool `json:"is_rest_assured"`
	// 是否有主播券
	HasKolCoupon bool `json:"has_kol_coupon"`
	// 是否闪购直播间
	IsFlashLiveRoom bool `json:"is_flash_live_room"`
	// 商品总数（不含闪购），最多100
	ProductCount int32 `json:"product_count"`
	// 商品列表（不含闪购）
	Products []ProductsItem `json:"products"`
	// 闪购商品总数
	FlashProductsCount int32 `json:"flash_products_count"`
	// 闪购商品列表
	FlashProducts []FlashProductsItem `json:"flash_products"`
	// 直播间ID
	RoomId int64 `json:"room_id"`
	// 昵称
	AuthorNickname string `json:"author_nickname"`
	// 百应id
	AuthorBuyinId int64 `json:"author_buyin_id"`
	// 福袋商品
	LotteryProducts []LotteryProductsItem `json:"lottery_products"`
}
type AdsPromotionRatio struct {
	// 新客推广费率
	ShareRatio int64 `json:"share_ratio"`
	// 老客推广费率
	OldFansRatio int64 `json:"old_fans_ratio"`
}

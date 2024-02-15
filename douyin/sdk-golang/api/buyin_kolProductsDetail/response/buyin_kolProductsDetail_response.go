package buyin_kolProductsDetail_response

import (
	"doudian.com/open/sdk_golang/core"
)

type BuyinKolProductsDetailResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *BuyinKolProductsDetailData `json:"data"`
}
type ShopScore struct {
	// 文本
	Text string `json:"text"`
	// 得分
	Score string `json:"score"`
	// 等级（1:高 2:中 3:低）
	Level int16 `json:"level"`
}
type MonthSaleData struct {
	// 近30天商品总订单量
	OrderNum int64 `json:"order_num"`
	// 近30天商品总浏览量
	ViewNum int64 `json:"view_num"`
	// 近30天推广总达人数
	KolNum int64 `json:"kol_num"`
	// 近30天推广达人数、浏览量、和订单量明细
	DailyStatistics []DailyStatisticsItem `json:"daily_statistics"`
	// 数据是否可用
	Success bool `json:"success"`
}
type BrandInfo struct {
	// 品牌 id
	BrandId int64 `json:"brand_id"`
	// 品牌名称
	BrandNameCn string `json:"brand_name_cn"`
	// 品牌英文名
	BrandNameEn string `json:"brand_name_en"`
	// 数据是否可用
	Success bool `json:"success"`
}
type ActivityInfo struct {
	// 活动类型（0:未参加活动；1: 超值购）
	ActivityType int64 `json:"activity_type"`
	// 数据是否可用
	Success bool `json:"success"`
}
type AvailableCouponsItem struct {
	// 优惠券类型：1 平台券 2 店铺券 3 主播券
	CouponType int32 `json:"coupon_type"`
	// 优惠券类型描述：平台券/主播券/店铺券
	TypeDesc string `json:"type_desc"`
	// 优惠券内容描述
	DiscountDesc string `json:"discount_desc"`
	// 优惠券领取开始时间
	ApplyStartTime string `json:"apply_start_time"`
	// 优惠券领取结束时间
	ApplyEndTime string `json:"apply_end_time"`
	// 优惠券有效期类型：1固定有效期类型，2浮动有效期类型
	ValidityType int64 `json:"validity_type"`
	// 1固定有效期类型，优惠券使用开始时间
	UseStartTime string `json:"use_start_time"`
	// 1固定有效期类型，优惠券使用结束时间
	UseEndTime string `json:"use_end_time"`
	// 2浮动有效期类型，领取优惠券后有效期，单位s
	ValidPeriod int64 `json:"valid_period"`
}
type PresellInfo struct {
	// 预售类型。0:非预售，1;全款预售，2:阶梯库存
	PresellType int64 `json:"presell_type"`
	// 数据是否可用
	Success bool `json:"success"`
}
type RightsInfo struct {
	// 是否提供安心购服务
	IsAssured bool `json:"is_assured"`
	// 数据是否可用
	Success bool `json:"success"`
}
type LogisticsInfo struct {
	// 商品物流说明
	Text string `json:"text"`
	// 数据是否可用
	Success bool `json:"success"`
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
type ShareInfo struct {
	// 是否可分销
	Sharable bool `json:"sharable"`
	// 数据是否可用
	Success bool `json:"success"`
}
type Tags struct {
	// 是否有【抖in好物】标签
	HasDouinGoodsTag string `json:"has_douin_goods_tag"`
	// 是否有品牌旗舰店标签（[品牌]黑标）
	HasShopBrandTag bool `json:"has_shop_brand_tag"`
	// 数据是否可用
	Success bool `json:"success"`
}
type QualificationInfo struct {
	// 是否具有短视频随心推资质
	HasSxt bool `json:"has_sxt"`
	// 数据是否可用
	Success bool `json:"success"`
}
type ShopInfo struct {
	// 店铺 id
	ShopId int64 `json:"shop_id"`
	// 店铺名称
	ShopName string `json:"shop_name"`
	// 商家得分
	ShopTotalScore *ShopTotalScore `json:"shop_total_score"`
	// 数据是否可用
	Success bool `json:"success"`
}
type BuyinKolProductsDetailData struct {
	// 商品列表
	Products []ProductsItem `json:"products"`
}
type PromotionInfo struct {
	// 普通佣金比例（10%返回10）
	CosRatio float64 `json:"cos_ratio"`
	// 普通佣金金额（单位为分）
	CosFee int64 `json:"cos_fee"`
	// 达人佣金比例
	KolCosRatio float64 `json:"kol_cos_ratio"`
	// 达人佣金金额（单位为分）
	KolCosFee int64 `json:"kol_cos_fee"`
	// 佣金类型（1-团长普通模式，2-团长竞价模式，3-普通佣金，4-商品定向佣金，5-店铺定向佣金，6-提报活动，7-招募佣金，8-托管团长，9-阶梯佣金）
	CommissionType int32 `json:"commission_type"`
	// 数据是否可用
	Success bool `json:"success"`
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
	ShopScore *ShopScore `json:"shop_score"`
	// 商品体验分
	ProductScore *ProductScore `json:"product_score"`
	// 物流体验分
	LogisticsScore *LogisticsScore `json:"logistics_score"`
	// 商家服务分
	ServiceScore *ServiceScore `json:"service_score"`
}
type BaseInfo struct {
	// 商品 id
	ProductId int64 `json:"product_id"`
	// 商品标题
	Title string `json:"title"`
	// 商品价格（单位：分）
	Price int64 `json:"price"`
	// 商品一级类目
	FirstCid int64 `json:"first_cid"`
	// 商品二级类目
	SecondCid int64 `json:"second_cid"`
	// 商品三级类目
	ThirdCid int64 `json:"third_cid"`
	// 是否有库存
	InStock bool `json:"in_stock"`
	// 总销量
	Sales int64 `json:"sales"`
	// 商品主图
	Cover string `json:"cover"`
	// 商品轮播图
	Imgs []string `json:"imgs"`
	// 商品链接
	DetailUrl string `json:"detail_url"`
	// 行业类目名称
	CategoryName string `json:"category_name"`
	// 行业类目 id
	CategoryId int64 `json:"category_id"`
	// 数据是否可用
	Success bool `json:"success"`
}
type CommentInfo struct {
	// 商品评分（5分制，保留一位）
	CommentScore float64 `json:"comment_score"`
	// 商品评价数目
	CommentNum int64 `json:"comment_num"`
	// 数据是否可用
	Success bool `json:"success"`
}
type CouponInfo struct {
	// 优惠券列表
	AvailableCoupons []AvailableCouponsItem `json:"available_coupons"`
	// 券后价（单位：分）
	CouponPrice int64 `json:"coupon_price"`
	// 数据是否可用
	Success bool `json:"success"`
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
type ProductsItem struct {
	// 基础信息
	BaseInfo *BaseInfo `json:"base_info"`
	// 推广信息
	PromotionInfo *PromotionInfo `json:"promotion_info"`
	// 评论信息
	CommentInfo *CommentInfo `json:"comment_info"`
	// 月售信息
	MonthSaleData *MonthSaleData `json:"month_sale_data"`
	// 分销信息
	ShareInfo *ShareInfo `json:"share_info"`
	// 品牌信息
	BrandInfo *BrandInfo `json:"brand_info"`
	// 商品标签
	Tags *Tags `json:"tags"`
	// 活动信息
	ActivityInfo *ActivityInfo `json:"activity_info"`
	// 券信息
	CouponInfo *CouponInfo `json:"coupon_info"`
	// 预售信息
	PresellInfo *PresellInfo `json:"presell_info"`
	// 权益信息
	RightsInfo *RightsInfo `json:"rights_info"`
	// 物流信息
	LogisticsInfo *LogisticsInfo `json:"logistics_info"`
	// 资质信息
	QualificationInfo *QualificationInfo `json:"qualification_info"`
	// 店铺信息
	ShopInfo *ShopInfo `json:"shop_info"`
}

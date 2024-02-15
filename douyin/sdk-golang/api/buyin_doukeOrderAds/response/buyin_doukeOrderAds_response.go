package buyin_doukeOrderAds_response

import (
	"doudian.com/open/sdk_golang/core"
)

type BuyinDoukeOrderAdsResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *Data `json:"data"`
}
type ProductTags struct {
	// 是否抖音超市(次日达)商品
	HasSupermarketTag bool `json:"has_supermarket_tag"`
	// 是否超值购商品
	HasSubsidyTag bool `json:"has_subsidy_tag"`
}
type PidInfo struct {
	// PID
	Pid string `json:"pid"`
	// 外部参数
	ExternalInfo string `json:"external_info"`
	// 分销回流类型。live： 直播间分销； ProductDetail：商品分销； Activity： 百补和秒杀推广；Shop：店铺导流
	MediaTypeName string `json:"media_type_name"`
}
type OrdersItem struct {
	// 商品分销订单-商品标签信息
	ProductTags *ProductTags `json:"product_tags"`
	// 收货时间
	ConfirmTime string `json:"confirm_time"`
	// 活动页物料ID
	MaterialId string `json:"material_id"`
	// 被分销达人百应ID（仅直播间分销订单拥有）
	AuthorBuyinId string `json:"author_buyin_id"`
	// 订单号
	OrderId string `json:"order_id"`
	// 商品id
	ProductId string `json:"product_id"`
	// 商品名称
	ProductName string `json:"product_name"`
	// 商品图片URL
	ProductImg string `json:"product_img"`
	// 作者账号昵称(抖音/火山作者)（仅直播间分销订单有）
	AuthorAccount string `json:"author_account"`
	// 商家名称
	ShopName string `json:"shop_name"`
	// 订单支付金额，单位分
	TotalPayAmount int64 `json:"total_pay_amount"`
	// 订单状态。PAY_SUCC：支付完成； REFUND：退款； SETTLE：结算；CONFIRM：确认收货
	FlowPoint string `json:"flow_point"`
	// 更新时间 [联盟侧订单更新时间]
	UpdateTime string `json:"update_time"`
	// 付款时间
	PaySuccessTime string `json:"pay_success_time"`
	// 结算时间，结算前为空字符串
	SettleTime string `json:"settle_time"`
	// 预估参与结算金额
	PayGoodsAmount int64 `json:"pay_goods_amount"`
	// 实际参与结算金额
	SettledGoodsAmount int64 `json:"settled_goods_amount"`
	// 店铺ID
	ShopId int64 `json:"shop_id"`
	// 分销订单相关参数
	PidInfo *PidInfo `json:"pid_info"`
	// 商品数目
	ItemNum int64 `json:"item_num"`
	// 退款订单退款时间
	RefundTime string `json:"refund_time"`
	// 预估推广费收入，单位分
	AdsEstimatedCommission int64 `json:"ads_estimated_commission"`
	// 实际推广费收入，单位分
	AdsRealCommission int64 `json:"ads_real_commission"`
	// 达人抖音号（仅直播间分销订单有）
	AuthorShortId string `json:"author_short_id"`
	// 推广费率
	AdsPromotionRate int64 `json:"ads_promotion_rate"`
	// 带货体裁。shop_list：橱窗；video：视频；live：直播；others：其他
	MediaType string `json:"media_type"`
	// 活动页推广活动Id，1000-超值购 1001-秒杀;若distribution_type为Mix，则为Mix活动ID
	AdsActivityId int64 `json:"ads_activity_id"`
	// 推广技术服务费
	AdsEstimatedTechServiceFee int64 `json:"ads_estimated_tech_service_fee"`
	// 直播间分销订单：新老粉
	AdsFansType string `json:"ads_fans_type"`
	// 商品参与的活动id，0: 未参加活动 1: 超值购（活动页单品推广）
	ProductActivityId string `json:"product_activity_id"`
}
type Data struct {
	// 订单列表
	Orders []OrdersItem `json:"orders"`
	// 下一页的索引
	Cursor string `json:"cursor"`
}
type BuyinDoukeOrderAdsData struct {
	// 结果
	Data *Data `json:"data"`
}

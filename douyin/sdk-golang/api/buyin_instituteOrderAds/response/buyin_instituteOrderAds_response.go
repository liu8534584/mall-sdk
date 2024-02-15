package buyin_instituteOrderAds_response

import (
	"doudian.com/open/sdk_golang/core"
)

type BuyinInstituteOrderAdsResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *BuyinInstituteOrderAdsData `json:"data"`
}
type PidInfo struct {
	// PID
	Pid string `json:"pid"`
	// 外部参数
	ExternalInfo string `json:"external_info"`
	// 分销类型，直播：Live, 商品: ProductDetail
	MediaTypeName string `json:"media_type_name"`
}
type OrdersItem struct {
	// 订单号
	OrderId string `json:"order_id"`
	// 商品id
	ProductId string `json:"product_id"`
	// 商品名称
	ProductName string `json:"product_name"`
	// 商品图片URL
	ProductImg string `json:"product_img"`
	// 作者账号昵称(抖音/火山作者)
	AuthorAccount string `json:"author_account"`
	// 作者抖店open_id
	AuthorOpenid string `json:"author_openid"`
	// 商家名称
	ShopName string `json:"shop_name"`
	// 订单支付金额，单位分
	TotalPayAmount int64 `json:"total_pay_amount"`
	// 订单状态(PAY_SUCC:支付完成 REFUND:退款 SETTLE:结算)
	FlowPoint string `json:"flow_point"`
	// App名称（抖音，火山）
	App string `json:"app"`
	// 下单用户抖店open_id
	BuyerOpenid string `json:"buyer_openid"`
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
	// 选品App client_key
	PickSourceClientKey string `json:"pick_source_client_key"`
	// 废弃字段
	AdsSplitRate int64 `json:"ads_split_rate"`
	// 直播间分销渠道预估佣金收入，单位分
	AdsEstimatedCommission int64 `json:"ads_estimated_commission"`
	// 直播间分销渠道实际佣金收入，单位分
	AdsRealCommission int64 `json:"ads_real_commission"`
	// 达人抖音号
	AuthorShortId string `json:"author_short_id"`
	// 达人/自播商家给直播间分销渠道设置的推广费率
	AdsPromotionRate int64 `json:"ads_promotion_rate"`
	// 达人百应ID
	AuthorBuyinId string `json:"author_buyin_id"`
	// 带货体裁。shop_list：橱窗；video：视频；live：直播；others：其他
	MediaType string `json:"media_type"`
	// 直播间分销订单用户类型 new：新客 old：老客
	AdsFansType string `json:"ads_fans_type"`
}
type Data struct {
	// 订单列表
	Orders []OrdersItem `json:"orders"`
	// 下一页索引
	Cursor string `json:"cursor"`
}
type BuyinInstituteOrderAdsData struct {
	// 结果
	Data *Data `json:"data"`
}

package buyin_commonShareCommandParse_response

import (
	"doudian.com/open/sdk_golang/core"
)

type BuyinCommonShareCommandParseResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *BuyinCommonShareCommandParseData `json:"data"`
}
type ActivityInfo struct {
	// 活动页自定义参数，对应活动页转链接口的同名参数
	ExtraParams string `json:"extra_params"`
	// 活动物料Id
	MaterialId int64 `json:"material_id"`
	// 1:百亿补贴2:秒杀活动3:自建活动页
	ActivityId int64 `json:"activity_id"`
	// 可推广的运营自建活动页活动id；activity_id为3时，返回该项
	MixActivityId string `json:"mix_activity_id"`
	// 可推广的运营自建活动页链接；activity_id为3时，返回该项
	ActivityUrl string `json:"activity_url"`
}
type RedpackInfo struct {
	// 抖客红包ID
	RedpackId string `json:"redpack_id"`
}
type CommonShareCommandParseInfo struct {
	// 口令类型 1:商品 2:直播间 3:直播预告 4:活动页 5:抖客红包 6:商品优惠
	CommandType int16 `json:"command_type"`
	// 商品信息，command_type为1或6返回该项
	ProductInfo *ProductInfo `json:"product_info"`
	// 直播间信息，仅command_type为2返回该项
	LiveInfo *LiveInfo `json:"live_info"`
	// 直播预告信息，仅command_type为3返回该项
	LiveAppointInfo *LiveAppointInfo `json:"live_appoint_info"`
	// 活动页信息，仅command_type为4返回该项
	ActivityInfo *ActivityInfo `json:"activity_info"`
	// 抖客红包信息，仅command_type为5返回该项
	RedpackInfo *RedpackInfo `json:"redpack_info"`
}
type BuyinCommonShareCommandParseData struct {
	// 口令信息
	CommonShareCommandParseInfo *CommonShareCommandParseInfo `json:"common_share_command_parse_info"`
}
type ProductInfo struct {
	// 团长参数
	InsActivityParam string `json:"ins_activity_param"`
	// 商品ID
	ProductId int64 `json:"product_id"`
}
type LiveInfo struct {
	// 主播百应ID
	AuthorBuyinId string `json:"author_buyin_id"`
	// 直播间绑定的商品ID，product_id专属于该直播间，可能已经失效
	ProductId int64 `json:"product_id"`
}
type LiveAppointInfo struct {
	// 直播预告对应的达人百应ID
	AuthorBuyinId string `json:"author_buyin_id"`
	// 直播预告ID
	LiveAppointmentId string `json:"live_appointment_id"`
}

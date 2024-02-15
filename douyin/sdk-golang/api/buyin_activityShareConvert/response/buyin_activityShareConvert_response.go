package buyin_activityShareConvert_response

import (
	"doudian.com/open/sdk_golang/core"
)

type BuyinActivityShareConvertResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	BuyinActivityShareConvertData
}
type Qrcode struct {
	// 图片URL
	Url string `json:"url"`
	// 宽带
	Width int32 `json:"width"`
	// 高度
	Height int32 `json:"height"`
}
type Data struct {
	// 抖口令
	ShareCommand string `json:"share_command"`
	// Deeplink，允许外部APP直接唤起抖音直播间
	DeepLink string `json:"deep_link"`
	// 二维码
	Qrcode Qrcode `json:"qrcode"`
}
type BuyinActivityShareConvertData struct {
	// 转链数据
	Data Data `json:"data"`
}

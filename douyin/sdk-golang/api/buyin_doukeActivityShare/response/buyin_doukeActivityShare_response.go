package buyin_doukeActivityShare_response

import (
	"doudian.com/open/sdk_golang/core"
)

type BuyinDoukeActivityShareResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *Data `json:"data"`
}
type Data struct {
	// 抖口令
	ShareCommand string `json:"share_command"`
	// Deeplink，允许外部APP直接唤起抖音
	DeepLink string `json:"deep_link"`
	// 二维码
	Qrcode *Qrcode `json:"qrcode"`
	// Zlink，允许外部APP直接唤起抖音直播间，未安装抖音应用用户，唤起抖音下载页，有效期60天
	ZLink string `json:"z_link"`
}
type BuyinDoukeActivityShareData struct {
	// 转链数据
	Data *Data `json:"data"`
}
type Qrcode struct {
	// 图片URL
	Url string `json:"url"`
	// 图片宽度
	Width int32 `json:"width"`
	// 图片高度
	Height int32 `json:"height"`
}

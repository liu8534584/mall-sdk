package buyin_kolProductShare_response

import (
	"doudian.com/open/sdk_golang/core"
)

type BuyinKolProductShareResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	BuyinKolProductShareData
}
type QrCode struct {
	// 图片地址
	Url string `json:"url"`
	// 宽度
	Width int32 `json:"width"`
	// 高度
	Height int32 `json:"height"`
}
type Data struct {
	// 抖口令
	DyPassword string `json:"dy_password"`
	// 二维码
	QrCode QrCode `json:"qr_code"`
	// deeplink
	DyDeeplink string `json:"dy_deeplink"`
	// 是否限制抖客推送
	IsLimitDouke bool   `json:"is_limit_douke"`
	ShareLink    string `json:"share_link"`
}
type BuyinKolProductShareData struct {
	// 返回数据
	Data Data `json:"data"`
}

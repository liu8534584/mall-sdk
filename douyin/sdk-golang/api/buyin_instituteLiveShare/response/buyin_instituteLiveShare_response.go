package buyin_instituteLiveShare_response

import (
	"doudian.com/open/sdk_golang/core"
)

type BuyinInstituteLiveShareResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *BuyinInstituteLiveShareData `json:"data"`
}
type QrCode struct {
	// URL
	Url string `json:"url"`
	// 宽度
	Width int32 `json:"width"`
	// 高度
	Height int32 `json:"height"`
}
type BuyinInstituteLiveShareData struct {
	// 抖口令
	DyPassword string `json:"dy_password"`
	// 二维码
	QrCode *QrCode `json:"qr_code"`
	// Deeplink，允许外部APP直接唤起抖音直播间
	DyDeeplink string `json:"dy_deeplink"`
	// Zlink，允许外部APP直接唤起抖音直播间，未安装抖音应用用户，唤起抖音下载页，有效期60天
	DyZlink string `json:"dy_zlink"`
}

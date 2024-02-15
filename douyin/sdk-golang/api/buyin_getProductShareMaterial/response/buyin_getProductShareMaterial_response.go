package buyin_getProductShareMaterial_response

import (
	"doudian.com/open/sdk_golang/core"
)

type BuyinGetProductShareMaterialResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *BuyinGetProductShareMaterialData `json:"data"`
}
type BuyinGetProductShareMaterialData struct {
	// 分销商品列表文件
	FileUrl string `json:"file_url"`
}

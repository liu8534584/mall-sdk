package buyin_kolPidDel_response

import (
	"doudian.com/open/sdk_golang/core"
)

type BuyinKolPidDelResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *BuyinKolPidDelData `json:"data"`
}
type BuyinKolPidDelData struct {
}

package buyin_kolPidEdit_response

import (
	"doudian.com/open/sdk_golang/core"
)

type BuyinKolPidEditResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *BuyinKolPidEditData `json:"data"`
}
type BuyinKolPidEditData struct {
}

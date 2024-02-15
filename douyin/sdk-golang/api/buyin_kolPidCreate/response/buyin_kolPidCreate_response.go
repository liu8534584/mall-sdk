package buyin_kolPidCreate_response

import (
	"doudian.com/open/sdk_golang/core"
)

type BuyinKolPidCreateResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *BuyinKolPidCreateData `json:"data"`
}
type BuyinKolPidCreateData struct {
	// 达人PID
	Pid string `json:"pid"`
}

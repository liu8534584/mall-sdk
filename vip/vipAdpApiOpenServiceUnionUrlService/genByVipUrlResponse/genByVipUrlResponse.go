package genByVipUrlResponse

import (
	"github.com/liu8534584/mall-sdk/vip/vipAdpApiOpenServiceUnionUrlService/cpsLinkUrlInfo"
)

type GenByVipUrlResponse struct {
	ReturnCode string        `json:"returnCode"`
	Result     UrlInfoResult `json:"result"`
}

type UrlInfoResult struct {
	UrlInfoLists []cpsLinkUrlInfo.UrlInfoList `json:"urlInfoList"`
}

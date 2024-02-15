package response

import "github.com/liu8534584/mall-sdk/goodsCenter/response/base"

type ShortItemIds2LongResp struct {
	Code     base.Code               `json:"code"`
	Message  string                  `json:"msg"`
	Demotion base.Demotion           `json:"demotion"`
	Data     map[string]TbFullItemId `json:"data"`
}
type TbFullItemId struct {
	FullItemId   string `json:"fullItemId"`
	UniqueItemId string `json:"uniqueItemId"`
}

package response

import "github.com/liu8534584/mall-sdk/goodsCenter/response/base"

type ItemDetailResp struct {
	Code     base.Code     `json:"code"`
	Message  string        `json:"msg"`
	Demotion base.Demotion `json:"demotion"`
	Data     ItemDetail    `json:"data"`
}

type BathItemDetailResp struct {
	Code     base.Code     `json:"code"`
	Message  string        `json:"msg"`
	Demotion base.Demotion `json:"demotion"`
	Data     []ItemDetail  `json:"data"`
}

type BathItemDetailV2Resp struct {
	Code     base.Code             `json:"code"`
	Message  string                `json:"msg"`
	Demotion base.Demotion         `json:"demotion"`
	Data     map[string]ItemDetail `json:"data"`
}

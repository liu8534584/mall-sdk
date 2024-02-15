package response

import (
	"github.com/liu8534584/mall-sdk/goodsCenter/response/base"
)

type ResolveTitleResponse struct {
	Code     base.Code     `json:"code"`
	Message  string        `json:"msg"`
	Demotion base.Demotion `json:"demotion"`
	Data     ItemDetail    `json:"data"`
}

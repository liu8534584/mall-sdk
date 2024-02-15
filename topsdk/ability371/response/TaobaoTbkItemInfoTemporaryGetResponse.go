package response

import (
	"github.com/liu8534584/mall-sdk/topsdk/ability371/domain"
)

type TaobaoTbkItemInfoTemporaryGetResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   淘宝客商品
	*/
	Results []domain.TaobaoTbkItemInfoTemporaryGetNTbkItem `json:"results,omitempty" `
}

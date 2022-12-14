package response

import (
	"github.com/liu8534584/mall-sdk/topsdk/ability411/domain"
)

type TaobaoTbkScMembergroupOptionalResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   θΏεη»ζ
	*/
	Data domain.TaobaoTbkScMembergroupOptionalMapData `json:"data,omitempty" `
}

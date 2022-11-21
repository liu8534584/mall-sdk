package response

import (
	"github.com/liu8534584/mall-sdk/topsdk/ability427/domain"
)

type TaobaoTbkScAdzoneCreateResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   MapData
	*/
	Data domain.TaobaoTbkScAdzoneCreateMapData `json:"data,omitempty" `
}

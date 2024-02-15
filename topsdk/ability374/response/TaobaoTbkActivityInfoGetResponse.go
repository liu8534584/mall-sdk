package response

import (
	"github.com/liu8534584/mall-sdk/topsdk/ability374/domain"
)

type TaobaoTbkActivityInfoGetResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   返回结果对象
	*/
	Data domain.TaobaoTbkActivityInfoGetData `json:"data,omitempty" `
}

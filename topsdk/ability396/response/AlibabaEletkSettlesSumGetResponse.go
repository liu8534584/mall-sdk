package response

import (
	"github.com/liu8534584/mall-sdk/topsdk/ability396/domain"
)

type AlibabaEletkSettlesSumGetResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   接口返回model
	*/
	Result domain.AlibabaEletkSettlesSumGetResultDTO `json:"result,omitempty" `
}

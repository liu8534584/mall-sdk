package response

import (
	"github.com/liu8534584/mall-sdk/topsdk/defaultability/domain"
)

type AlibabaAlscUnionElemePromotionStorepromotionGetResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   门店推广
	*/
	Data domain.AlibabaAlscUnionElemePromotionStorepromotionGetStorePromotionDto `json:"data,omitempty" `
	/*
	   返回码
	*/
	ResultCode int64 `json:"result_code,omitempty" `
	/*
	   返回消息
	*/
	Message string `json:"message,omitempty" `
}

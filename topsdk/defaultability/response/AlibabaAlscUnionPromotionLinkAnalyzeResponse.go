package response

import "github.com/liu8534584/mall-sdk/topsdk/defaultability/domain"

type AlibabaAlscUnionPromotionLinkAnalyzeResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   推广对象
	*/
	Data domain.AlibabaAlscUnionPromotionLinkAnalyzePromotionObject `json:"data,omitempty" `
	/*
	   请求结果码
	*/
	ResultCode int64 `json:"result_code,omitempty" `
	/*
	   返回描述
	*/
	Message string `json:"message,omitempty" `
}

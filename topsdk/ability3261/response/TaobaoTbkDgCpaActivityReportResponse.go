package response

import (
	"github.com/liu8534584/mall-sdk/topsdk/ability3261/domain"
)

type TaobaoTbkDgCpaActivityReportResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   返回模型
	*/
	Result domain.TaobaoTbkDgCpaActivityReportRpcResult `json:"result,omitempty" `
}

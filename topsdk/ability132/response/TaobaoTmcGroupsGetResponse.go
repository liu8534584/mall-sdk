package response

import (
	"github.com/liu8534584/mall-sdk/topsdk/ability132/domain"
)

type TaobaoTmcGroupsGetResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   dasdasd
	*/
	Groups []domain.TaobaoTmcGroupsGetTmcGroup `json:"groups,omitempty" `
	/*
	   分组总数
	*/
	TotalResults int64 `json:"total_results,omitempty" `
}

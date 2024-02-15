package response

import (
	"github.com/liu8534584/mall-sdk/topsdk/defaultability/domain"
)

type TaobaoTbkScMaterialOptionalUpgradeResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   搜索到符合条件的结果总数
	*/
	TotalResults int64 `json:"total_results,omitempty" `
	/*
	   resultList
	*/
	ResultList []domain.TaobaoTbkScMaterialOptionalUpgradeMapData `json:"result_list,omitempty" `
	/*
	   uvid结果信息，传入但未使用uvid时会返回原因
	*/
	UvidMsg string `json:"uvid_msg,omitempty" `
}

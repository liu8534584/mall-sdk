package queryResponse

import "github.com/liu8534584/mall-sdk/vip/vipAdpApiOpenServiceUnionGoodsService/getByGoodsIdsResponse"

type UnionGoodsServiceQueryResponse struct {
	ReturnCode string                               `json:"returnCode"`
	Result     UnionGoodsServiceQueryResponseResult `json:"result"`
}

type UnionGoodsServiceQueryResponseResult struct {
	Total         int                                  `json:"total"`
	GoodsInfoList []getByGoodsIdsResponse.VipGoodsInfo `json:"goodsInfoList"`
	PageSize      int                                  `json:"pageSize"`
	SortFields    []struct {
		FieldName string `json:"fieldName"`
		FieldDesc string `json:"fieldDesc"`
	} `json:"sortFields"`
	Page int `json:"page"`
}

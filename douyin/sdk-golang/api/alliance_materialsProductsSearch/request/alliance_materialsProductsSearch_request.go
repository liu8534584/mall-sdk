package alliance_materialsProductsSearch_request

import (
	"doudian.com/open/sdk_golang/api/alliance_materialsProductsSearch/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type AllianceMaterialsProductsSearchRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *AllianceMaterialsProductsSearchParam
}

func (c *AllianceMaterialsProductsSearchRequest) GetUrlPath() string {
	return "/alliance/materialsProductsSearch"
}

func New() *AllianceMaterialsProductsSearchRequest {
	request := &AllianceMaterialsProductsSearchRequest{
		Param: &AllianceMaterialsProductsSearchParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}

func (c *AllianceMaterialsProductsSearchRequest) Execute(accessToken *doudian_sdk.AccessToken) (*alliance_materialsProductsSearch_response.AllianceMaterialsProductsSearchResponse, error) {
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &alliance_materialsProductsSearch_response.AllianceMaterialsProductsSearchResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}

func (c *AllianceMaterialsProductsSearchRequest) GetParamObject() interface{} {
	return c.Param
}

func (c *AllianceMaterialsProductsSearchRequest) GetParams() *AllianceMaterialsProductsSearchParam {
	return c.Param
}

type AllianceMaterialsProductsSearchParam struct {
	// 商品标题，返回标题中包含某个关键词的商品
	Title string `json:"title,omitempty"`
	// 筛选商品一级类目
	FirstCids []int64 `json:"first_cids,omitempty"`
	// 筛选商品二级类目
	SecondCids []int64 `json:"second_cids,omitempty"`
	// 筛选商品三级类目
	ThirdCids []int64 `json:"third_cids,omitempty"`
	// 筛选价格区间-最小值（单位为分）
	PriceMin int32 `json:"price_min"`
	// 筛选价格区间-最大值（单位为分）
	PriceMax int32 `json:"price_max"`
	// 筛选历史销量区间-最小值
	SellNumMin int32 `json:"sell_num_min"`
	// 筛选历史销量区间-最大值
	SellNumMax int32 `json:"sell_num_max"`
	// 召回结果排序条件，0默认排序1历史销量排序2价格排序3佣金金额排序4佣金比例排序；
	SearchType int32 `json:"search_type"`
	// 排序顺序（0升序1降序）
	SortType int32 `json:"sort_type"`
	// 筛选普通佣金区间-最小值（单位为分）
	CosFeeMin int32 `json:"cos_fee_min"`
	// 筛选普通佣金区间-最大值（单位为分）
	CosFeeMax int32 `json:"cos_fee_max"`
	// 筛选普通佣金率区间-最小值 （乘100，例如1.1%为110）
	CosRatioMin int32 `json:"cos_ratio_min"`
	// 筛选普通佣金率区间-最大值（乘100，例如1.1%为110）
	CosRatioMax int32 `json:"cos_ratio_max"`
	// 分页（从1开始）
	Page int64 `json:"page"`
	// 每页的数量（小于等于20）
	PageSize int64 `json:"page_size"`
	// 获取商品分销状态。1: 仅返回可分销商品；0:返回全量商品
	ShareStatus int32 `json:"share_status"`
}

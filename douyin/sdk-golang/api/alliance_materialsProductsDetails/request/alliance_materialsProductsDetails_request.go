package alliance_materialsProductsDetails_request

import (
	"doudian.com/open/sdk_golang/api/alliance_materialsProductsDetails/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type AllianceMaterialsProductsDetailsRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *AllianceMaterialsProductsDetailsParam
}

func (c *AllianceMaterialsProductsDetailsRequest) GetUrlPath() string {
	return "/alliance/materialsProductsDetails"
}

func New() *AllianceMaterialsProductsDetailsRequest {
	request := &AllianceMaterialsProductsDetailsRequest{
		Param: &AllianceMaterialsProductsDetailsParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}

func (c *AllianceMaterialsProductsDetailsRequest) Execute(accessToken *doudian_sdk.AccessToken) (*alliance_materialsProductsDetails_response.AllianceMaterialsProductsDetailsResponse, error) {
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &alliance_materialsProductsDetails_response.AllianceMaterialsProductsDetailsResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}

func (c *AllianceMaterialsProductsDetailsRequest) GetParamObject() interface{} {
	return c.Param
}

func (c *AllianceMaterialsProductsDetailsRequest) GetParams() *AllianceMaterialsProductsDetailsParam {
	return c.Param
}

type AllianceMaterialsProductsDetailsParam struct {
	// 商品ID，最多20个。注意：如商品推广在抖店后台下架，接口会返回为空。
	ProductIds []int64 `json:"product_ids"`
	// 是否返回分销状态；true-返回；false-不返回；默认不返回。 根据响应参数sharable字段判断。
	WithShareStatus bool `json:"with_share_status"`
	// 是否返回品牌信息；true-返回；false-不返回；默认不返回
	WithBrandInfo bool `json:"with_brand_info"`
}

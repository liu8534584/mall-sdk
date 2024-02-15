package buyin_materialsProductStatus_request

import (
	"doudian.com/open/sdk_golang/api/buyin_materialsProductStatus/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type BuyinMaterialsProductStatusRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *BuyinMaterialsProductStatusParam
}

func (c *BuyinMaterialsProductStatusRequest) GetUrlPath() string {
	return "/buyin/materialsProductStatus"
}

func New() *BuyinMaterialsProductStatusRequest {
	request := &BuyinMaterialsProductStatusRequest{
		Param: &BuyinMaterialsProductStatusParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}

func (c *BuyinMaterialsProductStatusRequest) Execute(accessToken *doudian_sdk.AccessToken) (*buyin_materialsProductStatus_response.BuyinMaterialsProductStatusResponse, error) {
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &buyin_materialsProductStatus_response.BuyinMaterialsProductStatusResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}

func (c *BuyinMaterialsProductStatusRequest) GetParamObject() interface{} {
	return c.Param
}

func (c *BuyinMaterialsProductStatusRequest) GetParams() *BuyinMaterialsProductStatusParam {
	return c.Param
}

type BuyinMaterialsProductStatusParam struct {
	// 商品 url 列表
	Products []string `json:"products"`
}

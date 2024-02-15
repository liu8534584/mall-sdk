package buyin_getProductShareMaterial_request

import (
	"doudian.com/open/sdk_golang/api/buyin_getProductShareMaterial/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type BuyinGetProductShareMaterialRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *BuyinGetProductShareMaterialParam
}

func (c *BuyinGetProductShareMaterialRequest) GetUrlPath() string {
	return "/buyin/getProductShareMaterial"
}

func New() *BuyinGetProductShareMaterialRequest {
	request := &BuyinGetProductShareMaterialRequest{
		Param: &BuyinGetProductShareMaterialParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}

func (c *BuyinGetProductShareMaterialRequest) Execute(accessToken *doudian_sdk.AccessToken) (*buyin_getProductShareMaterial_response.BuyinGetProductShareMaterialResponse, error) {
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &buyin_getProductShareMaterial_response.BuyinGetProductShareMaterialResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}

func (c *BuyinGetProductShareMaterialRequest) GetParamObject() interface{} {
	return c.Param
}

func (c *BuyinGetProductShareMaterialRequest) GetParams() *BuyinGetProductShareMaterialParam {
	return c.Param
}

type BuyinGetProductShareMaterialParam struct {
}

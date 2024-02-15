package buyin_productSkus_request

import (
	"doudian.com/open/sdk_golang/api/buyin_productSkus/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type BuyinProductSkusRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *BuyinProductSkusParam
}

func (c *BuyinProductSkusRequest) GetUrlPath() string {
	return "/buyin/productSkus"
}

func New() *BuyinProductSkusRequest {
	request := &BuyinProductSkusRequest{
		Param: &BuyinProductSkusParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}

func (c *BuyinProductSkusRequest) Execute(accessToken *doudian_sdk.AccessToken) (*buyin_productSkus_response.BuyinProductSkusResponse, error) {
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &buyin_productSkus_response.BuyinProductSkusResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}

func (c *BuyinProductSkusRequest) GetParamObject() interface{} {
	return c.Param
}

func (c *BuyinProductSkusRequest) GetParams() *BuyinProductSkusParam {
	return c.Param
}

type BuyinProductSkusParam struct {
	// 商品 id
	ProductId int64 `json:"product_id"`
}

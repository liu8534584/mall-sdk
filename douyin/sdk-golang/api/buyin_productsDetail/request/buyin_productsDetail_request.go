package buyin_productsDetail_request

import (
	"doudian.com/open/sdk_golang/api/buyin_productsDetail/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type BuyinProductsDetailRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *BuyinProductsDetailParam
}

func (c *BuyinProductsDetailRequest) GetUrlPath() string {
	return "/buyin/productsDetail"
}

func New() *BuyinProductsDetailRequest {
	request := &BuyinProductsDetailRequest{
		Param: &BuyinProductsDetailParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}

func (c *BuyinProductsDetailRequest) Execute(accessToken *doudian_sdk.AccessToken) (*buyin_productsDetail_response.BuyinProductsDetailResponse, error) {
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &buyin_productsDetail_response.BuyinProductsDetailResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}

func (c *BuyinProductsDetailRequest) GetParamObject() interface{} {
	return c.Param
}

func (c *BuyinProductsDetailRequest) GetParams() *BuyinProductsDetailParam {
	return c.Param
}

type BuyinProductsDetailParam struct {
	// 商品 id 列表
	ProductIds *[]int64 `json:"product_ids"`
	// 字段（需要返回的字段，多个用英文逗号隔开。不传只返回商品基础信息，强烈建议按需获取数据）
	Fields string `json:"fields"`
}

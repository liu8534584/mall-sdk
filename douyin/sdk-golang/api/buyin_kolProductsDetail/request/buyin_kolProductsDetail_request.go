package buyin_kolProductsDetail_request

import (
	"doudian.com/open/sdk_golang/api/buyin_kolProductsDetail/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type BuyinKolProductsDetailRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *BuyinKolProductsDetailParam
}

func (c *BuyinKolProductsDetailRequest) GetUrlPath() string {
	return "/buyin/kolProductsDetail"
}

func New() *BuyinKolProductsDetailRequest {
	request := &BuyinKolProductsDetailRequest{
		Param: &BuyinKolProductsDetailParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}

func (c *BuyinKolProductsDetailRequest) Execute(accessToken *doudian_sdk.AccessToken) (*buyin_kolProductsDetail_response.BuyinKolProductsDetailResponse, error) {
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &buyin_kolProductsDetail_response.BuyinKolProductsDetailResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}

func (c *BuyinKolProductsDetailRequest) GetParamObject() interface{} {
	return c.Param
}

func (c *BuyinKolProductsDetailRequest) GetParams() *BuyinKolProductsDetailParam {
	return c.Param
}

type BuyinKolProductsDetailParam struct {
	// 商品 id 列表
	ProductIds *[]int64 `json:"product_ids"`
	// 需要获取的字段（取products字段下的一级字段，多个用英文逗号隔开，如果不传默认只返回商品基础信息。强烈建议数据按需获取）
	Fields string `json:"fields"`
}

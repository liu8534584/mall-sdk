package buyin_shareCommandParse_request

import (
	"doudian.com/open/sdk_golang/api/buyin_shareCommandParse/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type BuyinShareCommandParseRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *BuyinShareCommandParseParam
}

func (c *BuyinShareCommandParseRequest) GetUrlPath() string {
	return "/buyin/shareCommandParse"
}

func New() *BuyinShareCommandParseRequest {
	request := &BuyinShareCommandParseRequest{
		Param: &BuyinShareCommandParseParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}

func (c *BuyinShareCommandParseRequest) Execute(accessToken *doudian_sdk.AccessToken) (*buyin_shareCommandParse_response.BuyinShareCommandParseResponse, error) {
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &buyin_shareCommandParse_response.BuyinShareCommandParseResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}

func (c *BuyinShareCommandParseRequest) GetParamObject() interface{} {
	return c.Param
}

func (c *BuyinShareCommandParseRequest) GetParams() *BuyinShareCommandParseParam {
	return c.Param
}

type BuyinShareCommandParseParam struct {
	// 抖音口令（仅商品口令）
	Command string `json:"command"`
}

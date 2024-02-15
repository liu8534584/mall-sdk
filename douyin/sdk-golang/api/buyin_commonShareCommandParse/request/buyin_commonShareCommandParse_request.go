package buyin_commonShareCommandParse_request

import (
	"doudian.com/open/sdk_golang/api/buyin_commonShareCommandParse/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type BuyinCommonShareCommandParseRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *BuyinCommonShareCommandParseParam
}

func (c *BuyinCommonShareCommandParseRequest) GetUrlPath() string {
	return "/buyin/commonShareCommandParse"
}

func New() *BuyinCommonShareCommandParseRequest {
	request := &BuyinCommonShareCommandParseRequest{
		Param: &BuyinCommonShareCommandParseParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}

func (c *BuyinCommonShareCommandParseRequest) Execute(accessToken *doudian_sdk.AccessToken) (*buyin_commonShareCommandParse_response.BuyinCommonShareCommandParseResponse, error) {
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &buyin_commonShareCommandParse_response.BuyinCommonShareCommandParseResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}

func (c *BuyinCommonShareCommandParseRequest) GetParamObject() interface{} {
	return c.Param
}

func (c *BuyinCommonShareCommandParseRequest) GetParams() *BuyinCommonShareCommandParseParam {
	return c.Param
}

type BuyinCommonShareCommandParseParam struct {
	// 抖口令或短链
	Command *string `json:"command,omitempty"`
}

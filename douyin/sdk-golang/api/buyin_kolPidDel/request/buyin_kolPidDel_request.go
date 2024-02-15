package buyin_kolPidDel_request

import (
	"doudian.com/open/sdk_golang/api/buyin_kolPidDel/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type BuyinKolPidDelRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *BuyinKolPidDelParam
}

func (c *BuyinKolPidDelRequest) GetUrlPath() string {
	return "/buyin/kolPidDel"
}

func New() *BuyinKolPidDelRequest {
	request := &BuyinKolPidDelRequest{
		Param: &BuyinKolPidDelParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}

func (c *BuyinKolPidDelRequest) Execute(accessToken *doudian_sdk.AccessToken) (*buyin_kolPidDel_response.BuyinKolPidDelResponse, error) {
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &buyin_kolPidDel_response.BuyinKolPidDelResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}

func (c *BuyinKolPidDelRequest) GetParamObject() interface{} {
	return c.Param
}

func (c *BuyinKolPidDelRequest) GetParams() *BuyinKolPidDelParam {
	return c.Param
}

type BuyinKolPidDelParam struct {
	// PID列表，多个用,分隔，最多200个
	Pids string `json:"pids"`
}

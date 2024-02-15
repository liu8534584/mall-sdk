package buyin_kolPidEdit_request

import (
	"doudian.com/open/sdk_golang/api/buyin_kolPidEdit/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type BuyinKolPidEditRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *BuyinKolPidEditParam
}

func (c *BuyinKolPidEditRequest) GetUrlPath() string {
	return "/buyin/kolPidEdit"
}

func New() *BuyinKolPidEditRequest {
	request := &BuyinKolPidEditRequest{
		Param: &BuyinKolPidEditParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}

func (c *BuyinKolPidEditRequest) Execute(accessToken *doudian_sdk.AccessToken) (*buyin_kolPidEdit_response.BuyinKolPidEditResponse, error) {
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &buyin_kolPidEdit_response.BuyinKolPidEditResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}

func (c *BuyinKolPidEditRequest) GetParamObject() interface{} {
	return c.Param
}

func (c *BuyinKolPidEditRequest) GetParams() *BuyinKolPidEditParam {
	return c.Param
}

type BuyinKolPidEditParam struct {
	// PID
	Pid string `json:"pid"`
	// 推广位名称，不能和已有的重复(最长10个字符)
	SiteName string `json:"site_name"`
}

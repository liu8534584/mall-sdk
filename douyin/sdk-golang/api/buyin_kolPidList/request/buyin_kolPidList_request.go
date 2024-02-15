package buyin_kolPidList_request

import (
	"doudian.com/open/sdk_golang/api/buyin_kolPidList/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type BuyinKolPidListRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *BuyinKolPidListParam
}

func (c *BuyinKolPidListRequest) GetUrlPath() string {
	return "/buyin/kolPidList"
}

func New() *BuyinKolPidListRequest {
	request := &BuyinKolPidListRequest{
		Param: &BuyinKolPidListParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}

func (c *BuyinKolPidListRequest) Execute(accessToken *doudian_sdk.AccessToken) (*buyin_kolPidList_response.BuyinKolPidListResponse, error) {
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &buyin_kolPidList_response.BuyinKolPidListResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}

func (c *BuyinKolPidListRequest) GetParamObject() interface{} {
	return c.Param
}

func (c *BuyinKolPidListRequest) GetParams() *BuyinKolPidListParam {
	return c.Param
}

type BuyinKolPidListParam struct {
	// 页数，默认从1开始
	Page int64 `json:"page"`
	// 每页大小。默认100，最大2000
	PageSize int64 `json:"page_size"`
	// 推广媒体类型，0:其他, 1:微信, 2:QQ, 3:微博, 4:抖音, 5:头条
	MediaType int32 `json:"media_type"`
	// PID
	Pid string `json:"pid"`
}

package buyin_kolPidCreate_request

import (
	"doudian.com/open/sdk_golang/api/buyin_kolPidCreate/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type BuyinKolPidCreateRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *BuyinKolPidCreateParam
}

func (c *BuyinKolPidCreateRequest) GetUrlPath() string {
	return "/buyin/kolPidCreate"
}

func New() *BuyinKolPidCreateRequest {
	request := &BuyinKolPidCreateRequest{
		Param: &BuyinKolPidCreateParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}

func (c *BuyinKolPidCreateRequest) Execute(accessToken *doudian_sdk.AccessToken) (*buyin_kolPidCreate_response.BuyinKolPidCreateResponse, error) {
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &buyin_kolPidCreate_response.BuyinKolPidCreateResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}

func (c *BuyinKolPidCreateRequest) GetParamObject() interface{} {
	return c.Param
}

func (c *BuyinKolPidCreateRequest) GetParams() *BuyinKolPidCreateParam {
	return c.Param
}

type BuyinKolPidCreateParam struct {
	// 推广媒体类型，0:其他, 1:微信, 2:QQ, 3:微博, 4:抖音, 5:头条
	MediaType int32 `json:"media_type"`
	// 渠道名称，没有media_id时自动创建渠道；存在media_id时无效
	MediaName string `json:"media_name"`
	// 渠道ID，和media_name二选一
	MediaId int64 `json:"media_id"`
	// 推广位名称(最长10个字符)
	SiteName string `json:"site_name"`
}

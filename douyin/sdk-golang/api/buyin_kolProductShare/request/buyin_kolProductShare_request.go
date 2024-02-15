package buyin_kolProductShare_request

import (
	"doudian.com/open/sdk_golang/api/buyin_kolProductShare/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type BuyinKolProductShareRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *BuyinKolProductShareParam
}

func (c *BuyinKolProductShareRequest) GetUrlPath() string {
	return "/buyin/kolProductShare"
}

func New() *BuyinKolProductShareRequest {
	request := &BuyinKolProductShareRequest{
		Param: &BuyinKolProductShareParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}

func (c *BuyinKolProductShareRequest) Execute(accessToken *doudian_sdk.AccessToken) (*buyin_kolProductShare_response.BuyinKolProductShareResponse, error) {
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &buyin_kolProductShare_response.BuyinKolProductShareResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}

func (c *BuyinKolProductShareRequest) GetParamObject() interface{} {
	return c.Param
}

func (c *BuyinKolProductShareRequest) GetParams() *BuyinKolProductShareParam {
	return c.Param
}

type BuyinKolProductShareParam struct {
	// 商品URL/口令/短链接
	ProductUrl string `json:"product_url"`
	// 达人PID
	Pid string `json:"pid"`
	// 自定义字段，只允许 数字、字母和_，限制长度为20
	ExternalInfo string `json:"external_info"`
	// 是否需要二维码，需要会导致响应耗时增加
	NeedQrCode    bool `json:"need_qr_code"`
	NeedShareLink bool `json:"need_share_link"`
}

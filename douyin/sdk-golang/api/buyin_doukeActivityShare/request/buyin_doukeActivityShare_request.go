package buyin_doukeActivityShare_request

import (
	"doudian.com/open/sdk_golang/api/buyin_doukeActivityShare/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type BuyinDoukeActivityShareRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *BuyinDoukeActivityShareParam
}

func (c *BuyinDoukeActivityShareRequest) GetUrlPath() string {
	return "/buyin/doukeActivityShare"
}

func New() *BuyinDoukeActivityShareRequest {
	request := &BuyinDoukeActivityShareRequest{
		Param: &BuyinDoukeActivityShareParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}

func (c *BuyinDoukeActivityShareRequest) Execute(accessToken *doudian_sdk.AccessToken) (*buyin_doukeActivityShare_response.BuyinDoukeActivityShareResponse, error) {
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &buyin_doukeActivityShare_response.BuyinDoukeActivityShareResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}

func (c *BuyinDoukeActivityShareRequest) GetParamObject() interface{} {
	return c.Param
}

func (c *BuyinDoukeActivityShareRequest) GetParams() *BuyinDoukeActivityShareParam {
	return c.Param
}

type BuyinDoukeActivityShareParam struct {
	// 活动页物料ID
	MaterialId *int64 `json:"material_id,omitempty"`
	// 抖客PID
	Pid *string `json:"pid,omitempty"`
	// 是否需要二维码（生成二维码会增加接口耗时）
	NeedQrCode *bool `json:"need_qr_code,omitempty"`
	// 自定义字段，只允许 数字、字母和_，限制长度为40
	ExternalInfo string `json:"external_info,omitempty"`
	// 回流端标识 0：抖音 1：抖音极速版
	Platform int32 `json:"platform,omitempty"`
	// 活动页转链自定义参数，json格式，键值都是字符串
	ExtraParams string `json:"extra_params,omitempty"`
	// 活动页面置顶商品ID（仅支持秒杀、超值购、超市频道）
	ProductId int64 `json:"product_id,omitempty"`
}

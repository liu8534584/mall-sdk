package buyin_activityShareConvert_request

import (
	"doudian.com/open/sdk_golang/api/buyin_activityShareConvert/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type BuyinActivityShareConvertRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *BuyinActivityShareConvertParam 
}
func (c *BuyinActivityShareConvertRequest) GetUrlPath() string{
	return "/buyin/activityShareConvert"
}


func New() *BuyinActivityShareConvertRequest{
	request := &BuyinActivityShareConvertRequest{
		Param: &BuyinActivityShareConvertParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *BuyinActivityShareConvertRequest) Execute(accessToken *doudian_sdk.AccessToken) (*buyin_activityShareConvert_response.BuyinActivityShareConvertResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &buyin_activityShareConvert_response.BuyinActivityShareConvertResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *BuyinActivityShareConvertRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *BuyinActivityShareConvertRequest) GetParams() *BuyinActivityShareConvertParam{
	return c.Param
}


type BuyinActivityShareConvertParam struct {
	// 1:百亿补贴 2:秒杀活动 3:自建活动页
	ActivityId int64 `json:"activity_id"`
	// 必须是达人的PID
	Pid string `json:"pid"`
	// 是否需要二维码（生成二维码会增加接口耗时）
	NeedQrCode bool `json:"need_qr_code"`
	// 自定义字段，只允许 数字、字母和_，限制长度为20
	ExternalInfo string `json:"external_info"`
	// 可推广的运营自建活动页活动id；activity_id为3时，该项必填
	MixActivityId string `json:"mix_activity_id"`
	// 可推广的运营自建活动页链接，支持 h5 和 schema 两种形式；可不填
	ActivityUrl string `json:"activity_url"`
}

package buyin_instituteLiveShare_request

import (
	"doudian.com/open/sdk_golang/api/buyin_instituteLiveShare/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type BuyinInstituteLiveShareRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *BuyinInstituteLiveShareParam
}

func (c *BuyinInstituteLiveShareRequest) GetUrlPath() string {
	return "/buyin/instituteLiveShare"
}

func New() *BuyinInstituteLiveShareRequest {
	request := &BuyinInstituteLiveShareRequest{
		Param: &BuyinInstituteLiveShareParam{},
	}
	request.SetConfig(doudian_sdk.GlobalLiveConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}

func (c *BuyinInstituteLiveShareRequest) Execute(accessToken *doudian_sdk.AccessToken) (*buyin_instituteLiveShare_response.BuyinInstituteLiveShareResponse, error) {
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &buyin_instituteLiveShare_response.BuyinInstituteLiveShareResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil
}

func (c *BuyinInstituteLiveShareRequest) GetParamObject() interface{} {
	return c.Param
}

func (c *BuyinInstituteLiveShareRequest) GetParams() *BuyinInstituteLiveShareParam {
	return c.Param
}

type PidInfo struct {
	// 抖客PID
	Pid string `json:"pid"`
	// 自定义字段，只允许 数字、字母和_，限制长度为20
	ExternalInfo string `json:"external_info"`
}
type BuyinInstituteLiveShareParam struct {
	// PID信息
	PidInfo *PidInfo `json:"pid_info"`
	// 抖店用户open_id【废弃，将下线】
	OpenId string `json:"open_id"`
	// 主播百应ID
	BuyinId string `json:"buyin_id"`
	// 是否需要二维码(获取二维码将增加响应耗时，推荐false)
	NeedQrCode bool `json:"need_qr_code"`
	// 直播间口令或者短链接（优先级 buyin_id > open_id > dy_code））
	DyCode string `json:"dy_code"`
	// 直播间绑定的商品id，回流唤起对应商详页
	ProductId int64 `json:"product_id"`
	// 回流端标识 0：抖音 1：抖音极速版
	Platform int32 `json:"platform"`
}

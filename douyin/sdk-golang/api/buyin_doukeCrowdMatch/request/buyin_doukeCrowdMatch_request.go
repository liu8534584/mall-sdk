package buyin_doukeCrowdMatch_request

import (
	"doudian.com/open/sdk_golang/api/buyin_doukeCrowdMatch/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type BuyinDoukeCrowdMatchRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *BuyinDoukeCrowdMatchParam
}

func (c *BuyinDoukeCrowdMatchRequest) GetUrlPath() string {
	return "/buyin/doukeCrowdMatch"
}

func New() *BuyinDoukeCrowdMatchRequest {
	request := &BuyinDoukeCrowdMatchRequest{
		Param: &BuyinDoukeCrowdMatchParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}

func (c *BuyinDoukeCrowdMatchRequest) Execute(accessToken *doudian_sdk.AccessToken) (*buyin_doukeCrowdMatch_response.BuyinDoukeCrowdMatchResponse, error) {
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &buyin_doukeCrowdMatch_response.BuyinDoukeCrowdMatchResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil
}

func (c *BuyinDoukeCrowdMatchRequest) GetParamObject() interface{} {
	return c.Param
}

func (c *BuyinDoukeCrowdMatchRequest) GetParams() *BuyinDoukeCrowdMatchParam {
	return c.Param
}

type DeviceFingerprintInfo struct {
	// 设备指纹ID
	DeviceFingerprintId *string `json:"device_fingerprint_id,omitempty"`
	// 设备指纹类型。1: IDFA，2: IMEI，3: ANDROID_ID，4: GOOGLE_AID，5: OAID
	DeviceFingerprintType *int32 `json:"device_fingerprint_type,omitempty"`
	// 支持md5值，原值toUppercase之后的md5值
	IsMd5 *bool `json:"is_md5,omitempty"`
}
type BuyinDoukeCrowdMatchParam struct {
	// 设备指纹信息
	DeviceFingerprintInfo *DeviceFingerprintInfo `json:"device_fingerprint_info,omitempty"`
}

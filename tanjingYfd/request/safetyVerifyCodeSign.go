package request

import (
	"encoding/json"
	tanjing "github.com/liu8534584/mall-sdk/tanjingYfd"
	"github.com/liu8534584/mall-sdk/tanjingYfd/tanjingResponse"
)

// SafetyVerifyCodeSignRequest 商家机器人登录接口/**
type SafetyVerifyCodeSignRequest struct {
	tanjing.TanjingBaseRequest
	params *SafetyVerifyCodeSignParams
}

func NewSafetyVerifyCodeSignRequest(config *tanjing.TanjingBaseConfig) *SafetyVerifyCodeSignRequest {
	sign := SafetyVerifyCodeSignRequest{
		params: &SafetyVerifyCodeSignParams{},
	}
	sign.SetConfig(config)
	sign.SetClient(tanjing.DefaultTanjingClient)

	return &sign
}

func (s *SafetyVerifyCodeSignRequest) Execute(token string) (*tanjingResponse.TanjingBaseResp, error) {
	execute, err := s.GetClient().Execute(s, token)
	if err != nil {
		return nil, err
	}

	var resp *tanjingResponse.TanjingBaseResp
	_ = json.Unmarshal([]byte(execute), &resp)
	return resp, nil
}

func (s *SafetyVerifyCodeSignRequest) GetApiParamsObject() interface{} {
	return s.params
}

func (s *SafetyVerifyCodeSignRequest) GetApiParams() *SafetyVerifyCodeSignParams {
	return s.params
}

func (s *SafetyVerifyCodeSignRequest) GetApiPath() string {
	return "/scrm/robot/safety-verify-code-sign"
}

type SafetyVerifyCodeSignParams struct {
	MerchantNo string `json:"merchant_no"` //商家编号（必填）
	SerialNo   string `json:"serial_no"`   //操作编号（取本次扫码登录返回的操作编号）
	VerifyCode string `json:"verify_code"` //验证码
}

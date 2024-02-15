package request

import (
	"encoding/json"
	tanjing "github.com/liu8534584/mall-sdk/tanjingYfd"
	"github.com/liu8534584/mall-sdk/tanjingYfd/tanjingResponse"
)

// ScanOutRequest 商家机器人登录接口/**
type ScanOutRequest struct {
	tanjing.TanjingBaseRequest
	params *ScanOutRequestParams
}

func NewScanOutRequest(config *tanjing.TanjingBaseConfig) *ScanOutRequest {
	sign := ScanOutRequest{
		params: &ScanOutRequestParams{},
	}
	sign.SetConfig(config)
	sign.SetClient(tanjing.DefaultTanjingClient)

	return &sign
}

func (s *ScanOutRequest) Execute(token string) (*tanjingResponse.TanjingBaseResp, error) {
	execute, err := s.GetClient().Execute(s, token)
	if err != nil {
		return nil, err
	}

	var resp *tanjingResponse.TanjingBaseResp
	_ = json.Unmarshal([]byte(execute), &resp)
	return resp, nil
}

func (s *ScanOutRequest) GetApiParamsObject() interface{} {
	return s.params
}

func (s *ScanOutRequest) GetApiParams() *ScanOutRequestParams {
	return s.params
}

func (s *ScanOutRequest) GetApiPath() string {
	return "/scrm/robot/sign-out"
}

type ScanOutRequestParams struct {
	MerchantNo    string `json:"merchant_no"`     //商家编号（必填）
	RobotSerialNo string `json:"robot_serial_no"` //机器人编号，第一次登录至开放平台的机器人，机器人编号传空。当机器人第二次登录时必传对应的机器人编号
	Extend        string `json:"extend"`          //透传字段
}

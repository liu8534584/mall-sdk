package request

import (
	"encoding/json"
	tanjing "github.com/liu8534584/mall-sdk/tanjingYfd"
	"github.com/liu8534584/mall-sdk/tanjingYfd/tanjingResponse"
)

// ScanCodeSignRequest 商家机器人登录接口/**
type ScanCodeSignRequest struct {
	tanjing.TanjingBaseRequest
	params *ScanCodeSignRequestParams
}

func NewScanCodeSignRequest(config *tanjing.TanjingBaseConfig) *ScanCodeSignRequest {
	sign := ScanCodeSignRequest{
		params: &ScanCodeSignRequestParams{},
	}
	sign.SetConfig(config)
	sign.SetClient(tanjing.DefaultTanjingClient)

	return &sign
}

func (s *ScanCodeSignRequest) Execute(token string) (*tanjingResponse.ScanCodeSignResp, error) {
	execute, err := s.GetClient().Execute(s, token)
	if err != nil {
		return nil, err
	}

	var resp *tanjingResponse.ScanCodeSignResp
	_ = json.Unmarshal([]byte(execute), &resp)
	return resp, nil
}

func (s *ScanCodeSignRequest) GetApiParamsObject() interface{} {
	return s.params
}

func (s *ScanCodeSignRequest) GetApiParams() *ScanCodeSignRequestParams {
	return s.params
}

func (s *ScanCodeSignRequest) GetApiPath() string {
	return "/scrm/robot/scan-code-sign"
}

type ScanCodeSignRequestParams struct {
	MerchantNo            string `json:"merchant_no"`              //商家编号（必填）
	ProtocolType          string `json:"protocol_type"`            //协议类型（1：IPAD协议 9：PC协议版 12：Mac协议 13：安卓平板协议）（必填）
	IsCacheLogin          bool   `json:"is_cache_login"`           //是否缓存登录
	RobotSerialNo         string `json:"robot_serial_no"`          //机器人编号，第一次登录至开放平台的机器人，机器人编号传空。当机器人第二次登录时必传对应的机器人编号
	RegionCode            string `json:"region_code"`              //扫码设备的地区编码:地级市行政区域代码
	RelationSerialNo      string `json:"relation_serial_no"`       //关联编号
	IsVipDevice           bool   `json:"is_vip_device"`            // 是否VIP设备
	IsOfficialAccountAuth bool   `json:"is_official_account_auth"` //默认传false
}

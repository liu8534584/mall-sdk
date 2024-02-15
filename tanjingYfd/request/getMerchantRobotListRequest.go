package request

import (
	"encoding/json"
	tanjing "github.com/liu8534584/mall-sdk/tanjingYfd"
	"github.com/liu8534584/mall-sdk/tanjingYfd/tanjingResponse"
)

// GetMerchantRobotListRequest 商家机器人登录接口/**
type GetMerchantRobotListRequest struct {
	tanjing.TanjingBaseRequest
	params *GetMerchantRobotListRequestParams
}

func NewGetMerchantRobotListRequest(config *tanjing.TanjingBaseConfig) *GetMerchantRobotListRequest {
	sign := GetMerchantRobotListRequest{
		params: &GetMerchantRobotListRequestParams{},
	}
	sign.SetConfig(config)
	sign.SetClient(tanjing.DefaultTanjingClient)

	return &sign
}

func (s *GetMerchantRobotListRequest) Execute(token string) (*tanjingResponse.TanjingBaseResp, error) {
	execute, err := s.GetClient().Execute(s, token)
	if err != nil {
		return nil, err
	}

	var resp *tanjingResponse.TanjingBaseResp
	_ = json.Unmarshal([]byte(execute), &resp)
	return resp, nil
}

func (s *GetMerchantRobotListRequest) GetApiParamsObject() interface{} {
	return s.params
}

func (s *GetMerchantRobotListRequest) GetApiParams() *GetMerchantRobotListRequestParams {
	return s.params
}

func (s *GetMerchantRobotListRequest) GetApiPath() string {
	return "/scrm/Robot/get-merchant-robot-page-list"
}

type GetMerchantRobotListRequestParams struct {
	MerchantNo     string   `json:"merchant_no"`        //商家编号（必填）
	RobotSerialNos []string `json:"robot_serial_nos"`   //机器人编号，第一次登录至开放平台的机器人，机器人编号传空。当机器人第二次登录时必传对应的机器人编号
	Type           int      `json:"chatroom_serial_no"` //机器人类型 0全部 10平台号 20托管号 30扫码号
	PageSize       int      `json:"page_size"`          //分页大小
	MaxId          int      `json:"max_id"`             // 最大机器人id，用于分页
}

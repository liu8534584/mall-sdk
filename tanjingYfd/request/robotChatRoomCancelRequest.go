package request

import (
	"encoding/json"
	tanjing "github.com/liu8534584/mall-sdk/tanjingYfd"
	"github.com/liu8534584/mall-sdk/tanjingYfd/tanjingResponse"
)

// RobotChatRoomCancelRequest 机器人取消关注群/**
type RobotChatRoomCancelRequest struct {
	tanjing.TanjingBaseRequest
	params *RobotChatRoomCancelRequestParams
}

func NewRobotChatRoomCancelRequest(config *tanjing.TanjingBaseConfig) *RobotChatRoomCancelRequest {
	sign := RobotChatRoomCancelRequest{
		params: &RobotChatRoomCancelRequestParams{},
	}
	sign.SetConfig(config)
	sign.SetClient(tanjing.DefaultTanjingClient)

	return &sign
}

func (r *RobotChatRoomCancelRequest) Execute(token string) (*tanjingResponse.TanjingBaseResp, error) {
	execute, err := r.GetClient().Execute(r, token)
	if err != nil {
		return nil, err
	}

	var resp *tanjingResponse.TanjingBaseResp
	_ = json.Unmarshal([]byte(execute), &resp)
	return resp, nil
}

func (r *RobotChatRoomCancelRequest) GetApiParamsObject() interface{} {
	return r.params
}

func (r *RobotChatRoomCancelRequest) GetApiParams() *RobotChatRoomCancelRequestParams {
	return r.params
}

func (r *RobotChatRoomCancelRequest) GetApiPath() string {
	return "/sync/ChatRoom/RobotChatRoomCancel"
}

type RobotChatRoomCancelRequestParams struct {
	VcMerchantNo       string `json:"vcMerchantNo"`       //商家编号
	VcChatRoomSerialNo string `json:"vcChatRoomSerialNo"` //群编号
	VcRobotSerialNo    string `json:"vcRobotSerialNo"`    //机器人编号
}

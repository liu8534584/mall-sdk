package request

import (
	"encoding/json"
	tanjing "github.com/liu8534584/mall-sdk/tanjingYfd"
	"github.com/liu8534584/mall-sdk/tanjingYfd/tanjingResponse"
)

// RobotChatRoomOpen 机器人关注群/**
type RobotChatRoomOpen struct {
	tanjing.TanjingBaseRequest
	params *RobotChatRoomOpenParams
}

func NewRobotChatRoomOpen(config *tanjing.TanjingBaseConfig) *RobotChatRoomOpen {
	sign := RobotChatRoomOpen{
		params: &RobotChatRoomOpenParams{},
	}
	sign.SetConfig(config)
	sign.SetClient(tanjing.DefaultTanjingClient)

	return &sign
}

func (r *RobotChatRoomOpen) Execute(token string) (*tanjingResponse.TanjingBaseResp, error) {
	execute, err := r.GetClient().Execute(r, token)
	if err != nil {
		return nil, err
	}

	var resp *tanjingResponse.TanjingBaseResp
	_ = json.Unmarshal([]byte(execute), &resp)
	return resp, nil
}

func (r *RobotChatRoomOpen) GetApiParamsObject() interface{} {
	return r.params
}

func (r *RobotChatRoomOpen) GetApiParams() *RobotChatRoomOpenParams {
	return r.params
}

func (r *RobotChatRoomOpen) GetApiPath() string {
	return "/sync/ChatRoom/RobotChatRoomOpen"
}

type RobotChatRoomOpenParams struct {
	VcMerchantNo       string `json:"vcMerchantNo"`       //商家编号
	VcChatRoomSerialNo string `json:"vcChatRoomSerialNo"` //群编号
	VcRobotSerialNo    string `json:"vcRobotSerialNo"`    //机器人编号
}

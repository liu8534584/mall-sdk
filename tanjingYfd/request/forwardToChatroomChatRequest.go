package request

import (
	"encoding/json"
	tanjing "github.com/liu8534584/mall-sdk/tanjingYfd"
	"github.com/liu8534584/mall-sdk/tanjingYfd/tanjingResponse"
)

// ForwardToChatroomChat 商家机器人登录接口/**
type ForwardToChatroomChat struct {
	tanjing.TanjingBaseRequest
	params *ForwardToChatroomChatParams
}

func NewForwardToChatroomChat(config *tanjing.TanjingBaseConfig) *ForwardToChatroomChat {
	sign := ForwardToChatroomChat{
		params: &ForwardToChatroomChatParams{},
	}
	sign.SetConfig(config)
	sign.SetClient(tanjing.DefaultTanjingClient)

	return &sign
}

func (s *ForwardToChatroomChat) Execute(token string) (*tanjingResponse.TanjingBaseResp, error) {
	execute, err := s.GetClient().Execute(s, token)
	if err != nil {
		return nil, err
	}

	var resp *tanjingResponse.TanjingBaseResp
	_ = json.Unmarshal([]byte(execute), &resp)
	return resp, nil
}

func (s *ForwardToChatroomChat) GetApiParamsObject() interface{} {
	return s.params
}

func (s *ForwardToChatroomChat) GetApiParams() *ForwardToChatroomChatParams {
	return s.params
}

func (s *ForwardToChatroomChat) GetApiPath() string {
	return "/scrm/ChatMessages/forward-to-chatroom-chat"
}

type ForwardToChatroomChatParams struct {
	MerchantNo       string `json:"merchant_no"`        //商家编号（必填）
	RobotSerialNo    string `json:"robot_serial_no"`    //机器人编号，第一次登录至开放平台的机器人，机器人编号传空。当机器人第二次登录时必传对应的机器人编号
	ChatroomSerialNo string `json:"chatroom_serial_no"` //群编号
	MsgType          int64  `json:"msg_type"`           //消息类型（2002：图片；2004：视频；2007表情包）
	Content          string `json:"content"`            //消息xml
}

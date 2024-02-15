package request

import (
	"encoding/json"
	tanjing "github.com/liu8534584/mall-sdk/tanjingYfd"
	"github.com/liu8534584/mall-sdk/tanjingYfd/tanjingResponse"
)

// SendPrivateChatMessagesRequest 私聊信息发送/**
type SendPrivateChatMessagesRequest struct {
	tanjing.TanjingBaseRequest
	params *SendPrivateChatMessagesRequestParams
}

func NewSendPrivateChatMessagesRequest(config *tanjing.TanjingBaseConfig) *SendPrivateChatMessagesRequest {
	sign := SendPrivateChatMessagesRequest{
		params: &SendPrivateChatMessagesRequestParams{},
	}
	sign.SetConfig(config)
	sign.SetClient(tanjing.DefaultTanjingClient)

	return &sign
}

func (r *SendPrivateChatMessagesRequest) Execute(token string) (*tanjingResponse.TanjingBaseResp, error) {
	execute, err := r.GetClient().Execute(r, token)
	if err != nil {
		return nil, err
	}

	var resp *tanjingResponse.TanjingBaseResp
	_ = json.Unmarshal([]byte(execute), &resp)
	return resp, nil
}

func (r *SendPrivateChatMessagesRequest) GetApiParamsObject() interface{} {
	return r.params
}

func (r *SendPrivateChatMessagesRequest) GetApiParams() *SendPrivateChatMessagesRequestParams {
	return r.params
}

func (r *SendPrivateChatMessagesRequest) GetApiPath() string {
	return "/scrm/ChatMessages/SendPrivateChatMessages"
}

type SendPrivateChatMessagesRequestParams struct {
	VcMerchantNo    string             `json:"vcMerchantNo"`    //商家编号
	VcRobotSerialNo string             `json:"vcRobotSerialNo"` //机器人编号
	VcRelaSerialNo  string             `json:"vcRelaSerialNo"`  //商家业务流水号
	VcToWxSerialNo  string             `json:"vcToWxSerialNo"`  //微信编号
	Data            []SendGroupMsgInfo `json:"Data"`
}

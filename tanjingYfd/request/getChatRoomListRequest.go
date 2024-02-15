package request

import (
	"encoding/json"
	tanjing "github.com/liu8534584/mall-sdk/tanjingYfd"
	"github.com/liu8534584/mall-sdk/tanjingYfd/tanjingResponse"
)

// GetChatRoomListRequest 获取群列表/**
type GetChatRoomListRequest struct {
	tanjing.TanjingBaseRequest
	params *GetChatRoomListRequestParams
}

func NewGetChatRoomListRequest(config *tanjing.TanjingBaseConfig) *GetChatRoomListRequest {
	sign := GetChatRoomListRequest{
		params: &GetChatRoomListRequestParams{},
	}
	sign.SetConfig(config)
	sign.SetClient(tanjing.DefaultTanjingClient)

	return &sign
}

func (r *GetChatRoomListRequest) Execute(token string) (*tanjingResponse.GetChatRoomListResp, error) {
	execute, err := r.GetClient().Execute(r, token)
	if err != nil {
		return nil, err
	}

	var resp *tanjingResponse.GetChatRoomListResp
	_ = json.Unmarshal([]byte(execute), &resp)
	return resp, nil
}

func (r *GetChatRoomListRequest) GetApiParamsObject() interface{} {
	return r.params
}

func (r *GetChatRoomListRequest) GetApiParams() *GetChatRoomListRequestParams {
	return r.params
}

func (r *GetChatRoomListRequest) GetApiPath() string {
	return "/sync/ChatRoom/GetChatRoomList"
}

type GetChatRoomListRequestParams struct {
	VcMerchantNo       string `json:"vcMerchantNo"`       //商家编号
	VcChatRoomSerialNo string `json:"vcChatRoomSerialNo"` //群编号，值不传的话，查询全部
	IsOpenMessage      uint8  `json:"isOpenMessage"`      //是否已关注（10 是 11 否）, 0 查询全部
	VcRobotSerialNo    string `json:"vcRobotSerialNo"`    // 机器人编号
}

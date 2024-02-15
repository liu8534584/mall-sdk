package request

import (
	"encoding/json"
	tanjing "github.com/liu8534584/mall-sdk/tanjingYfd"
	"github.com/liu8534584/mall-sdk/tanjingYfd/tanjingResponse"
)

// GetTagInFriendsRequest 获取标签列表/**
type GetTagInFriendsRequest struct {
	tanjing.TanjingBaseRequest
	params *GetTagInFriendsRequestParams
}

func NewGetTagInFriendsRequest(config *tanjing.TanjingBaseConfig) *GetTagInFriendsRequest {
	sign := GetTagInFriendsRequest{
		params: &GetTagInFriendsRequestParams{},
	}
	sign.SetConfig(config)
	sign.SetClient(tanjing.DefaultTanjingClient)

	return &sign
}

func (r *GetTagInFriendsRequest) Execute(token string) (*tanjingResponse.GetTagInFriendsResp, error) {
	execute, err := r.GetClient().Execute(r, token)
	if err != nil {
		return nil, err
	}

	var resp *tanjingResponse.GetTagInFriendsResp
	_ = json.Unmarshal([]byte(execute), &resp)
	return resp, nil
}

func (r *GetTagInFriendsRequest) GetApiParamsObject() interface{} {
	return r.params
}

func (r *GetTagInFriendsRequest) GetApiParams() *GetTagInFriendsRequestParams {
	return r.params
}

func (r *GetTagInFriendsRequest) GetApiPath() string {
	return "/sync/Friend/GetTagInFriends"
}

type GetTagInFriendsRequestParams struct {
	VcMerchantNo    string `json:"vcMerchantNo"`    //商家编号
	VcRobotSerialNo string `json:"vcRobotSerialNo"` //机器人编号
	NTagId          string `json:"nTagId"`          // 标签id
}

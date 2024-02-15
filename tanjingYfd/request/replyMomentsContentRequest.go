package request

import (
	"encoding/json"
	tanjing "github.com/liu8534584/mall-sdk/tanjingYfd"
	"github.com/liu8534584/mall-sdk/tanjingYfd/tanjingResponse"
)

// ReplyMomentsContentRequest 朋友圈评论/**
type ReplyMomentsContentRequest struct {
	tanjing.TanjingBaseRequest
	params *ReplyMomentsContentRequestParams
}

func NewReplyMomentsContentRequest(config *tanjing.TanjingBaseConfig) *ReplyMomentsContentRequest {
	sign := ReplyMomentsContentRequest{
		params: &ReplyMomentsContentRequestParams{},
	}
	sign.SetConfig(config)
	sign.SetClient(tanjing.DefaultTanjingClient)

	return &sign
}

func (r *ReplyMomentsContentRequest) Execute(token string) (*tanjingResponse.TanjingBaseResp, error) {
	execute, err := r.GetClient().Execute(r, token)
	if err != nil {
		return nil, err
	}

	var resp *tanjingResponse.TanjingBaseResp
	_ = json.Unmarshal([]byte(execute), &resp)
	return resp, nil
}

func (r *ReplyMomentsContentRequest) GetApiParamsObject() interface{} {
	return r.params
}

func (r *ReplyMomentsContentRequest) GetApiParams() *ReplyMomentsContentRequestParams {
	return r.params
}

func (r *ReplyMomentsContentRequest) GetApiPath() string {
	return "/scrm/Moments/ReplyMomentsContent"
}

type ReplyMomentsContentRequestParams struct {
	VcMerchantNo    string `json:"vcMerchantNo"`    //商家编号
	VcRobotSerialNo string `json:"vcRobotSerialNo"` //机器人编号
	NTimelineId     string `json:"nTimelineId"`     //朋友圈ID
	NReplyId        int64  `json:"nReplyId"`        //评论ID，默认为0
	VcReplyMsg      string `json:"vcReplyMsg"`      // 评论内容
}

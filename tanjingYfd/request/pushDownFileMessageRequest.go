package request

import (
	"encoding/json"
	tanjing "github.com/liu8534584/mall-sdk/tanjingYfd"
	"github.com/liu8534584/mall-sdk/tanjingYfd/tanjingResponse"
)

// PushDownFileMessageRequest 机器人取消关注群/**
type PushDownFileMessageRequest struct {
	tanjing.TanjingBaseRequest
	params *PushDownFileMessageRequestParams
}

func NewPushDownFileMessageRequest(config *tanjing.TanjingBaseConfig) *PushDownFileMessageRequest {
	sign := PushDownFileMessageRequest{
		params: &PushDownFileMessageRequestParams{},
	}
	sign.SetConfig(config)
	sign.SetClient(tanjing.DefaultTanjingClient)

	return &sign
}

func (r *PushDownFileMessageRequest) Execute(token string) (*tanjingResponse.PushDownFileMessageResp, error) {
	execute, err := r.GetClient().Execute(r, token)
	if err != nil {
		return nil, err
	}

	var resp *tanjingResponse.PushDownFileMessageResp
	_ = json.Unmarshal([]byte(execute), &resp)
	return resp, nil
}

func (r *PushDownFileMessageRequest) GetApiParamsObject() interface{} {
	return r.params
}

func (r *PushDownFileMessageRequest) GetApiParams() *PushDownFileMessageRequestParams {
	return r.params
}

func (r *PushDownFileMessageRequest) GetApiPath() string {
	return "/scrm/ChatMessages/PushDownFileMessage"
}

type PushDownFileMessageRequestParams struct {
	VcMerchantNo       string `json:"vcMerchantNo"`       //商家编码
	VcRobotSerialNo    string `json:"vcRobotSerialNo"`    //机器人编号
	VcDownFileSerialNo string `json:"vcDownFileSerialNo"` //下载文件编码
	NVideoDownType     string `json:"nVideoDownType"`     //下载视频/视频封面图/语音文件/图片 （1：下载视频封面/图片文件/语音（视频封面图：只下载视频的第一帧图片），2，下载视频，3 下载视频和视频封面），0 其他类型
}

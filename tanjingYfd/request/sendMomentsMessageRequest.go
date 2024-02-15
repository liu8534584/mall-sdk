package request

import (
	"encoding/json"
	tanjing "github.com/liu8534584/mall-sdk/tanjingYfd"
	"github.com/liu8534584/mall-sdk/tanjingYfd/tanjingResponse"
)

// SendMomentsMessageRequest 发送朋友圈/**
type SendMomentsMessageRequest struct {
	tanjing.TanjingBaseRequest
	params *SendMomentsMessageRequestParams
}

func NewSendMomentsMessageRequest(config *tanjing.TanjingBaseConfig) *SendMomentsMessageRequest {
	sign := SendMomentsMessageRequest{
		params: &SendMomentsMessageRequestParams{},
	}
	sign.SetConfig(config)
	sign.SetClient(tanjing.DefaultTanjingClient)

	return &sign
}

func (r *SendMomentsMessageRequest) Execute(token string) (*tanjingResponse.TanjingBaseResp, error) {
	execute, err := r.GetClient().Execute(r, token)
	if err != nil {
		return nil, err
	}

	var resp *tanjingResponse.TanjingBaseResp
	_ = json.Unmarshal([]byte(execute), &resp)
	return resp, nil
}

func (r *SendMomentsMessageRequest) GetApiParamsObject() interface{} {
	return r.params
}

func (r *SendMomentsMessageRequest) GetApiParams() *SendMomentsMessageRequestParams {
	return r.params
}

func (r *SendMomentsMessageRequest) GetApiPath() string {
	return "/scrm/Moments/SendMomentsMessage"
}

type SendMomentsMessageRequestParams struct {
	VcMerchantNo       string                `json:"vcMerchantNo"`    //商家编号
	VcRobotSerialNo    string                `json:"vcRobotSerialNo"` //机器人编号
	NMomentType        int                   `json:"nMomentType"`     //消息类型（纯文字：2001 图文：2002 视频：2004 链接：2005）
	VcContentDesc      string                `json:"vcContentDesc"`   //文字内容
	VcTitle            string                `json:"vcTitle"`         // 链接标题
	VcContentUrl       string                `json:"vcContentUrl"`    // 链接地址
	MediaList          []MomentsMessageMedia `json:"MediaList"`
	Tags               []MomentsMessageTag   `json:"Tags"`
	VcContactSerialNos string                `json:"vcContactSerialNos"` // 指定人的编号
	NPrivateType       int                   `json:"nPrivateType"`       //可见类型 0：默认无限制 1：指定不可见 2：指定可见 3：仅自己可见
}

type MomentsMessageTag struct {
	NTagId    string `json:"nTagId"`    //标签ID
	VcTagName string `json:"vcTagName"` //标签名称
}

type MomentsMessageMedia struct {
	VcMediaUrl     string `json:"vcMediaUrl"`     //图片/视频地址（图片最多传9张）
	NVideoDuration string `json:"nVideoDuration"` //视频时长,时长单位：秒；必须传时长且时长要正确，否则会发送失败，当时长不正确时可能会有很大的禁封风险
	VcCoverImgUrl  string `json:"vcCoverImgUrl"`  //视频封面图（第一帧），此参数不传会报错提示：视频类型朋友圈必须传有效的封面图片地址
}

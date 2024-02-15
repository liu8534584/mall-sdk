package request

import (
	"encoding/json"
	tanjing "github.com/liu8534584/mall-sdk/tanjingYfd"
	"github.com/liu8534584/mall-sdk/tanjingYfd/tanjingResponse"
)

// SendGroupChatMessagesRequest 群消息发送/**
type SendGroupChatMessagesRequest struct {
	tanjing.TanjingBaseRequest
	params *SendGroupChatMessagesRequestParams
}

func NewSendGroupChatMessagesRequest(config *tanjing.TanjingBaseConfig) *SendGroupChatMessagesRequest {
	sign := SendGroupChatMessagesRequest{
		params: &SendGroupChatMessagesRequestParams{},
	}
	sign.SetConfig(config)
	sign.SetClient(tanjing.DefaultTanjingClient)

	return &sign
}

func (r *SendGroupChatMessagesRequest) Execute(token string) (*tanjingResponse.TanjingBaseResp, error) {
	execute, err := r.GetClient().Execute(r, token)
	if err != nil {
		return nil, err
	}

	var resp *tanjingResponse.TanjingBaseResp
	_ = json.Unmarshal([]byte(execute), &resp)
	return resp, nil
}

func (r *SendGroupChatMessagesRequest) GetApiParamsObject() interface{} {
	return r.params
}

func (r *SendGroupChatMessagesRequest) GetApiParams() *SendGroupChatMessagesRequestParams {
	return r.params
}

func (r *SendGroupChatMessagesRequest) GetApiPath() string {
	return "/scrm/ChatMessages/SendGroupChatMessages"
}

type SendGroupChatMessagesRequestParams struct {
	VcMerchantNo       string             `json:"vcMerchantNo"`       //商家编号
	VcRobotSerialNo    string             `json:"vcRobotSerialNo"`    //机器人编号
	VcRelaSerialNo     string             `json:"vcRelaSerialNo"`     //商家业务流水号
	VcChatRoomSerialNo string             `json:"vcChatRoomSerialNo"` //群编号
	Data               []SendGroupMsgInfo `json:"Data"`
}

type SendGroupMsgInfo struct {
	NMsgNum  int `json:"nMsgNum"`  // 消息编号(整型,用于区分同一组的消息，消息编号的值需要为正整数)
	NMsgType int `json:"nMsgType"` //消息类型 2001 文字 2002 图片 2003 语音(只支持amr格式) 2004 视频 2005 链接 2006 好友名片 2008 企微好友名片 2010 文件 2013小程序 2016 音乐 2017 视频号消息 2019 视频号直播间 2020 视频号名片 2024 聊天记录消息 2025 笔记消息
	//文字内容（如果是图片或者链接则传图片地址[链接的图片不宜过大，建议160x160px，小于10k]
	//如果是（企微）好友名片，则个微传好友的微信编号（企微传好友的微信id）
	//如果是语音,则传语音的地址（语音的后缀必须为amr示例：http://downsc.chinaz.net/Files/DownLoadsound1/201910/12087.amr） 如果是小程序则传小程序的XML
	//如果是视频消息，则传视频的封面图【视频第一帧的图片链接地址，此处必传，否则视频消息类型会失败】）（@人任意位置文本消息：xxxxxxxxx@$好友编号$xxxxxxxxxxxxx@$好友编号$xxxxx（ps：xxx代表文本消息，@$好友编号$代表@的群内好友））
	//如果是视频号消息/视频号直播间消息/视频号名片，则传对应的XML；如果是笔记消息/聊天记录消息，则传该笔记消息/聊天记录消息的base64编码xml；
	MsgContent      string   `json:"msgContent"`
	NVoiceTime      int64    `json:"nVoiceTime"`      //语音时长/视频时长,时长单位：秒；当消息类型为以上两种类型时，必须传时长且时长要正确，否则会发送失败，当时长不正确时可能会有很大的禁封风险
	VcTitle         string   `json:"vcTitle"`         //链接标题
	VcDesc          string   `json:"vcDesc"`          //链接描述
	VcHref          string   `json:"vcHref"`          //链接URL，当消息为视频时，此处传视频的链接地址
	NIsHit          int      `json:"nIsHit"`          //是否艾特必填，值为空或则不为以下提供的值，调用接口可以成功，但是消息无法发送至群内。数组内仅文本消息或空消息支持@人 (0 不@人 2 @部分群成员)
	NAtLocation     int      `json:"nAtLocation"`     //@人在文本的所在位置 0 文本开始位置 1文本结束位置 2任意位置
	VcAtWxSerialNos []string `json:"vcAtWxSerialNos"` //指定艾特部分人的编号(多个用','隔开,如果不用艾特则传空)
}

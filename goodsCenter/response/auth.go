package response

import "github.com/liu8534584/mall-sdk/goodsCenter/response/base"

type AuthStatusRes struct {
	AppId        string     `json:"appId"`
	MiniAppUrl   string     `json:"miniAppUrl"`
	OauthUrl     string     `json:"oauthUrl"`
	OauthHttpUrl string     `json:"oauthHttpUrl"`
	IsNeedOauth  bool       `json:"isNeedOauth"`
	PddPid       PddPidInfo `json:"pddPid"`
}

type PddPidInfo struct {
	PddPid          string `json:"pddPid"`
	Uid             string `json:"uid"`
	Sid             string `json:"sid"`
	PddCustomParams string `json:"pddCustomParams"`
}

type AuthStatusResp struct {
	Code     base.Code     `json:"code"`
	Message  string        `json:"msg"`
	Demotion base.Demotion `json:"demotion"`
	Data     AuthStatusRes `json:"data"`
}

package vipLinkCheckRequest

import (
	"encoding/json"
	"github.com/liu8534584/mall-sdk/vip"
	"github.com/liu8534584/mall-sdk/vip/vipAdpApiOpenServiceUnionUrlService/vipLinkCheckResponse"
)

type UnionUrlServiceVipLinkCheckRequest struct {
	vip.BaseVipApiRequest
	Param *UnionUrlServiceVipLinkCheckParam
}

func (v *UnionUrlServiceVipLinkCheckRequest) GetServiceName() string {
	return "com.vip.adp.api.open.service.UnionUrlService"
}

func (v *UnionUrlServiceVipLinkCheckRequest) GetMethodName() string {
	return "vipLinkCheck"
}

func New(config *vip.VipBaseConfig) *UnionUrlServiceVipLinkCheckRequest {
	request := UnionUrlServiceVipLinkCheckRequest{
		Param: &UnionUrlServiceVipLinkCheckParam{},
	}
	request.SetConfig(config)
	request.SetClient(vip.DefaultVipApiClient)

	return &request
}

func (v *UnionUrlServiceVipLinkCheckRequest) Execute() (*vipLinkCheckResponse.VipLinkCheckResponse, error) {
	responseJson, err := v.GetClient().Request(v, false)
	if err != nil {
		return nil, err
	}
	NewResponseObj := &vipLinkCheckResponse.VipLinkCheckResponse{}
	_ = json.Unmarshal([]byte(responseJson), &NewResponseObj)

	return NewResponseObj, err
}

func (v *UnionUrlServiceVipLinkCheckRequest) GetParamsObject() interface{} {
	return v.Param
}

func (v *UnionUrlServiceVipLinkCheckRequest) GetParams() *UnionUrlServiceVipLinkCheckParam {
	return v.Param
}

type UnionUrlServiceVipLinkCheckParam struct {
	VipLinkCheckReq VipLinkCheckReq `json:"vipLinkCheckReq"`
}
type VipLinkCheckReq struct {
	Source    string `json:"source"`    //请求源：调用方自定义，标识请求应用即可，无业务含义
	Content   string `json:"content"`   //检查的内容(长度不超过10000)
	RequestId string `json:"requestId"` //请求id：UUID
}

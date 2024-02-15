package genByGoodsIdRequest

import (
	"encoding/json"
	"github.com/liu8534584/mall-sdk/vip"
	"github.com/liu8534584/mall-sdk/vip/vipAdpApiOpenServiceUnionUrlService/genByGoodsIdResponse"
)

type UnionUrlServiceGenByGoodsIdRequest struct {
	vip.BaseVipApiRequest
	Param *UnionUrlServiceGenByGoodsIdParam
}

func (v *UnionUrlServiceGenByGoodsIdRequest) GetServiceName() string {
	return "com.vip.adp.api.open.service.UnionUrlService"
}

func (v *UnionUrlServiceGenByGoodsIdRequest) GetMethodName() string {
	return "genByGoodsId"
}

func New(config *vip.VipBaseConfig) *UnionUrlServiceGenByGoodsIdRequest {
	request := UnionUrlServiceGenByGoodsIdRequest{
		Param: &UnionUrlServiceGenByGoodsIdParam{},
	}
	request.SetConfig(config)
	request.SetClient(vip.DefaultVipApiClient)

	return &request
}

func (v *UnionUrlServiceGenByGoodsIdRequest) Execute() (*genByGoodsIdResponse.GenByGoodsIdResponse, error) {
	responseJson, err := v.GetClient().Request(v, false)
	if err != nil {
		return nil, err
	}
	NewResponseObj := &genByGoodsIdResponse.GenByGoodsIdResponse{}
	_ = json.Unmarshal([]byte(responseJson), &NewResponseObj)

	return NewResponseObj, err
}

func (v *UnionUrlServiceGenByGoodsIdRequest) GetParamsObject() interface{} {
	return v.Param
}

func (v *UnionUrlServiceGenByGoodsIdRequest) GetParams() *UnionUrlServiceGenByGoodsIdParam {
	return v.Param
}

type UnionUrlServiceGenByGoodsIdParam struct {
	GoodsIdList            []string               `json:"goodsIdList"`
	StatParam              string                 `json:"statParam"` //自定义渠道统计参数
	ChanTag                string                 `json:"chanTag"`
	RequestId              string                 `json:"requestId"`   //请求id：UUID
	GenShortUrl            bool                   `json:"genShortUrl"` //是否压缩生成的链接,默认false(理论上生成的链接无需压缩也能满足大部分推广情景，如非必要，请勿选择压缩，压缩的链接有有效期(目前是一个月))
	UrlGenByGoodsIdRequest UrlGenByGoodsIdRequest `json:"urlGenByGoodsIdRequest"`
}
type UrlGenByGoodsIdRequest struct {
	OpenId   string `json:"openId"`
	RealCall bool   `json:"realCall"`
	AdCode   string `json:"adCode"`
	Platform int32  `json:"platform"`
}

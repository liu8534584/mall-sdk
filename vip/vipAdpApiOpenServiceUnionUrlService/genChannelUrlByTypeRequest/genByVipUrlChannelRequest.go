package genChennelUrlRequest

import (
	"encoding/json"
	"github.com/liu8534584/mall-sdk/vip"
	genChennelUrlResponse "github.com/liu8534584/mall-sdk/vip/vipAdpApiOpenServiceUnionUrlService/genChannelUrlByTypeResponse"
)

type UnionUrlServiceGenChannelUrlByTypeRequest struct {
	vip.BaseVipApiRequest
	Param *UnionUrlServiceGenChannelUrlByTypeParam
}

func (v *UnionUrlServiceGenChannelUrlByTypeRequest) GetServiceName() string {
	return "com.vip.adp.api.open.service.UnionUrlService"
}

func (v *UnionUrlServiceGenChannelUrlByTypeRequest) GetMethodName() string {
	return "getChannelUrlByType"
}

func New(config *vip.VipBaseConfig) *UnionUrlServiceGenChannelUrlByTypeRequest {
	request := UnionUrlServiceGenChannelUrlByTypeRequest{
		Param: &UnionUrlServiceGenChannelUrlByTypeParam{},
	}
	request.SetConfig(config)
	request.SetClient(vip.DefaultVipApiClient)

	return &request
}

func (v *UnionUrlServiceGenChannelUrlByTypeRequest) Execute() (*genChennelUrlResponse.VipGenByChannelTypeResponse, error) {
	responseJson, err := v.GetClient().Request(v, false)
	if err != nil {
		return nil, err
	}
	NewResponseObj := &genChennelUrlResponse.VipGenByChannelTypeResponse{}
	_ = json.Unmarshal([]byte(responseJson), &NewResponseObj)

	return NewResponseObj, err
}

func (v *UnionUrlServiceGenChannelUrlByTypeRequest) GetParamsObject() interface{} {
	return v.Param
}

func (v *UnionUrlServiceGenChannelUrlByTypeRequest) GetParams() *UnionUrlServiceGenChannelUrlByTypeParam {
	return v.Param
}

type UnionUrlServiceGenChannelUrlByTypeParam struct {
	Request UnionUrlServiceGenChannelUrlByTypeRequestFiled `json:"request"`
}

type UnionUrlServiceGenChannelUrlByTypeRequestFiled struct {
	Type                   string `json:"type"`                   //否  生链类型
	Ucode                  string `json:"ucode"`                  //  否 渠道用户码
	CommissionDiscountRate string `json:"commissionDiscountRate"` //否 佣金打折系数（0-1000对应0%-100%）
	SubsidyDiscountRate    string `json:"subsidyDiscountRate"`    //否 补贴打折系数（0-1000对应0%-100%）
	StatParam              string `json:"statParam"`              // 否 渠道用户标识
	ChanTag                string `json:"chanTag"`                //否 chanTag=pid，即推广位标识 (必传)，用来标记推广中的某个资源位，比如APP的banner、icon等（必传）（不能含有特殊字符，仅限字母、数字、下划线， 长度最大64）， 如果没有推广位 则传默认推广位标识: default_pid
	RequestId              string `json:"requestId"`              //否 请求id：UUID
	UserId                 string `json:"userId"`                 // 否 用户id
	AppKey                 string `json:"appKey"`                 //否 渠道appKey
	CompressShortUrl       bool   `json:"compressShortUrl"`       //否 是否压缩链接
	GenerateByUcode        bool   `json:"generateByUcode"`        // 否 是否根据传入ucode转链
	OpenId                 string `json:"openId"`                 // 否 渠道用户在渠道侧的用户唯一标识（必传）（如已接入sdk，该值为渠道用户授权绑定唯品会账号的标识 ，如没接入sdk，该值为渠道用户在渠道侧的用户唯一标识） （不能含有特殊字符，仅限字母、数字、下划线， 长度最大32），用于识别用户并给用户返利分佣（非常重要！传错影响给用户返利分佣），如果当前调用与用户无关 比如后台job触发拉取商品 则传默认标识：default_open_id (调转链接口与商品接口时，同一用户openId传参需要一致)。
	RealCall               bool   `json:"realCall"`               // 否 是否实时调用（必传），true：由用户实时触发的请求，实时给用户展示联盟返回的商品信息或者实时给用户转链生成推广链接。 false：不是由用户实时触发，由渠道后台job触发的请求，比如渠道后台job定期调联盟接口拉取商品到渠道自己的库，请按实际情况传该参数。
}

package goodsListRequest

import (
	"encoding/json"
	"github.com/liu8534584/mall-sdk/vip"
	"github.com/liu8534584/mall-sdk/vip/vipAdpApiOpenServiceUnionGoodsService/goodsListResponse"
)

type UnionGoodsServiceGoodsListRequest struct {
	vip.BaseVipApiRequest
	Param *UnionGoodsServiceGoodsListParam
}

func (v *UnionGoodsServiceGoodsListRequest) GetServiceName() string {
	return "com.vip.adp.api.open.service.UnionGoodsService"
}

func (v *UnionGoodsServiceGoodsListRequest) GetMethodName() string {
	return "goodsList"
}

func New(config *vip.VipBaseConfig) *UnionGoodsServiceGoodsListRequest {
	request := UnionGoodsServiceGoodsListRequest{
		Param: &UnionGoodsServiceGoodsListParam{},
	}
	request.SetConfig(config)
	request.SetClient(vip.DefaultVipApiClient)

	return &request
}

func (v *UnionGoodsServiceGoodsListRequest) Execute() (*goodsListResponse.UnionGoodsServiceGoodsListResponse, error) {
	responseJson, err := v.GetClient().Request(v, false)
	if err != nil {
		return nil, err
	}
	NewResponseObj := &goodsListResponse.UnionGoodsServiceGoodsListResponse{}
	_ = json.Unmarshal([]byte(responseJson), &NewResponseObj)

	return NewResponseObj, err
}

func (v *UnionGoodsServiceGoodsListRequest) GetParamsObject() interface{} {
	return v.Param
}

func (v *UnionGoodsServiceGoodsListRequest) GetParams() *UnionGoodsServiceGoodsListParam {
	return v.Param
}

type UnionGoodsServiceGoodsListParam struct {
	Request RequestObj `json:"request"`
}
type RequestObj struct {
	ChannelType int    `json:"channelType"` //频道类型:0-超高佣，1-出单爆款; 当请求类型为频道时必传
	Page        int    `json:"page"`
	PageSize    int    `json:"pageSize"` //分页大小:默认20，最大100
	RequestId   string `json:"requestId"`
	SourceType  int    `json:"sourceType"` //请求源类型：0-频道，1-组货
	JxCode      string `json:"jxCode"`     //精选组货码：当请求源类型为组货时必传
	ChanTag     string `json:"chanTag"`
	OpenId      string `json:"openId"`
	RealCall    bool   `json:"realCall"`
}

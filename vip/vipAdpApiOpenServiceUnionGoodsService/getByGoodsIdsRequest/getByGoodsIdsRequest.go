package getByGoodsIdsRequest

import (
	"encoding/json"
	"github.com/liu8534584/mall-sdk/vip"
	"github.com/liu8534584/mall-sdk/vip/vipAdpApiOpenServiceUnionGoodsService"
	"github.com/liu8534584/mall-sdk/vip/vipAdpApiOpenServiceUnionGoodsService/getByGoodsIdsResponse"
)

type UnionGoodsServiceGetByGoodsIdsRequest struct {
	vip.BaseVipApiRequest
	Param *UnionGoodsServiceGetByGoodsIdsParam
}

func (v *UnionGoodsServiceGetByGoodsIdsRequest) GetServiceName() string {
	return "com.vip.adp.api.open.service.UnionGoodsService"
}

func (v *UnionGoodsServiceGetByGoodsIdsRequest) GetMethodName() string {
	return "getByGoodsIds"
}

func New(config *vip.VipBaseConfig) *UnionGoodsServiceGetByGoodsIdsRequest {
	request := UnionGoodsServiceGetByGoodsIdsRequest{
		Param: &UnionGoodsServiceGetByGoodsIdsParam{},
	}
	request.SetConfig(config)
	request.SetClient(vip.DefaultVipApiClient)

	return &request
}

func (v *UnionGoodsServiceGetByGoodsIdsRequest) Execute() (*getByGoodsIdsResponse.UnionGoodsServiceGetByGoodsIdsResponse, error) {
	responseJson, err := v.GetClient().Request(v, false)
	if err != nil {
		return nil, err
	}
	NewResponseObj := &getByGoodsIdsResponse.UnionGoodsServiceGetByGoodsIdsResponse{}
	_ = json.Unmarshal([]byte(responseJson), &NewResponseObj)

	return NewResponseObj, err
}

func (v *UnionGoodsServiceGetByGoodsIdsRequest) GetParamsObject() interface{} {
	return v.Param
}

func (v *UnionGoodsServiceGetByGoodsIdsRequest) GetParams() *UnionGoodsServiceGetByGoodsIdsParam {
	return v.Param
}

type UnionGoodsServiceGetByGoodsIdsParam struct {
	GoodsIdList                []string                                                         `json:"goodsIdList"`
	RequestId                  string                                                           `json:"requestId"` //请求id：UUID
	ChanTag                    string                                                           `json:"chanTag"`   //渠道标识
	GoodsInfoByGoodsIdsRequest vipAdpApiOpenServiceUnionGoodsService.GoodsInfoByGoodsIdsRequest `json:"request"`
}

package meituanCardResquest

import (
	"encoding/json"
	"github.com/liu8534584/mall-sdk/meituan/meituanCard"
	"github.com/liu8534584/mall-sdk/meituan/meituanCard/meituanCardResponse"
)

type QueryCouponRequest struct {
	param *QueryCouponParams
	meituanCard.DdBaseRequest
}

func NewQueryCoupon(config *meituanCard.MeituanCardConfig, param QueryCouponParams) *QueryCouponRequest {
	request := QueryCouponRequest{
		param: &param,
	}
	request.SetConfig(config)
	request.SetClient(meituanCard.DefaultMeituanCardClient)

	return &request
}

func (m *QueryCouponRequest) GetApiParamsObject() interface{} {
	return m.param
}

func (m *QueryCouponRequest) GetParams() *QueryCouponParams {
	return m.param
}

func (m *QueryCouponRequest) GetApiPath() string {
	return "/cps_open/common/api/v1/query_coupon"
}

func (m *QueryCouponRequest) IsPost() bool {
	return true
}

func (m *QueryCouponRequest) Execute() (*meituanCardResponse.QueryCouponResponse, error) {
	execute, err := m.GetClient().Execute(m)
	if err != nil {
		return nil, err
	}
	var res meituanCardResponse.QueryCouponResponse
	_ = json.Unmarshal([]byte(execute), &res)
	//if res.Code != 200 {
	//	var errorRes meituanDdResponse.DdErrorResponse
	//	_ = json.Unmarshal([]byte(execute), &errorRes)
	//	return nil, errors.New(errorRes.Msg)
	//}
	//
	return &res, err
}

type QueryCouponParams struct {
	Longitude    int      `json:"longitude,omitempty"` //经度的百万倍
	Latitude     int      `json:"latitude,omitempty"`  //维度的百万倍
	PriceCap     int      `json:"priceCap,omitempty"`
	PriceFloor   int      `json:"priceFloor,omitempty"`
	VpSkuViewIds []string `json:"vpSkuViewIds,omitempty"`
	ListTopiId   int      `json:"listTopiId,omitempty"`
	PageSize     int      `json:"pageSize,omitempty"`
	PageNo       int      `json:"pageNo,omitempty"`
	SortField    int      `json:"sortField,omitempty"`    //1 售价  2 销量 3 佣金
	AscDescOrder int      `json:"ascDescOrder,omitempty"` //	1升序  2 降序
}

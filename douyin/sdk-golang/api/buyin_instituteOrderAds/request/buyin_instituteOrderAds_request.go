package buyin_instituteOrderAds_request

import (
	buyin_kolOrderAds_response "doudian.com/open/sdk_golang/api/buyin_kolOrderAds/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type BuyinInstituteOrderAdsRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *BuyinInstituteOrderAdsParam
}

func (c *BuyinInstituteOrderAdsRequest) GetUrlPath() string {
	return "/buyin/instituteOrderAds"
}

func New() *BuyinInstituteOrderAdsRequest {
	request := &BuyinInstituteOrderAdsRequest{
		Param: &BuyinInstituteOrderAdsParam{},
	}
	request.SetConfig(doudian_sdk.GlobalLiveConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}

func (c *BuyinInstituteOrderAdsRequest) Execute(accessToken *doudian_sdk.AccessToken) (*buyin_kolOrderAds_response.BuyinKolOrderAdsResponse, error) {
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &buyin_kolOrderAds_response.BuyinKolOrderAdsResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}

func (c *BuyinInstituteOrderAdsRequest) GetParamObject() interface{} {
	return c.Param
}

func (c *BuyinInstituteOrderAdsRequest) GetParams() *BuyinInstituteOrderAdsParam {
	return c.Param
}

type BuyinInstituteOrderAdsParam struct {
	// 每页订单数目，取值区间： [1, 200]
	Size int64 `json:"size,omitempty"`
	// 下一页索引（第一页传“0”）
	Cursor string `json:"cursor,omitempty"`
	// 支付时间开始，最早支持90天前
	StartTime string `json:"start_time,omitempty"`
	// 支付时间结束
	EndTime string `json:"end_time,omitempty"`
	// 订单号。多个订单号用英文 , 分隔，最多支持10个订单号
	OrderIds string `json:"order_ids,omitempty"`
	// 机构PID
	Pid string `json:"pid,omitempty"`
	// 分销类型。直播间:Live；商品：ProductDetail
	DistributionType string `json:"distribution_type,omitempty"`
	// 用户抖店open_id
	OpenId string `json:"open_id,omitempty"`
	// 查询时间类型。pay: 支付时间（默认）; update：联盟侧更新时间，非订单状态更新时间
	TimeType string `json:"time_type,omitempty"`
}

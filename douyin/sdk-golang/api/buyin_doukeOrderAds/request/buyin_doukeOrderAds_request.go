package buyin_doukeOrderAds_request

import (
	"doudian.com/open/sdk_golang/api/buyin_doukeOrderAds/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type BuyinDoukeOrderAdsRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *BuyinDoukeOrderAdsParam
}

func (c *BuyinDoukeOrderAdsRequest) GetUrlPath() string {
	return "/buyin/doukeOrderAds"
}

func New() *BuyinDoukeOrderAdsRequest {
	request := &BuyinDoukeOrderAdsRequest{
		Param: &BuyinDoukeOrderAdsParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}

func (c *BuyinDoukeOrderAdsRequest) Execute(accessToken *doudian_sdk.AccessToken) (*buyin_doukeOrderAds_response.BuyinDoukeOrderAdsResponse, error) {
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &buyin_doukeOrderAds_response.BuyinDoukeOrderAdsResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}

func (c *BuyinDoukeOrderAdsRequest) GetParamObject() interface{} {
	return c.Param
}

func (c *BuyinDoukeOrderAdsRequest) GetParams() *BuyinDoukeOrderAdsParam {
	return c.Param
}

type BuyinDoukeOrderAdsParam struct {
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
	// 抖客PID
	Pid string `json:"pid,omitempty"`
	// 分销类型。Live-直播间，ProductDetail-商品详情，Activity-活动（百亿补贴/秒杀）Mix-H5自建活动页
	DistributionType string `json:"distribution_type,omitempty"`
	// 查询时间类型。pay: 支付时间（默认）; update：联盟侧更新时间，非订单状态更新时间
	TimeType string `json:"time_type,omitempty"`
	// 不填或者填0：查询分销订单，1：查询比价订单
	QueryOrderType int64 `json:"query_order_type,omitempty"`
}

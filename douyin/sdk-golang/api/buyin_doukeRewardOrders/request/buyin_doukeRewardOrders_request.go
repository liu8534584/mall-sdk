package buyin_doukeRewardOrders_request

import (
	"doudian.com/open/sdk_golang/api/buyin_doukeRewardOrders/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type BuyinDoukeRewardOrdersRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *BuyinDoukeRewardOrdersParam
}

func (c *BuyinDoukeRewardOrdersRequest) GetUrlPath() string {
	return "/buyin/doukeRewardOrders"
}

func New() *BuyinDoukeRewardOrdersRequest {
	request := &BuyinDoukeRewardOrdersRequest{
		Param: &BuyinDoukeRewardOrdersParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}

func (c *BuyinDoukeRewardOrdersRequest) Execute(accessToken *doudian_sdk.AccessToken) (*buyin_doukeRewardOrders_response.BuyinDoukeRewardOrdersResponse, error) {
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &buyin_doukeRewardOrders_response.BuyinDoukeRewardOrdersResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}

func (c *BuyinDoukeRewardOrdersRequest) GetParamObject() interface{} {
	return c.Param
}

func (c *BuyinDoukeRewardOrdersRequest) GetParams() *BuyinDoukeRewardOrdersParam {
	return c.Param
}

type BuyinDoukeRewardOrdersParam struct {
	// 页数，从1开始
	Page *int64 `json:"page,omitempty"`
	// 每页订单数目，最大为20
	PageSize *int64 `json:"page_size,omitempty"`
	// 独立抖客PID
	Pid string `json:"pid,omitempty"`
	// 查询开始日期，格式：20230308，奖励订单只保存近30天的订单
	StartDate string `json:"start_date,omitempty"`
	// 查询结束日期，格式：20230308，奖励订单只保存近30天的订单
	EndDate string `json:"end_date,omitempty"`
	// 默认全部，0:全部 1:商品分销 2: 直播间分销 3: mix自建活动页分销 4:频道页分销 99:其他
	DistributionType int64 `json:"distribution_type,omitempty"`
}

package buyin_distributionLiveProductList_request

import (
	"doudian.com/open/sdk_golang/api/buyin_distributionLiveProductList/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type BuyinDistributionLiveProductListRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *BuyinDistributionLiveProductListParam
}

func (c *BuyinDistributionLiveProductListRequest) GetUrlPath() string {
	return "/buyin/distributionLiveProductList"
}

func New() *BuyinDistributionLiveProductListRequest {
	request := &BuyinDistributionLiveProductListRequest{
		Param: &BuyinDistributionLiveProductListParam{},
	}
	request.SetConfig(doudian_sdk.GlobalLiveConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}

func (c *BuyinDistributionLiveProductListRequest) Execute(accessToken *doudian_sdk.AccessToken) (*buyin_distributionLiveProductList_response.BuyinDistributionLiveProductListResponse, error) {
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &buyin_distributionLiveProductList_response.BuyinDistributionLiveProductListResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}

func (c *BuyinDistributionLiveProductListRequest) GetParamObject() interface{} {
	return c.Param
}

func (c *BuyinDistributionLiveProductListRequest) GetParams() *BuyinDistributionLiveProductListParam {
	return c.Param
}

type BuyinDistributionLiveProductListParam struct {
	// 主播的百应ID
	AuthorBuyinId string `json:"author_buyin_id"`
}

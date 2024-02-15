package buyin_kolMaterialsProductsDetails_request

import (
	"doudian.com/open/sdk_golang/api/buyin_kolMaterialsProductsDetails/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type BuyinKolMaterialsProductsDetailsRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *BuyinKolMaterialsProductsDetailsParam
}

func (c *BuyinKolMaterialsProductsDetailsRequest) GetUrlPath() string {
	return "/buyin/kolMaterialsProductsDetails"
}

func New() *BuyinKolMaterialsProductsDetailsRequest {
	request := &BuyinKolMaterialsProductsDetailsRequest{
		Param: &BuyinKolMaterialsProductsDetailsParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}

func (c *BuyinKolMaterialsProductsDetailsRequest) Execute(accessToken *doudian_sdk.AccessToken) (*buyin_kolMaterialsProductsDetails_response.BuyinKolMaterialsProductsDetailsResponse, error) {
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &buyin_kolMaterialsProductsDetails_response.BuyinKolMaterialsProductsDetailsResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}

func (c *BuyinKolMaterialsProductsDetailsRequest) GetParamObject() interface{} {
	return c.Param
}

func (c *BuyinKolMaterialsProductsDetailsRequest) GetParams() *BuyinKolMaterialsProductsDetailsParam {
	return c.Param
}

type BuyinKolMaterialsProductsDetailsParam struct {
	// 批量商品ID（最多20个)
	ProductIds []int64 `json:"product_ids"`
	// 返回是否可分销状态 sharable 字段
	WithShareStatus bool `json:"with_share_status"`
}

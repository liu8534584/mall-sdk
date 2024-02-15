package alliance_materialsProductCategory_request

import (
	"doudian.com/open/sdk_golang/api/alliance_materialsProductCategory/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type AllianceMaterialsProductCategoryRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *AllianceMaterialsProductCategoryParam
}

func (c *AllianceMaterialsProductCategoryRequest) GetUrlPath() string {
	return "/alliance/materialsProductCategory"
}

func New() *AllianceMaterialsProductCategoryRequest {
	request := &AllianceMaterialsProductCategoryRequest{
		Param: &AllianceMaterialsProductCategoryParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}

func (c *AllianceMaterialsProductCategoryRequest) Execute(accessToken *doudian_sdk.AccessToken) (*alliance_materialsProductCategory_response.AllianceMaterialsProductCategoryResponse, error) {
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &alliance_materialsProductCategory_response.AllianceMaterialsProductCategoryResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}

func (c *AllianceMaterialsProductCategoryRequest) GetParamObject() interface{} {
	return c.Param
}

func (c *AllianceMaterialsProductCategoryRequest) GetParams() *AllianceMaterialsProductCategoryParam {
	return c.Param
}

type AllianceMaterialsProductCategoryParam struct {
	// 父类目 （0表示查询一级类目列表）
	ParentId int64 `json:"parent_id"`
}

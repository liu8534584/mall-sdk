package buyin_liveShareMaterial_request

import (
	"doudian.com/open/sdk_golang/api/buyin_liveShareMaterial/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type BuyinLiveShareMaterialRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *BuyinLiveShareMaterialParam
}

func (c *BuyinLiveShareMaterialRequest) GetUrlPath() string {
	return "/buyin/liveShareMaterial"
}

func New() *BuyinLiveShareMaterialRequest {
	request := &BuyinLiveShareMaterialRequest{
		Param: &BuyinLiveShareMaterialParam{},
	}
	request.SetConfig(doudian_sdk.GlobalLiveConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}

func (c *BuyinLiveShareMaterialRequest) Execute(accessToken *doudian_sdk.AccessToken) (*buyin_liveShareMaterial_response.BuyinLiveShareMaterialResponse, error) {
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &buyin_liveShareMaterial_response.BuyinLiveShareMaterialResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}

func (c *BuyinLiveShareMaterialRequest) GetParamObject() interface{} {
	return c.Param
}

func (c *BuyinLiveShareMaterialRequest) GetParams() *BuyinLiveShareMaterialParam {
	return c.Param
}

type BuyinLiveShareMaterialParam struct {
	// 1 品牌 2 达人(默认值)
	AuthorType int32 `json:"author_type"`
	// 作者电商等级，0-7级
	AuthorLevels []int32 `json:"author_levels"`
	// 商品行业类目，列表类型；使用【/alliance/activityProductCategoryList】接口获取
	FristCids []string `json:"frist_cids"`
	// 达人昵称或者账号
	AuthorInfo string `json:"author_info"`
	// 分页，1开始
	Page int64 `json:"page"`
	// 分页大小
	PageSize int64 `json:"page_size"`
	// 排序字段: 1-综合；2-销量；3-佣金率；4-粉丝数
	SortBy int32 `json:"sort_by"`
	// 排序方式：0-降序；1-升序
	SortType int32 `json:"sort_type"`
	// 开关播状态：0:所有，1:开播，2:关播
	LiveStatus int64 `json:"live_status"`
}

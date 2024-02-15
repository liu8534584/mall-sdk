package meituanDdResquest

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/liu8534584/mall-sdk/meituan/meituanDd"
	"github.com/liu8534584/mall-sdk/meituan/meituanDd/meituanDdResponse"
	"github.com/liu8534584/mall-sdk/utils"
	"github.com/spf13/cast"
	"time"
)

type DdSearchRequest struct {
	param *DdSearchParams
	Url   string
	meituanDd.DdBaseRequest
}

func NewDdSearchRequest(config *meituanDd.MeituanDdConfig, param DdSearchParams) *DdSearchRequest {
	param.RequestId = fmt.Sprintf("hs_%d", time.Now().Unix())
	param.UtmSource = config.UtmSource
	param.Version = "1.0"
	param.Timestamp = time.Now().Unix()
	param.PlatformId = config.PlatformId
	param.PromotionId = config.PromotionId
	request := DdSearchRequest{param: &param}
	request.param.AccessToken = request.GetClient().GetAccessToken(config.AppKey, config.UtmSource, request.param.Timestamp)
	getParams := map[string]string{
		"requestId":   param.RequestId,
		"utmSource":   param.UtmSource,
		"version":     param.Version,
		"accessToken": request.param.AccessToken,
		"timestamp":   cast.ToString(param.Timestamp),
	}
	if param.UtmMedium != "" {
		param.UtmMedium = request.GetClient().GetUtmMedium(config.AppKey, param.UtmMedium)
	}
	request.Url = utils.SetQuery(meituanDd.DdBaseUrl+"/api/search/deals", getParams)
	request.SetClient(meituanDd.DefaultMeituanDdClient)
	request.SetConfig(config)
	return &request
}

func (m *DdSearchRequest) GetApiParamsObject() interface{} {
	return m.param
}

func (m *DdSearchRequest) GetParams() *DdSearchParams {
	return m.param
}

func (m *DdSearchRequest) GetApiPath() string {
	return m.Url
}

func (m *DdSearchRequest) IsPost() bool {
	return true
}

func (m *DdSearchRequest) Execute() (*meituanDdResponse.DdSearchResponse, error) {
	execute, err := m.GetClient().Execute(m)
	if err != nil {
		return nil, err
	}
	var res meituanDdResponse.DdSearchResponse
	_ = json.Unmarshal([]byte(execute), &res)
	if res.Code != 200 {
		var errorRes meituanDdResponse.DdErrorResponse
		_ = json.Unmarshal([]byte(execute), &errorRes)
		return nil, errors.New(errorRes.Msg)
	}

	return &res, err
}

type DdSearchParams struct {
	RequestId   string               `json:"requestId,omitempty"`
	UtmSource   string               `json:"utmSource,omitempty"`
	Version     string               `json:"version,omitempty"`
	AccessToken string               `json:"accessToken,omitempty"`
	UtmMedium   string               `json:"utmMedium,omitempty"`
	PromotionId string               `json:"promotionId,omitempty"`
	Timestamp   int64                `json:"timestamp,omitempty"`
	PlatformId  int64                `json:"platformId,omitempty"`
	Cat0Id      int64                `json:"cat0Id,omitempty"`
	Cat1Id      int64                `json:"cat1Id,omitempty"`
	Cat1Ids     []int64              `json:"cat1Ids,omitempty"`
	ShopIds     []string             `json:"shopIds,omitempty"`
	DealIds     []int64              `json:"dealIds,omitempty"`
	DealType    int64                `json:"dealType,omitempty"`
	SortType    int64                `json:"sortType,omitempty"`
	Filters     []ApiDealRangeFilter `json:"filters,omitempty"`
	KeyWords    string               `json:"keyWords,omitempty"`
	Geo         ApiGeoInfo           `json:"geo,omitempty"`
	Page        int64                `json:"page,omitempty"`
	Size        int64                `json:"size,omitempty"`
}

type ApiDealRangeFilter struct {
	FilterType int64 `json:"filterType"`
	Upper      int   `json:"upper"`
	Lower      int   `json:"lower"`
}

type ApiGeoInfo struct {
	CityId     int64   `json:"cityId,omitempty"`
	RegionId   int64   `json:"regionId,omitempty"`
	DistrictId int64   `json:"districtId,omitempty"`
	Lng        float64 `json:"lng,omitempty"`
	Lat        float64 `json:"lat,omitempty"`
	Radii      int64   `json:"radii,omitempty"`
}

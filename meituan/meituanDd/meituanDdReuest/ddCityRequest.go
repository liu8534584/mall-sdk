package meituanDdResquest

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/liu8534584/mall-sdk/meituan/meituanDd"
	"github.com/liu8534584/mall-sdk/meituan/meituanDd/meituanDdResponse"
	"time"
)

type DdCityRequest struct {
	param *DdCityParams
	meituanDd.DdBaseRequest
	Url string
}

func NewDdCityRequest(config *meituanDd.MeituanDdConfig, provinceId string) *DdCityRequest {
	request := DdCityRequest{param: &DdCityParams{
		RequestId:   fmt.Sprintf("hs_%d", time.Now().Unix()),
		UtmSource:   config.UtmSource,
		Version:     "1.0",
		PlatformId:  config.PlatformId,
		Timestamp:   time.Now().Unix(),
		PromotionId: config.PromotionId,
	}}
	request.Url = fmt.Sprintf(meituanDd.DdBaseUrl+"/api/province/%s/cities", provinceId)
	request.SetClient(meituanDd.DefaultMeituanDdClient)
	request.param.AccessToken = request.GetClient().GetAccessToken(config.AppKey, config.UtmSource, request.param.Timestamp)
	request.SetConfig(config)
	return &request
}

func (m *DdCityRequest) GetApiParamsObject() interface{} {
	return m.param
}

func (m *DdCityRequest) GetParams() *DdCityParams {
	return m.param
}

func (m *DdCityRequest) GetApiPath() string {
	return m.Url
}

func (m *DdCityRequest) IsPost() bool {
	return false
}

func (m *DdCityRequest) Execute() (*meituanDdResponse.DdCityResponse, error) {
	execute, err := m.GetClient().Execute(m)
	if err != nil {
		return nil, err
	}
	var res meituanDdResponse.DdCityResponse
	_ = json.Unmarshal([]byte(execute), &res)

	if res.Code != 200 {
		var errorRes meituanDdResponse.DdErrorResponse
		_ = json.Unmarshal([]byte(execute), &errorRes)
		return nil, errors.New(errorRes.Msg)
	}
	return &res, err
}

type DdCityParams struct {
	RequestId   string `json:"requestId"`
	UtmSource   string `json:"utmSource"`
	Version     string `json:"version"`
	AccessToken string `json:"accessToken"`
	PlatformId  int64  `json:"platformId"`
	UtmMedium   string `json:"utmMedium"`
	PromotionId string `json:"promotionId"`
	Timestamp   int64  `json:"timestamp"`
}

package meituanDdResquest

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/liu8534584/mall-sdk/meituan/meituanDd"
	"github.com/liu8534584/mall-sdk/meituan/meituanDd/meituanDdResponse"
	"time"
)

type DdRegionsRequest struct {
	param *DdRegionsParams
	Url   string
	meituanDd.DdBaseRequest
}

func NewDdRegionsRequest(config *meituanDd.MeituanDdConfig, cityId string, includeHot bool) *DdRegionsRequest {
	request := DdRegionsRequest{param: &DdRegionsParams{
		RequestId:   fmt.Sprintf("hs_%d", time.Now().Unix()),
		UtmSource:   config.UtmSource,
		Version:     "1.0",
		Timestamp:   time.Now().Unix(),
		IncludeHot:  includeHot,
		PromotionId: config.PromotionId,
		PlatformId:  config.PlatformId,
	}}
	request.Url = fmt.Sprintf(meituanDd.DdBaseUrl+"/api/mt/city/%s/regions", cityId)
	if config.PlatformId == 1 {
		request.Url = fmt.Sprintf(meituanDd.DdBaseUrl+"/api/city/%s/regions", cityId)
	}
	request.SetClient(meituanDd.DefaultMeituanDdClient)
	request.param.AccessToken = request.GetClient().GetAccessToken(config.AppKey, config.UtmSource, request.param.Timestamp)
	request.SetConfig(config)
	return &request
}

func (m *DdRegionsRequest) GetApiParamsObject() interface{} {
	return m.param
}

func (m *DdRegionsRequest) GetParams() *DdRegionsParams {
	return m.param
}

func (m *DdRegionsRequest) GetApiPath() string {
	return m.Url
}

func (m *DdRegionsRequest) IsPost() bool {
	return false
}

func (m *DdRegionsRequest) Execute() (*meituanDdResponse.DdRegionsResponse, error) {
	execute, err := m.GetClient().Execute(m)
	if err != nil {
		return nil, err
	}
	var res meituanDdResponse.DdRegionsResponse
	_ = json.Unmarshal([]byte(execute), &res)

	if res.Code != 200 {
		var errorRes meituanDdResponse.DdErrorResponse
		_ = json.Unmarshal([]byte(execute), &errorRes)
		return nil, errors.New(errorRes.Msg)
	}

	return &res, err
}

type DdRegionsParams struct {
	RequestId   string `json:"requestId"`
	UtmSource   string `json:"utmSource"`
	Version     string `json:"version"`
	AccessToken string `json:"accessToken"`
	UtmMedium   string `json:"utmMedium"`
	PromotionId string `json:"promotionId"`
	PlatformId  int64  `json:"platformId"`
	Timestamp   int64  `json:"timestamp"`
	IncludeHot  bool   `json:"includeHot"`
}

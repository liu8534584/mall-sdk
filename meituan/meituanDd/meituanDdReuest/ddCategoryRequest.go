package meituanDdResquest

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/liu8534584/mall-sdk/meituan/meituanDd"
	"github.com/liu8534584/mall-sdk/meituan/meituanDd/meituanDdResponse"
	"time"
)

type DdCategoryRequest struct {
	param *DdCategoryParams
	meituanDd.DdBaseRequest
	Url string
}

func NewDdCategoryRequest(config *meituanDd.MeituanDdConfig, cityId, cat0Id string) *DdCategoryRequest {
	request := DdCategoryRequest{param: &DdCategoryParams{
		RequestId:   fmt.Sprintf("hs_%d", time.Now().Unix()),
		UtmSource:   config.UtmSource,
		Version:     "1.0",
		PlatformId:  config.PlatformId,
		Timestamp:   time.Now().Unix(),
		Cat0Id:      cat0Id,
		PromotionId: config.PromotionId,
	}}
	request.Url = fmt.Sprintf(meituanDd.DdBaseUrl+"/api/city/%s/categories", cityId)
	request.SetClient(meituanDd.DefaultMeituanDdClient)
	request.param.AccessToken = request.GetClient().GetAccessToken(config.AppKey, config.UtmSource, request.param.Timestamp)
	request.SetConfig(config)
	return &request
}

func (m *DdCategoryRequest) GetApiParamsObject() interface{} {
	return m.param
}

func (m *DdCategoryRequest) GetParams() *DdCategoryParams {
	return m.param
}

func (m *DdCategoryRequest) GetApiPath() string {
	return m.Url
}

func (m *DdCategoryRequest) IsPost() bool {
	return false
}

func (m *DdCategoryRequest) Execute() (*meituanDdResponse.DdCategoryResponse, error) {
	execute, err := m.GetClient().Execute(m)
	if err != nil {
		return nil, err
	}
	var res meituanDdResponse.DdCategoryResponse
	_ = json.Unmarshal([]byte(execute), &res)

	if res.Code != 200 {
		var errorRes meituanDdResponse.DdErrorResponse
		_ = json.Unmarshal([]byte(execute), &errorRes)
		return nil, errors.New(errorRes.Msg)
	}
	return &res, err
}

type DdCategoryParams struct {
	RequestId   string `json:"requestId"`
	UtmSource   string `json:"utmSource"`
	Version     string `json:"version"`
	AccessToken string `json:"accessToken"`
	PlatformId  int64  `json:"platformId"`
	UtmMedium   string `json:"utmMedium"`
	PromotionId string `json:"promotionId"`
	Timestamp   int64  `json:"timestamp"`
	Cat0Id      string `json:"cat0Id"`
}

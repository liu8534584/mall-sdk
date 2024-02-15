package meituanDdResquest

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/liu8534584/mall-sdk/meituan/meituanDd"
	"github.com/liu8534584/mall-sdk/meituan/meituanDd/meituanDdResponse"
	"time"
)

type DdProvinceRequest struct {
	param *DdProvinceParams
	meituanDd.DdBaseRequest
}

func NewDdProvinceRequest(config *meituanDd.MeituanDdConfig) *DdProvinceRequest {
	request := DdProvinceRequest{param: &DdProvinceParams{
		RequestId:   fmt.Sprintf("hs_%d", time.Now().Unix()),
		UtmSource:   config.UtmSource,
		Version:     "1.0",
		PlatformId:  config.PlatformId,
		Timestamp:   time.Now().Unix(),
		PromotionId: config.PromotionId,
	}}
	request.SetClient(meituanDd.DefaultMeituanDdClient)
	request.param.AccessToken = request.GetClient().GetAccessToken(config.AppKey, config.UtmSource, request.param.Timestamp)
	request.SetConfig(config)
	return &request
}

func (m *DdProvinceRequest) GetApiParamsObject() interface{} {
	return m.param
}

func (m *DdProvinceRequest) GetParams() *DdProvinceParams {
	return m.param
}

func (m *DdProvinceRequest) GetApiMethod() string {
	return "POST"
}

func (m *DdProvinceRequest) GetApiPath() string {
	return meituanDd.DdBaseUrl + "/api/province/all"
}

func (m *DdProvinceRequest) IsPost() bool {
	return false
}

func (m *DdProvinceRequest) Execute() (*meituanDdResponse.DdProvinceResponse, error) {
	execute, err := m.GetClient().Execute(m)
	if err != nil {
		return nil, err
	}
	var res meituanDdResponse.DdProvinceResponse
	_ = json.Unmarshal([]byte(execute), &res)
	if res.Code != 200 {
		var errorRes meituanDdResponse.DdErrorResponse
		_ = json.Unmarshal([]byte(execute), &errorRes)
		return nil, errors.New(errorRes.Msg)
	}

	return &res, err
}

type DdProvinceParams struct {
	RequestId   string `json:"requestId"`
	UtmSource   string `json:"utmSource"`
	Version     string `json:"version"`
	AccessToken string `json:"accessToken"`
	PlatformId  int64  `json:"platformId"`
	UtmMedium   string `json:"utmMedium"`
	PromotionId string `json:"promotionId"`
	Timestamp   int64  `json:"timestamp"`
}

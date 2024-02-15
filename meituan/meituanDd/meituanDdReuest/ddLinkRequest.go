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

type DdLinkRequest struct {
	param *DdLinkParams
	Url   string
	meituanDd.DdBaseRequest
}

func NewDdLinkRequest(config *meituanDd.MeituanDdConfig, param DdLinkParams) *DdLinkRequest {
	param.RequestId = fmt.Sprintf("hs_%d", time.Now().Unix())
	param.UtmSource = config.UtmSource
	param.Version = "1.0"
	param.Timestamp = time.Now().Unix()
	param.PlatformId = config.PlatformId
	request := DdLinkRequest{param: &param}
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
	request.Url = utils.SetQuery(meituanDd.DdBaseUrl+"/api/promotion/link", getParams)
	request.SetClient(meituanDd.DefaultMeituanDdClient)
	request.SetConfig(config)
	return &request
}

func (m *DdLinkRequest) GetApiParamsObject() interface{} {
	return m.param
}

func (m *DdLinkRequest) GetParams() *DdLinkParams {
	return m.param
}

func (m *DdLinkRequest) GetApiPath() string {
	return m.Url
}

func (m *DdLinkRequest) IsPost() bool {
	return true
}

func (m *DdLinkRequest) Execute() (*meituanDdResponse.DdLinkResponse, error) {
	execute, err := m.GetClient().Execute(m)
	if err != nil {
		return nil, err
	}
	var res meituanDdResponse.DdLinkResponse
	_ = json.Unmarshal([]byte(execute), &res)
	if res.Code != 200 {
		var errorRes meituanDdResponse.DdErrorResponse
		_ = json.Unmarshal([]byte(execute), &errorRes)
		return nil, errors.New(errorRes.Msg)
	}

	return &res, err
}

type DdLinkParams struct {
	RequestId    string  `json:"requestId"`
	UtmSource    string  `json:"utmSource"`
	Version      string  `json:"version"`
	AccessToken  string  `json:"accessToken"`
	UtmMedium    string  `json:"utmMedium"`
	PromotionId  string  `json:"promotionId"`
	Timestamp    int64   `json:"timestamp"`
	PlatformId   int64   `json:"platformId"`
	Activity     string  `json:"activity"`     //活动ID (从分销联盟后台--【我要推广】中获取)
	PageLevel    int     `json:"pageLevel"`    //	页面分销层级 (详见PageLevelEnum)
	UserLevel    float64 `json:"userLevel"`    //	渠道用户等级，用于处理页面展示佣金信息，一般在二级分销页面需要传递。默认为1，请求值需要在0-1之间，例如:0.8
	DemandQrInfo bool    `json:"demandQrInfo"` //	是否需要返回二维码，默认为false。一般在微信二级分销场景下需要，请渠道结合实际情况使用该字段，返回二维码可能会导致接口响应时间变长，取链失败等情况。
}

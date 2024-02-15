package meituanCardResquest

import (
	"encoding/json"
	"github.com/liu8534584/mall-sdk/meituan/meituanCard"
	"github.com/liu8534584/mall-sdk/meituan/meituanCard/meituanCardResponse"
)

type GetReferralLinkRequest struct {
	param *GetReferralLinkParams
	meituanCard.DdBaseRequest
}

func NewGetReferralLink(config *meituanCard.MeituanCardConfig, param GetReferralLinkParams) *GetReferralLinkRequest {
	request := GetReferralLinkRequest{
		param: &param,
	}
	request.SetConfig(config)
	request.SetClient(meituanCard.DefaultMeituanCardClient)

	return &request
}

func (m *GetReferralLinkRequest) GetApiParamsObject() interface{} {
	return m.param
}

func (m *GetReferralLinkRequest) GetParams() *GetReferralLinkParams {
	return m.param
}

func (m *GetReferralLinkRequest) GetApiPath() string {
	return "/cps_open/common/api/v1/get_referral_link"
}

func (m *GetReferralLinkRequest) IsPost() bool {
	return true
}

func (m *GetReferralLinkRequest) Execute() (*meituanCardResponse.GetReferralLinkResponse, error) {
	execute, err := m.GetClient().Execute(m)
	if err != nil {
		return nil, err
	}
	var res meituanCardResponse.GetReferralLinkResponse
	_ = json.Unmarshal([]byte(execute), &res)

	return &res, err
}

type GetReferralLinkParams struct {
	ActId     string `json:"actId,omitempty"`     //活动物料 id
	SkuViewId string `json:"skuViewId,omitempty"` //商品券 di
	Sid       string `json:"sid,omitempty"`       //跟单参数  64 位限制
	LinkType  int    `json:"linkType,omitempty"`  //1：H5长链接  2：H5短链接  3：小程序唤起链接
	Text      string `json:"text,omitempty"`
}

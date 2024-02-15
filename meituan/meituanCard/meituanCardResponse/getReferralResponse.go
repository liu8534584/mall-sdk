package meituanCardResponse

type GetReferralLinkResponse struct {
	MeituanCardBaseResp
	Data      string `json:"data"`
	SkuViewId string `json:"skuViewId"`
}

package meituanDd

type MeituanDdConfig struct {
	AppKey      string `json:"appKey"`
	UtmSource   string `json:"utmSource"`
	PlatformId  int64  `json:"platformId"`
	PromotionId string `json:"promotionId"`
}

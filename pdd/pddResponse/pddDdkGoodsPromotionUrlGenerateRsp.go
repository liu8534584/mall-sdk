package pddResponse

type PddDdkGoodsPromotionUrlGenerateRsp struct {
	GoodsPromotionUrlGenerateResponse PddDdkGoodsPromotionUrlGenerateInfoRsp `json:"goods_promotion_url_generate_response"`
	ErrorResponse                     PddErrorResponse                       `json:"error_response"`
}

type PddDdkGoodsPromotionUrlGenerateInfoRsp struct {
	GoodsPromotionUrlList []PddDdkGoodsPromotionUrlGenerateUrlInfoRsp `json:"goods_promotion_url_list"`
}

type PddDdkGoodsPromotionUrlGenerateUrlInfoRsp struct {
	MobileShortUrl string                                                  `json:"mobile_short_url"`
	MobileUrl      string                                                  `json:"mobile_url"`
	SchemaUrl      string                                                  `json:"schema_url"`
	ShortUrl       string                                                  `json:"short_url"`
	Url            string                                                  `json:"url"`
	QqAppInfo      PddDdkGoodsPromotionUrlGenerateUrlInfoQqAppInfoResponse `json:"qq_app_info"`
	WeAppInfo      PddDdkGoodsPromotionUrlGenerateUrlInfoWeAppInfoResponse `json:"we_app_info"`
}

type PddDdkGoodsPromotionUrlGenerateUrlInfoQqAppInfoResponse struct {
	AppId             string `json:"app_id"`
	BannerUrl         string `json:"banner_url"`
	Desc              string `json:"desc"`
	PagePath          string `json:"page_path"`
	QqAppIconUrl      string `json:"qq_app_icon_url"`
	SourceDisplayName string `json:"source_display_name"`
	Title             string `json:"title"`
	UserName          string `json:"user_name"`
}

type PddDdkGoodsPromotionUrlGenerateUrlInfoWeAppInfoResponse struct {
	AppId             string `json:"app_id"`
	BannerUrl         string `json:"banner_url"`
	Desc              string `json:"desc"`
	PagePath          string `json:"page_path"`
	SourceDisplayName string `json:"source_display_name"`
	Title             string `json:"title"`
	UserName          string `json:"user_name"`
	WeAppIconUrl      string `json:"we_app_icon_url"`
}

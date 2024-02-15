package pddResponse

type PddDdkRpPromUrlGenerateRsp struct {
	RpPromotionUrlGenerateRResponse PddRpPromotionUrlGenerateResponse `json:"rp_promotion_url_generate_response"`
	ErrorResponse                   PddErrorResponse                  `json:"error_response"`
}

type PddRpPromotionUrlGenerateResponse struct {
	ResourceList []PddRpPromotionUrlGenerateResourceListResponse `json:"resource_list"`
	UrlList      []PddRpPromotionUrlGenerateUrlListResponse      `json:"url_list"`
}

type PddRpPromotionUrlGenerateResourceListResponse struct {
	Desc string `json:"desc"`
	Url  string `json:"url"`
}

type PddRpPromotionUrlGenerateUrlListResponse struct {
	MobileShortUrl           string                                            `json:"mobile_short_url"`
	MobileUrl                string                                            `json:"mobile_url"`
	MultiGroupMobileShortUrl string                                            `json:"multi_group_mobile_short_url"`
	MultiGroupMobileUrl      string                                            `json:"multi_group_mobile_url"`
	MultiGroupShortUrl       string                                            `json:"multi_group_short_url"`
	MultiGroupUrl            string                                            `json:"multi_group_url"`
	SchemaUrl                string                                            `json:"schema_url"`
	ShortUrl                 string                                            `json:"short_url"`
	Url                      string                                            `json:"url"`
	QqAppInfo                PddRpPromotionUrlGenerateUrlListQqAppInfoResponse `json:"qq_app_info"`
	WeAppInfo                PddRpPromotionUrlGenerateUrlListWeAppInfoResponse `json:"we_app_info"`
}

type PddRpPromotionUrlGenerateUrlListQqAppInfoResponse struct {
	AppId             string `json:"app_id"`
	BannerUrl         string `json:"banner_url"`
	Desc              string `json:"desc"`
	PagePath          string `json:"page_path"`
	QqAppIconUrl      string `json:"qq_app_icon_url"`
	SourceDisplayName string `json:"source_display_name"`
	Title             string `json:"title"`
	UserName          string `json:"user_name"`
}

type PddRpPromotionUrlGenerateUrlListWeAppInfoResponse struct {
	AppId             string `json:"app_id"`
	BannerUrl         string `json:"banner_url"`
	Desc              string `json:"desc"`
	PagePath          string `json:"page_path"`
	SourceDisplayName string `json:"source_display_name"`
	Title             string `json:"title"`
	UserName          string `json:"user_name"`
	WeAppIconUrl      string `json:"we_app_icon_url"`
}

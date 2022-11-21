package pddResponse

type PddDdkResourceUrlGenRsp struct {
	ResourceUrlResponse pddDdkResourceUrlGenInfoResponse `json:"resource_url_response"`
	ErrorResponse       PddErrorResponse                 `json:"error_response"`
}

type pddDdkResourceUrlGenInfoResponse struct {
	Sign          string                              `json:"sign"`
	WeAppInfo     pddDdkResourceUrlWeAppInfoResponse  `json:"we_app_info"`
	SingleUrlList pddDdkResourceUrlGenInfoUrlResponse `json:"single_url_list"`
	MultiUrlList  pddDdkResourceUrlGenInfoUrlResponse `json:"multi_url_list"`
}

type pddDdkResourceUrlGenInfoUrlResponse struct {
	ShortUrl string `json:"short_url"`
	Url      string `json:"url"`
}

type pddDdkResourceUrlWeAppInfoResponse struct {
	AppId             string `json:"app_id"`
	BannerUrl         string `json:"banner_url"`
	Desc              string `json:"desc"`
	PagePath          string `json:"page_path"`
	SourceDisplayName string `json:"source_display_name"`
	Title             string `json:"title"`
	UserName          string `json:"user_name"`
	WeAppIconUrl      string `json:"we_app_icon_url"`
}

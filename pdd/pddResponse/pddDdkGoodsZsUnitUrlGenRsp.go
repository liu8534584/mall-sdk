package pddResponse

type PddDdkGoodsZsUnitUrlGenRsp struct {
	GoodsZsUnitUrlGenRsp GoodsZsUnitUrlGenRsp `json:"goods_zs_unit_generate_response"`
	ErrorResponse        PddErrorResponse     `json:"error_response"`
}

type GoodsZsUnitUrlGenRsp struct {
	MobileShortUrl           string `json:"mobile_short_url"`
	MobileUrl                string `json:"mobile_url"`
	MultiGroupMobileShortUrl string `json:"multi_group_mobile_short_url"`
	MultiGroupMobileUrl      string `json:"multi_group_mobile_url"`
	MultiGroupShortUrl       string `json:"multi_group_short_url"`
	MultiGroupUrl            string `json:"multi_group_url"`
	ShortUrl                 string `json:"short_url"`
	Url                      string `json:"url"`
	WeixinShortLink          string `json:"weixin_short_link"`
}

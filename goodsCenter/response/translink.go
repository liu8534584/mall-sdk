package response

import "github.com/liu8534584/mall-sdk/goodsCenter/response/base"

type TransLinkRes struct {
	Code     base.Code     `json:"code"`
	Message  string        `json:"msg"`
	Demotion base.Demotion `json:"demotion"`
	Data     ItemTransLink `json:"data"`
}

type ItemTransLink struct {
	TransLink []ItemTransLinkDetailRes `json:"transLink"`
}

type RedirectUrlRes struct {
	Url             string `json:"url"`
	SourceUrl       string `json:"sourceUrl"`
	ItemId          string `json:"itemId"`
	ShopId          string `json:"shopId"`
	Type            int    `json:"type"`
	ActivityId      string `json:"activityId"`
	IsFeizhuProduct bool   `json:"isFeizhuProduct"`
}

type BathRedirectUrlRes struct {
	Code     base.Code        `json:"code"`
	Message  string           `json:"msg"`
	Demotion base.Demotion    `json:"demotion"`
	Data     []RedirectUrlRes `json:"data"`
}

type BathRedirectUrlV4Res struct {
	Code     base.Code          `json:"code"`
	Message  string             `json:"msg"`
	Demotion base.Demotion      `json:"demotion"`
	Data     []RedirectUrlV4Res `json:"data"`
}

type RedirectUrlV4Res struct {
	SourceUrl        string   `json:"SourceUrl"` //原始链接
	GoodsType        int      `json:"goodsType"`
	ItemId           string   `json:"itemId"`
	ShopId           string   `json:"shopId"`
	ItemUrl          string   `json:"itemUrl"`
	ReplaceUrl       string   `json:"replaceUrl"`
	RealItemUrl      string   `json:"realItemUrl"`
	HasItemId        bool     `json:"hasItemId"`
	Contents         string   `json:"contents"`
	ActivityId       string   `json:"activityId"`
	CouponUrl        string   `json:"couponUrl"`
	DuoId            string   `json:"duoId"`
	HasKl            bool     `json:"hasKl"`
	KlList           []string `json:"klList"`
	BizSceneId       string   `json:"bizSceneId"`
	IsFeizhuProduct  bool     `json:"isFeizhuProduct"`
	TransLinkUrlType int      `json:"transLinkUrlType"`
	IsKzUrl          bool     `json:"isKzUrl"`
	TrueUrl          string   `json:"trueUrl"`
}

type ItemTransLinkDetailRes struct {
	ItemId    string `json:"itemId"`
	OriginUrl string `json:"originUrl"`
	ShortUrl  string `json:"shortUrl"`
	LongUrl   string `json:"longUrl"`
	SchemaUrl string `json:"schemaUrl"`
	Kl        string `json:"kl"`
	MiniUrl   string `json:"miniUrl"`
	MiniAppId string `json:"miniAppId"`
	QrCodeUrl string `json:"qrCodeUrl"`
}

type BathTransLinkRes struct {
	Code     base.Code     `json:"code"`
	Message  string        `json:"msg"`
	Demotion base.Demotion `json:"demotion"`
	Data     BathTransLink `json:"data"`
}

type BathTransLink struct {
	Content             string                `json:"content"`
	HtmlContent         string                `json:"htmlContent"`
	IsSuccess           bool                  `json:"isSuccess"`
	Toast               string                `json:"toast"`
	IsShowPddSuperRed   int                   `json:"isShowPddSuperRed"`
	IsSelectPddSuperRed int                   `json:"isSelectPddSuperRed"`
	IsShowTbSubsidy     int                   `json:"isShowTbSubsidy"`
	IsSelectTbSubsidy   int                   `json:"isSelectTbSubsidy"`
	TransLinkUrlType    int                   `json:"transLinkUrlType"`
	Link                []BathTransLinkDetail `json:"link"`
}

type BathTransLinkDetail struct {
	ItemId      string `json:"itemId"`
	ShopId      string `json:"shopId"`
	ItemUrl     string `json:"itemUrl"`
	ActivityId  string `json:"activityId"`
	TransLink   string `json:"transLink"`
	GoodsSource int    `json:"goodsSource"`
}

type ShareMiniCode struct {
	MiniCode string `json:"miniCode"`
}

type ShareMiniCodeRes struct {
	Code     base.Code     `json:"code"`
	Message  string        `json:"msg"`
	Demotion base.Demotion `json:"demotion"`
	Data     ShareMiniCode `json:"data"`
}

type ItemShareInfoRes struct {
	Code     base.Code     `json:"code"`
	Message  string        `json:"msg"`
	Demotion base.Demotion `json:"demotion"`
	Data     ItemShareInfo `json:"data"`
}

type ItemShareInfo struct {
	TransLinkInfo TransLinkInfo `json:"transLinkInfo"`
	ItemInfo      ItemDetail    `json:"itemInfo"`
}

type TransLinkInfo struct {
	Url         string `json:"url"`
	Kl          string `json:"kl"`
	SuperRedUrl string `json:"superRedUrl"`
	ShortUrl    string `json:"shortUrl"`
}

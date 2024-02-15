package request

import "github.com/liu8534584/mall-sdk/goodsCenter/response"

type ItemTransLinkReq struct {
	Items                []ItemTransLinkDetail `json:"items"`
	AppId                string                `json:"appId"`
	TaobaoPidInfo        TaobaoPidInfo         `json:"taobaoPid"`
	JdChannelId          int                   `json:"jdChannelId"`
	JdPid                string                `json:"jdPid"`
	PddPid               string                `json:"pddPid"`
	PddSid               string                `json:"pddSid"`
	DuoId                uint64                `json:"duoId"`
	CustomParams         string                `json:"customParams"`
	PddCustomParams      string                `json:"pddCustomParams"`
	DouyinPid            string                `json:"douyinPid"`
	GoodsSource          int                   `json:"goodsSource"`
	IsSelectPddSuperRed  int                   `json:"isSelectPddSuperRed"`
	IsSelectTbSubsidy    int                   `json:"isSelectTbSubsidy"`
	BizSceneId           string                `json:"bizSceneId"`
	PromotionType        string                `json:"promotionType"`
	IsGeneratePddAuthUrl bool                  `json:"isGeneratePddAuthUrl"`
	AppType              int                   `json:"appType"`
	AdCode               string                `json:"adCode"`
	Md                   string                `json:"md"`
	League               LeagueAccount         `json:"league"`
}

type LeagueAccount struct {
	AppKey    string `json:"appKey"`
	AppSecret string `json:"appSecret"`
}

type ItemTransLinkDetail struct {
	OriginTbItemId  string `json:"originTbItemId"`
	ItemId          string `json:"itemId"`
	IsFeizhuProduct bool   `json:"isfeizhuProduct"`
	ItemUrl         string `json:"itemUrl"`
	ActivityId      string `json:"activityId"`
	CouponId        string `json:"couponId"`
	CouponUrl       string `json:"couponUrl"`
	BizSceneId      string `json:"bizSceneId"`
	PromotionType   string `json:"promotionType"`
	DuoId           uint64 `json:"duoId"`
	ChannelType     uint64 `json:"channelType"`
	ResourceType    uint64 `json:"resourceType"`
	ShopId          string `json:"shopId"`
	IsPangolin      bool   `json:"isPangolin"`
	MaterialId      string `json:"materialId"`
	MixActivityId   string `json:"mixActivityId"`
	ActivityUrl     string `json:"activityUrl"`
	ActId           string `json:"actId"`
	//IsPinGou        bool   `json:"isPinGou"`
	PddSourceId            string `json:"pddSourceId"`
	PddSearchId            string `json:"pddSearchId"`
	IsTbSpecial            bool   `json:"isTbSpecial"`     //是否为淘宝特价
	JdGiftCouponKey        string `json:"jdGiftCouponKey"` //jd礼金批次号
	VipChannelType         string `json:"vipChannelType"`
	CommissionDiscountRate string `json:"commissionDiscountRate"`
	SubsidyDiscountRate    string `json:"subsidyDiscountRate"`
}

type RedirectUrlReq struct {
	Url  string `json:"url"`
	Type int    `json:"type"`
}
type RedirectUrlV2Req struct {
	Urls []string `json:"urls"`
}

type RedirectUrlV4Req struct {
	Contents string `json:"contents"`
}

type ShareInfoReq struct {
	AppName              string        `json:"appName"`
	GoodsSource          int           `json:"goodsSource"`
	TaobaoPidInfo        TaobaoPidInfo `json:"taobaoPidInfo"`
	PddPidInfo           PddPidInfo    `json:"pddPidInfo"`
	IsGeneratePddAuthUrl bool          `json:"isGeneratePddAuthUrl"`
	ItemId               string        `json:"itemId"`
	ShopId               string        `json:"shopId"`
	CouponUrl            string        `json:"couponUrl"`
	PddSearchId          string        `json:"pddSearchId"`
	CustomParams         string        `json:"customParams"`
	JdChannelId          int           `json:"jdChannelId"`
	JdItemUrl            string        `json:"jdItemUrl"`
	DouyinPid            string        `json:"douyinPid"`
	AdCode               string        `json:"adCode"`
	OriginItemId         string        `json:"originItemId"`
	ZsDuoId              int           `json:"zsDuoId"`
	IsTbSpecial          bool          `json:"isTbSpecial"` //是否为淘宝特价
}

type BathTransLinkReq struct {
	Content                 string        `json:"content"`
	IsSelectPddSuperRed     int           `json:"isSelectPddSuperRed"`
	IsSelectTbSubsidy       int           `json:"isSelectTbSubsidy"`
	DouyinPid               string        `json:"DouyinPid"`
	TaobaoPidInfo           TaobaoPidInfo `json:"taobaoPid"`
	JdPid                   string        `json:"jdPid"`
	TbElmPidInfo            TaobaoPidInfo `json:"tbElmPidInfo"`
	JdChannelId             int           `json:"jdChannelId"`
	PddPid                  string        `json:"pddPid"`
	PddSid                  string        `json:"pddSid"`
	PddUid                  string        `json:"pddUid"`
	PddCustomParams         string        `json:"pddCustomParams"`
	DyCustomParams          string        `json:"dyCustomParams"`
	SnCustomParams          string        `json:"snCustomParams"`
	JdCustomParams          string        `json:"jdCustomParams"`
	VipCustomParams         string        `json:"vipCustomParams"`
	MeituanCustomParams     string        `json:"meituanCustomParams"`
	MeituanCardCustomParams string        `json:"meituanCardCustomParams"`
	ElemCustomParams        string        `json:"elemCustomParams"`
	IsPangolin              bool          `json:"isPangolin"`
	AppType                 int           `json:"appType"`
	TxyxCustomParams        string        `json:"txyxCustomParams"`
	KaolaCustomParams       string        `json:"KaolaCustomParams"`
	TransLinkUrlType        int           `json:"transLinkUrlType"` // 1：好省短链接 2：tb口令 3：tb短链接 4：pdd小程序 5:pdd短链 6:jd短链
	TransLinkActType        int           `json:"transLinkActType"` // 1：pdd超红 2：tb百亿补贴
}

type ShareMiniCodeReq struct {
	CustomParams string `json:"customParams"`
	GoodsSource  int64  `json:"goodsSource"`
	ActId        string `json:"actId"`
	ActivityId   string `json:"activityId"`
}

type BathTransLinkV2Req struct {
	Content                 string                      `json:"content"`
	PatternContent          []response.RedirectUrlV4Res `json:"patternContent"`
	IsSelectPddSuperRed     int                         `json:"isSelectPddSuperRed"`
	IsSelectTbSubsidy       int                         `json:"isSelectTbSubsidy"`
	DouyinPid               string                      `json:"douyinPid"`
	TaobaoPidInfo           TaobaoPidInfo               `json:"taobaoPid"`
	TbElmPidInfo            TaobaoPidInfo               `json:"tbElmPidInfo"`
	JdChannelId             int                         `json:"jdChannelId"`
	JdPid                   string                      `json:"jdPid"`
	PddPid                  string                      `json:"pddPid"`
	PddSid                  string                      `json:"pddSid"`
	PddUid                  string                      `json:"pddUid"`
	PddCustomParams         string                      `json:"pddCustomParams"`
	DyCustomParams          string                      `json:"dyCustomParams"`
	JdCustomParams          string                      `json:"jdCustomParams"`
	VipCustomParams         string                      `json:"vipCustomParams"`
	MeituanCustomParams     string                      `json:"meituanCustomParams"`
	MeituanCardCustomParams string                      `json:"meituanCardCustomParams"`
	ElemCustomParams        string                      `json:"elemCustomParams"`
	SnCustomParams          string                      `json:"snCustomParams"`
	TxyxCustomParams        string                      `json:"txyxCustomParams"`
	KaolaCustomParams       string                      `json:"kaolaCustomParams"`
	IsPangolin              bool                        `json:"isPangolin"`
	AppType                 int                         `json:"appType"`
	TransLinkUrlType        int                         `json:"transLinkUrlType"` // 1：好省短链接 2：tb口令 3：tb短链接 4：pdd小程序 5:pdd短链 6:jd短链
	TransLinkActType        int                         `json:"transLinkActType"` // 1：pdd超红 2：tb百亿补贴
}

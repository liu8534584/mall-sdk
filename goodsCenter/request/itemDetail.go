package request

type ItemDetailParams struct {
	GoodsSource    int64         `json:"goodsSource"`
	AppId          string        `json:"appId,omitempty"`
	TaobaoPidInfo  TaobaoPidInfo `json:"taobaoPid"`
	PddPidInfo     PddPidInfo    `json:"pddPidInfo"`
	CustomParams   string        `json:"customParams"`
	JdChannelId    int           `json:"jdChannelId"`
	OriginTbItemId string        `json:"originTbItemId"`
	ItemId         string        `json:"itemId"`
	ShopId         string        `json:"shopId"`
	ItemUrl        string        `json:"itemUrl,omitempty"`
	LinkParams     string        `json:"linkParams"`
	ActivityId     string        `json:"activityId,omitempty"`
	ExtraParams    string        `json:"extraParams"`
	PddSearchId    string        `json:"pddSearchId"`
	ZsDuoId        string        `json:"zsDuoId"`
	BizSceneId     string        `json:"bizSceneId"`
	NeedMaterial   bool          `json:"needMaterial"`
	PositionInfo   PositionInfo  `json:"positionInfo"` //定位信息
}

type PositionInfo struct {
	Lon      float64 `json:"lon"`
	Lat      float64 `json:"lat"`
	CityCode string  `json:"cityCode"`
	SortType int     `json:"sortType"`
}

type BathItemDetailParams struct {
	AppName         string                   `json:"appName"`
	NeedMaterial    bool                     `json:"needMaterial"`
	GoodsSource     int                      `json:"goodsSource"`
	AppId           string                   `json:"appId,omitempty"`
	TaobaoPidInfo   TaobaoPidInfo            `json:"taobaoPid"`
	JdChannelId     int                      `json:"jdChannelId"`
	CustomParams    string                   `json:"customParams"`
	PddCustomParams string                   `json:"pddCustomParams"`
	PddPid          string                   `json:"pddPid"`
	PddSid          string                   `json:"pddSid"`
	CustomerParam   string                   `json:"customerParam"`
	ItemLists       []ItemSdkProductInfoLReq `json:"itemLists"`
}

type ItemSdkProductInfoLReq struct {
	ItemId  string `json:"itemId"`
	ShopId  string `json:"shopId"`
	ItemUrl string `json:"itemUrl,omitempty"`
}

type BathItemDetailV2Params struct {
	AppName         string                    `json:"appName"`
	NeedMaterial    bool                      `json:"needMaterial"`
	AppId           string                    `json:"appId,omitempty"`
	TaobaoPidInfo   TaobaoPidInfo             `json:"taobaoPid"`
	CustomParams    string                    `json:"customParams"`
	PddPid          string                    `json:"pddPid"`
	PddSid          string                    `json:"pddSid"`
	CustomerParam   string                    `json:"customerParam"`
	ItemLists       []ItemSdkProductInfoV2Req `json:"itemLists"`
	JdChannelId     int                       `json:"jdChannelId"`
	PddCustomParams string                    `json:"pddCustomParams"`
}

type ItemSdkProductInfoV2Req struct {
	ItemId      string `json:"itemId"`
	ShopId      string `json:"shopId"`
	ItemUrl     string `json:"itemUrl,omitempty"`
	GoodsSource int    `json:"goodsSource"`
}

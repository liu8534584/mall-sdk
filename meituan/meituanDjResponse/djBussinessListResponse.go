package meituanDjResponse

type DjBusinessResponse struct {
	RetStatus int64     `json:"retStatus"` //响应状态 0 表示正常，1 表示错误
	Message   string    `json:"message"`   //响应注释说明
	PoiList   []PoiList `json:"poiList"`
	PagePvId  string    `json:"pagePvId"`
}

type PoiList struct {
	WmPoiId          string        `json:"wmPoiId"`          // 商家 id
	Name             string        `json:"name"`             // 商家名称
	Picture          string        `json:"picture"`          // 商家图标，图片比例为 1:1 或 4:3
	WmPoiScore       string        `json:"wmPoiScore"`       // 商家评分
	MonthSaleNumStr  string        `json:"monthSaleNumStr"`  // 商家销量
	MinPriceTip      string        `json:"minPriceTip"`      // 起送价格
	ShippingFeeTip   string        `json:"shippingFeeTip"`   // 折扣后配送费
	AveragePrice     string        `json:"averagePrice"`     // 人均消费价格
	DeliveryTimeTip  string        `json:"deliveryTimeTip"`  // 配送时长
	DeliveryDistance string        `json:"deliveryDistance"` // 配送距离
	PoiRankLabel     string        `json:"poiRankLabel"`     // 榜单排名(例如:高新区奶茶排行榜第 1 名)
	DeliveryType     string        `json:"deliveryType"`     // 配送方式（当商家支持自取时，该字段有值，值为：支持自取）
	Ratio            string        `json:"ratio"`            // i32 佣金比例 (例如:46 表示 46/10000)
	SkuList          []SkuList     `json:"skuList"`          // 商家的菜品列表（最多 5 个菜品，有可能少于 5 个）
	ActionUrl        UserActionUrl `json:"actionUrl"`        //  用户行为链接(菜品维度)。注商品维度目前没有曝光和点击回 调链接。
}
type SkuList struct {
	Name         string        `json:"name"`         // 菜品 名称
	SkuId        string        `json:"skuId"`        // 菜品 id
	PicUrl       string        `json:"picUrl"`       // 菜品图
	ActPrice     string        `json:"actPrice"`     // 菜品售价
	OriPrice     string        `json:"oriPrice"`     // 菜品原价
	MonthSaleNum string        `json:"monthSaleNum"` // 菜品月销量
	ActionUrl    UserActionUrl `json:"actionUrl"`    //  用户行为链接(菜品维度)。注商品维度目前没有曝光和点击回 调链接。
}

type UserActionUrl struct {
	DpUrl           string `json:"dpUrl"`           //调起链接（跳转到 app 内的点餐页）
	LpUrl           string `json:"lpUrl"`           //调起链接（跳转到 h5 点餐页）
	WxPath          string `json:"wxPath"`          //微信小程序调起路径（到点餐页）
	WxAppid         string `json:"wxAppid"`         //微信小程序调起参数 appid
	ImpMonitorUrl   string `json:"impMonitorUrl"`   //曝光监控链接
	ClickMonitorUrl string `json:"clickMonitorUrl"` //点击监控链接
	WxAppOrgId      string `json:"wxAppOrgid"`      //微信小程序原始 id
}

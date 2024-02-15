package kaolaResponse

type QueryGoodsInfoResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data []GoodsInfo `json:"data"`
}

type GoodsInfo struct {
	ActivityInfo struct {
		ActivityLable string `json:"activityLable"`
		NoPostage     int    `json:"noPostage"`
	} `json:"activityInfo"`
	BaseInfo struct {
		BizType          int      `json:"bizType"`
		BrandCountryName string   `json:"brandCountryName"`
		BrandName        string   `json:"brandName"`
		DetailImgList    []string `json:"detailImgList"`
		GoodsSubTitle    string   `json:"goodsSubTitle"`
		GoodsTitle       string   `json:"goodsTitle"`
		ImageList        []string `json:"imageList"`
		ImportType       int      `json:"importType"`
		InterPurch       int      `json:"interPurch"`
		OnlineStatus     int      `json:"onlineStatus"`
		Self             int      `json:"self"`
		Store            int      `json:"store"`
	} `json:"baseInfo"`
	CategoryInfo []struct {
		CategoryId   int    `json:"categoryId"`
		CategoryName string `json:"categoryName"`
		Level        int    `json:"level"`
	} `json:"categoryInfo"`
	CommissionInfo struct {
		CommissionRate         float64 `json:"commissionRate"`
		GroupBuyCommissionRate float64 `json:"groupBuyCommissionRate"`
		ExpireType             int     `json:"expireType"`
	} `json:"commissionInfo"`
	DepositInfo struct {
		Deposit          bool `json:"deposit"`
		DepositEndTime   int  `json:"depositEndTime"`
		DepositStartTime int  `json:"depositStartTime"`
		PayEndTime       int  `json:"payEndTime"`
		PayStartTime     int  `json:"payStartTime"`
	} `json:"depositInfo"`
	GoodsId  int `json:"goodsId"`
	LinkInfo struct {
		GoodsDetailUrl   string `json:"goodsDetailUrl"`
		GroupBuyShareUrl string `json:"groupBuyShareUrl"`
		GoodsPCUrl       string `json:"goodsPCUrl"`
		MiniShareUrl     string `json:"miniShareUrl"`
		ShareUrl         string `json:"shareUrl"`
		ShortShareUrl    string `json:"shortShareUrl"`
	} `json:"linkInfo"`
	PriceInfo struct {
		CrossVipDiscount   bool    `json:"crossVipDiscount"`
		CurrentPrice       float64 `json:"currentPrice"`
		GroupBuyPrice      float64 `json:"groupBuyPrice"`
		MarketPrice        float64 `json:"marketPrice"`
		MemberCurrentPrice float64 `json:"memberCurrentPrice"`
		MemberPriceSpread  float64 `json:"memberPriceSpread"`
	} `json:"priceInfo"`
	PromotionInfo []struct {
		ActivityGoodsSellCount int     `json:"activityGoodsSellCount"`
		ActivityGoodsStore     int     `json:"activityGoodsStore"`
		ApplyUserType          int     `json:"applyUserType"`
		EndTime                string  `json:"endTime"`
		GoodsId                int     `json:"goodsId"`
		PromotionSalePrice     float64 `json:"promotionSalePrice"`
		SkuId                  string  `json:"skuId"`
		StartTime              string  `json:"startTime"`
		StoreLimit             bool    `json:"storeLimit"`
	} `json:"promotionInfo"`
}

package meituanDdResponse

type DdSearchResponse struct {
	Code int64    `json:"code"`
	Msg  PageInfo `json:"msg"`
}

type PageInfo struct {
	RecordCount int64             `json:"recordCount"`
	PageSize    int64             `json:"pageSize"`
	Page        int64             `json:"page"`
	Records     []ApiSearchDealVO `json:"records"`
}

type ApiSearchDealVO struct {
	DealBaseInfo ApiDealBaseInfoVO  `json:"dealBaseInfo"`
	DealDetail   ApiDealDetailVO    `json:"dealDetail"`
	ShopInfo     DealShopBaseInfoVO `json:"shopInfo"`
	CouponInfo   ApiCouponInfoVO    `json:"couponInfo"`
}
type ApiDealBaseInfoVO struct {
	DealGroupId     int64  `json:"dealGroupId"`
	DealTitle       string `json:"dealTitle"`
	DealDescription string `json:"dealDescription"`
	DefaultPic      string `json:"defaultPic"`
	RealPrice       int64  `json:"realPrice"`
	FinalPrice      int64  `json:"finalPrice"`
	MarketPrice     int64  `json:"marketPrice"`
	DealType        int64  `json:"dealType"`
}
type ApiDealDetailVO struct {
	MenuType      int64  `json:"menuType"`
	Menu          string `json:"menu"`
	PositiveRatio string `json:"positiveRatio"`
	FeedbackTag   string `json:"feedbackTag"`
}
type DealShopBaseInfoVO struct {
	ShopId       int64  `json:"shopId"`
	ShopUuid     string `json:"shopUuid"`
	ShopName     string `json:"shopName"`
	BranchName   string `json:"branchName"`
	BrandName    string `json:"brandName"`
	CateName     string `json:"cateName"`
	RegionName   string `json:"regionName"`
	Distance     int64  `json:"distance"`
	ShopPower    int64  `json:"shopPower"`
	Address      string `json:"address"`
	CityName     string `json:"cityName"`
	BizHourDesc  string `json:"bizHourDesc"`
	IsMustEat    bool   `json:"isMustEat"`
	IsBlackPearl bool   `json:"isBlackPearl"`
	PricePerson  int64  `json:"pricePerson"`
	GradeIcon    string `json:"gradeIcon"`
}
type ApiCouponInfoVO struct {
	DiscountPrice   int64   `json:"discountPrice"`
	PriceLimit      int64   `json:"priceLimit"`
	CommissionRatio float64 `json:"commissionRatio"`
	ValidStartDate  string  `json:"validStartDate"`
	ValidEndDate    string  `json:"validEndDate"`
	CouponPageUrl   string  `json:"couponPageUrl"`
	MiniProgramUrl  string  `json:"miniProgramUrl"`
}

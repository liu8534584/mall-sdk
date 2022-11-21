package dataoke

type GoodsDetailRes struct {
	Cache     bool        `json:"cache"`
	Code      int         `json:"code"`
	Data      GoodsDetail `json:"data"`
	Msg       string      `json:"msg"`
	RequestID string      `json:"requestId"`
	Time      int64       `json:"time"`
}

type GoodsDetail struct {
	ID                     int           `json:"id"`
	GoodsID                string        `json:"goodsId"`
	ItemLink               string        `json:"itemLink"`
	Title                  string        `json:"title"`
	Dtitle                 string        `json:"dtitle"`
	OriginalPrice          float64       `json:"originalPrice"`
	ActualPrice            float64       `json:"actualPrice"`
	ShopType               int           `json:"shopType"`
	MonthSales             int           `json:"monthSales"`
	TwoHoursSales          int           `json:"twoHoursSales"`
	DailySales             int           `json:"dailySales"`
	CommissionType         int           `json:"commissionType"`
	CommissionRate         int           `json:"commissionRate"`
	Desc                   string        `json:"desc"`
	CouponReceiveNum       int           `json:"couponReceiveNum"`
	CouponTotalNum         int           `json:"couponTotalNum"`
	CouponRemainCount      int           `json:"couponRemainCount"`
	CouponLink             string        `json:"couponLink"`
	CouponID               string        `json:"couponId"`
	CouponEndTime          string        `json:"couponEndTime"`
	CouponStartTime        string        `json:"couponStartTime"`
	CouponPrice            int           `json:"couponPrice"`
	CouponConditions       string        `json:"couponConditions"`
	ActivityType           int           `json:"activityType"`
	ActivityStartTime      string        `json:"activityStartTime"`
	ActivityEndTime        string        `json:"activityEndTime"`
	ShopName               string        `json:"shopName"`
	ShopLevel              int           `json:"shopLevel"`
	DescScore              float64       `json:"descScore"`
	DsrScore               float64       `json:"dsrScore"`
	DsrPercent             float64       `json:"dsrPercent"`
	ShipScore              float64       `json:"shipScore"`
	ShipPercent            float64       `json:"shipPercent"`
	ServiceScore           float64       `json:"serviceScore"`
	ServicePercent         float64       `json:"servicePercent"`
	Brand                  int           `json:"brand"`
	BrandID                int64         `json:"brandId"`
	BrandName              string        `json:"brandName"`
	HotPush                int           `json:"hotPush"`
	TeamName               string        `json:"teamName"`
	QuanMLink              int           `json:"quanMLink"`
	HzQuanOver             int           `json:"hzQuanOver"`
	Yunfeixian             int           `json:"yunfeixian"`
	EstimateAmount         int           `json:"estimateAmount"`
	FreeshipRemoteDistrict int           `json:"freeshipRemoteDistrict"`
	DiscountType           int           `json:"discountType"`
	DiscountFull           int           `json:"discountFull"`
	DiscountCut            int           `json:"discountCut"`
	Discounts              float64       `json:"discounts"`
	MarketGroup            []interface{} `json:"marketGroup"`
	ActivityInfo           []interface{} `json:"activityInfo"`
	InspectedGoods         int           `json:"inspectedGoods"`
	Divisor                int           `json:"divisor"`
	MainPic                string        `json:"mainPic"`
	MarketingMainPic       string        `json:"marketingMainPic"`
	SellerID               string        `json:"sellerId"`
	Cid                    int           `json:"cid"`
	Tbcid                  int           `json:"tbcid"`
	Subcid                 []interface{} `json:"subcid"`
	Haitao                 int           `json:"haitao"`
	Tchaoshi               int           `json:"tchaoshi"`
	Lowest                 int           `json:"lowest"`
	GoldSellers            int           `json:"goldSellers"`
	Video                  string        `json:"video"`
	CreateTime             string        `json:"createTime"`
	SpecialText            []interface{} `json:"specialText"`
	DirectCommissionType   int           `json:"directCommissionType"`
	DirectCommission       int           `json:"directCommission"`
	DirectCommissionLink   string        `json:"directCommissionLink"`
	ActivityID             string        `json:"activityId"`
	Sales24H               int           `json:"sales24h"`
	Imgs                   string        `json:"imgs"`
	Reimgs                 string        `json:"reimgs"`
	BrandWenan             string        `json:"brandWenan"`
	PreviewStartTime       string        `json:"previewStartTime"`
	ZsFullMinusInfo        string        `json:"zsFullMinusInfo"`
	DetailPics             string        `json:"detailPics"`
	IsSubdivision          int           `json:"isSubdivision"`
	SubdivisionID          int           `json:"subdivisionId"`
	SubdivisionName        string        `json:"subdivisionName"`
	SubdivisionRank        int           `json:"subdivisionRank"`
	ShopLogo               string        `json:"shopLogo"`
}

type GoodsDetailImage struct {
	Height      string        `json:"height"`
	HotAreaList []interface{} `json:"hotAreaList"`
	Img         string        `json:"img"`
	Width       string        `json:"width"`
}

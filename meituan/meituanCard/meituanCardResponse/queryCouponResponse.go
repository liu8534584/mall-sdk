package meituanCardResponse

type QueryCouponResponse struct {
	MeituanCardBaseResp
	Data    []MeituanCardCouponDetail `json:"data"`
	HasNext bool                      `json:"hasNext"`
}

type MeituanCardCouponDetail struct {
	AvailablePoiInfo    AvailablePoiInfo    `json:"availablePoiInfo"`
	BrandInfo           BrandInfo           `json:"brandInfo"`
	CommissionInfo      CommissionInfo      `json:"commissionInfo"`
	CouponPackDetail    CouponPackDetail    `json:"couponPackDetail"`
	DeliverablePoiInfo  *DeliverablePoiInfo `json:"deliverablePoiInfo"`
	PurchaseLimitInfo   PurchaseLimitInfo   `json:"purchaseLimitInfo"`
	CouponValidTimeInfo CouponValidTimeInfo `json:"couponValidTimeInfo"`
}

type AvailablePoiInfo struct {
	AvailablePoiNum int `json:"availablePoiNum"`
}

type BrandInfo struct {
	BrandName    *string `json:"brandName"`
	BrandLogoUrl *string `json:"brandLogoUrl"`
}

type CommissionInfo struct {
	CommissionPercent string      `json:"commissionPercent"`
	Commission        interface{} `json:"commission"`
}

type CouponPackDetail struct {
	Name          string `json:"name"`
	SkuViewId     string `json:"skuViewId"`
	CouponNum     int    `json:"couponNum"`
	ValidTime     int    `json:"validTime"`
	HeadUrl       string `json:"headUrl"`
	SaleVolume    string `json:"saleVolume"`
	StartTime     int    `json:"startTime"`
	EndTime       int    `json:"endTime"`
	SaleStatus    bool   `json:"saleStatus"`
	OriginalPrice string `json:"originalPrice"`
	SellPrice     string `json:"sellPrice"`
}

type DeliverablePoiInfo struct {
	PoiName          string `json:"poiName"`
	PoiLogoUrl       string `json:"poiLogoUrl"`
	DeliveryDistance string `json:"deliveryDistance"`
	DistributionCost string `json:"distributionCost"`
	DeliveryDuration string `json:"deliveryDuration"`
	LastDeliveryFee  string `json:"lastDeliveryFee"`
}

type PurchaseLimitInfo struct {
	SingleDayPurchaseLimit int `json:"singleDayPurchaseLimit"`
}

type CouponValidTimeInfo struct {
	CouponValidTimeType int `json:"couponValidTimeType"`
	CouponValidDay      int `json:"couponValidDay"`
	CouponValidSTime    int `json:"couponValidSTime"`
	CouponValidETime    int `json:"couponValidETime"`
}

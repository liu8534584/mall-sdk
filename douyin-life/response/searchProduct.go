package response

type SearchProductResponse struct {
	ApiResponse
	Data SearchProductData `json:"data"`
}

type SearchProductData struct {
	Cursor      int             `json:"cursor"`
	HasMore     bool            `json:"has_more"`
	ProductList []ProductDetail `json:"product_list"`
}
type ProductImg struct {
	URLList []string `json:"UrlList"`
}
type NearestPoi struct {
	Distance float64 `json:"distance"`
	Name     string  `json:"name"`
}
type SaleTags struct {
	Content string `json:"content"`
}
type FirstCategory struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
type SecondCategory struct {
	Name string `json:"name"`
	ID   int    `json:"id"`
}
type ThirdCategory struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
type CategoryInfo struct {
	FirstCategory  FirstCategory  `json:"first_category"`
	SecondCategory SecondCategory `json:"second_category"`
	ThirdCategory  ThirdCategory  `json:"third_category"`
}
type DyDetailPage struct {
	H5URL string `json:"h5_url"`
}
type SoldInfo struct {
	BlurredSoldCount string `json:"blurred_sold_count"`
}
type AvailableTask struct {
	IsRecommend    bool  `json:"is_recommend"`
	Status         int   `json:"status"`
	Type           int   `json:"type"`
	ValidTimeBegin int64 `json:"valid_time_begin"`
	ValidTimeEnd   int64 `json:"valid_time_end"`
	Commission     int   `json:"commission"`
	CommissionRate int   `json:"commission_rate"`
	ID             int64 `json:"id"`
}
type SeckillInfo struct {
	Status         int    `json:"status"`
	StartTime      uint64 `json:"start_time"`
	EndTime        uint64 `json:"end_time"`
	DiscountAmount int64  `json:"discount_amount"`
}

type CouponInfo struct {
	DiscountAmount int64  `json:"discount_amount"`
	UseStartTime   uint64 `json:"use_start_time"`
	UseEndTime     uint64 `json:"use_end_time"`
}

type MarketingInfo struct {
	SeckillInfo *SeckillInfo  `json:"seckill_info"`
	CouponList  []*CouponInfo `json:"coupon_list"`
}
type PriceInfo struct {
	Price       uint64 `json:"price"`
	OriginPrice uint64 `json:"origin_price"`
}
type ItemList struct {
	Unit  string `json:"unit"`
	Count int    `json:"count"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}
type ProductItemDetail struct {
	ItemList []ItemList `json:"item_list"`
	Title    string     `json:"title"`
}
type ProductDetail struct {
	NearestPoi        NearestPoi          `json:"nearest_poi,omitempty"`
	ProductImg        ProductImg          `json:"product_img"`
	SaleTags          []SaleTags          `json:"sale_tags"`
	Status            int                 `json:"status"`
	CategoryInfo      CategoryInfo        `json:"category_info"`
	DyDetailPage      DyDetailPage        `json:"dy_detail_page"`
	Name              string              `json:"name"`
	PoiCount          int                 `json:"poi_count"`
	SoldInfo          SoldInfo            `json:"sold_info"`
	SoldTimeBegin     int                 `json:"sold_time_begin"`
	SoldTimeEnd       int                 `json:"sold_time_end"`
	ID                int64               `json:"id"`
	NeedInvoice       bool                `json:"need_invoice"`
	AvailableTask     []AvailableTask     `json:"available_task"`
	MarketingInfo     MarketingInfo       `json:"marketing_info"`
	PriceInfo         PriceInfo           `json:"price_info"`
	ProductItemDetail []ProductItemDetail `json:"product_item_detail"`
}

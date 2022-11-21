package pddResponse

type PddDdkGoodsRecommendGetRsp struct {
	GoodsBasicDetailResponse PddDdkGoodsRecommendGetGoodsRsp `json:"goods_basic_detail_response"`
	ErrorResponse            PddErrorResponse                `json:"error_response"`
}

type PddDdkGoodsRecommendGetGoodsRsp struct {
	Total    uint64                               `json:"total"`
	SearchId string                               `json:"search_id"`
	listId   string                               `json:"list_id"`
	List     []PddDdkGoodsRecommendGetItemInfoRsp `json:"list"`
}

type PddDdkGoodsRecommendGetItemInfoRsp struct {
	ActivityPromotionRate      uint16   `json:"activity_promotion_rate"`
	ActivityTags               []uint64 `json:"activity_tags"`
	BrandName                  string   `json:"brand_name"`
	CashGiftAmount             uint64   `json:"cash_gift_amount"`
	CatId                      string   `json:"cat_id"`
	CatIds                     []uint64 `json:"cat_ids"`
	CouponDiscount             uint64   `json:"coupon_discount"`
	CouponEndTime              uint64   `json:"coupon_end_time"`
	CouponMinOrderAmount       uint64   `json:"coupon_min_order_amount"`
	CouponPrice                uint64   `json:"coupon_price"`
	CouponRemainQuantity       uint64   `json:"coupon_remain_quantity"`
	CouponStartTime            uint64   `json:"coupon_start_time"`
	CouponTotalQuantity        uint64   `json:"coupon_total_quantity"`
	CreateAt                   uint64   `json:"create_at"`
	DescTxt                    string   `json:"desc_txt"`
	ExtraCouponAmount          uint64   `json:"extra_coupon_amount"`
	GoodsDesc                  string   `json:"goods_desc"`
	GoodsImageUrl              string   `json:"goods_image_url"`
	GoodsLabels                []uint16 `json:"goods_labels"`
	GoodsName                  string   `json:"goods_name"`
	GoodsRate                  uint64   `json:"goods_rate"`
	GoodsSign                  string   `json:"goods_sign"`
	GoodsThumbnailUrl          string   `json:"goods_thumbnail_url"`
	GoodsType                  uint16   `json:"goods_type"`
	HasCoupon                  bool     `json:"has_coupon"`
	HasMaterial                bool     `json:"has_material"`
	LgstTxt                    string   `json:"lgst_txt"`
	MallId                     uint64   `json:"mall_id"`
	MallName                   string   `json:"mall_name"`
	MarketFee                  uint64   `json:"market_fee"`
	MerchantType               string   `json:"merchant_type"`
	MinGroupPrice              uint64   `json:"min_group_price"`
	MinNormalPrice             uint64   `json:"min_normal_price"`
	OptId                      string   `json:"opt_id"`
	OptIds                     []uint64 `json:"opt_ids"`
	OptName                    string   `json:"opt_name"`
	PredictPromotionRate       uint64   `json:"predict_promotion_rate"`
	PromotionRate              uint64   `json:"promotion_rate"`
	QrCodeImageUrl             string   `json:"qr_code_image_url"`
	RealtimeSalesTip           string   `json:"realtime_sales_tip"`
	SalesTip                   string   `json:"sales_tip"`
	SearchId                   string   `json:"search_id"`
	ServTxt                    string   `json:"serv_txt"`
	ShareDesc                  string   `json:"share_desc"`
	ShareRate                  uint16   `json:"share_rate"`
	SubsidyAmount              uint64   `json:"subsidy_amount"`
	SubsidyDuoAmountTenMillion uint16   `json:"subsidy_duo_amount_ten_million"`
	UnifiedTags                []string `json:"unified_tags"`
}

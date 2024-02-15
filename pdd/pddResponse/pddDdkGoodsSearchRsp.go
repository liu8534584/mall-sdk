package pddResponse

type PddDdkGoodsSearchRsp struct {
	GoodsSearchResponse PddDdkGoodsSearchInfoRsp `json:"goods_search_response"`
	ErrorResponse       PddErrorResponse         `json:"error_response"`
}

type PddDdkGoodsSearchInfoRsp struct {
	GoodsList  []PddDdkGoodsSearchGoodsInfoRsp `json:"goods_list"`
	ListId     string                          `json:"list_id"`
	SearchId   string                          `json:"search_id"`
	TotalCount int                             `json:"total_count"`
}

type PddDdkGoodsSearchGoodsInfoRsp struct {
	ActivityPromotionRate       uint16   `json:"activity_promotion_rate"`
	ActivityTags                []uint64 `json:"activity_tags"`
	ActivityType                int      `json:"activity_type"`
	BrandName                   string   `json:"brand_name"`
	CashGiftAmount              uint64   `json:"cash_gift_amount"`
	CltCpnBatchSn               string   `json:"clt_cpn_batch_sn"`
	CltCpnDiscount              uint64   `json:"clt_cpn_discount"`
	CltCpnEndTime               uint64   `json:"clt_cpn_end_time"`
	CltCpnMinAmt                uint64   `json:"clt_cpn_min_amt"`
	CltCpnAuantity              uint64   `json:"clt_cpn_auantity"`
	CltCpnRemainQuantity        uint64   `json:"clt_cpn_remain_quantity"`
	CltCpnStartTime             uint64   `json:"clt_cpn_start_time"`
	MallImgUrl                  string   `json:"mall_img_url"`
	CatIds                      []uint64 `json:"cat_ids"`
	CouponDiscount              uint64   `json:"coupon_discount"`
	CouponEndTime               uint64   `json:"coupon_end_time"`
	CouponMinOrderAmount        uint64   `json:"coupon_min_order_amount"`
	CouponPrice                 uint64   `json:"coupon_price"`
	CouponRemainQuantity        uint64   `json:"coupon_remain_quantity"`
	CouponStartTime             uint64   `json:"coupon_start_time"`
	CouponTotalQuantity         uint64   `json:"coupon_total_quantity"`
	CreateAt                    uint64   `json:"create_at"`
	DescTxt                     string   `json:"desc_txt"`
	ExtraCouponAmount           uint64   `json:"extra_coupon_amount"`
	GoodsDesc                   string   `json:"goods_desc"`
	GoodsImageUrl               string   `json:"goods_image_url"`
	GoodsLabels                 []uint16 `json:"goods_labels"`
	GoodsName                   string   `json:"goods_name"`
	GoodsRate                   uint64   `json:"goods_rate"`
	GoodsSign                   string   `json:"goods_sign"`
	GoodsId                     uint64   `json:"goods_id"`
	GoodsThumbnailUrl           string   `json:"goods_thumbnail_url"`
	GoodsType                   uint16   `json:"goods_type"`
	HasCoupon                   bool     `json:"has_coupon"`
	HasMallCoupon               bool     `json:"has_mall_coupon"`
	HasMaterial                 bool     `json:"has_material"`
	LgstTxt                     string   `json:"lgst_txt"`
	MallCouponDiscountPct       int      `json:"mall_coupon_discount_pct"`
	MallCouponEndTime           uint64   `json:"mall_coupon_end_time"`
	MallCouponMaxDiscountAmount int      `json:"mall_coupon_max_discount_amount"`
	MallCouponMinOrderAmount    int      `json:"mall_coupon_min_order_amount"`
	MallCouponRemainQuantity    uint64   `json:"mall_coupon_remain_quantity"`
	MallCouponStartTime         uint64   `json:"mall_coupon_start_time"`
	MallCouponTotalQuantity     uint64   `json:"mall_coupon_total_quantity"`
	MallCouponId                uint64   `json:"mall_coupon_id"`
	MallCps                     int      `json:"mall_cps"`
	MallId                      uint64   `json:"mall_id"`
	MallName                    string   `json:"mall_name"`
	MarketFee                   uint64   `json:"market_fee"`
	MerchantType                int      `json:"merchant_type"`
	MinGroupPrice               uint64   `json:"min_group_price"`
	MinNormalPrice              uint64   `json:"min_normal_price"`
	OnlySceneAuth               bool     `json:"only_scene_auth"`
	PlanType                    uint64   `json:"plan_type"`
	OptId                       uint64   `json:"opt_id"`
	OptIds                      []uint64 `json:"opt_ids"`
	OptName                     string   `json:"opt_name"`
	PredictPromotionRate        uint64   `json:"predict_promotion_rate"`
	PromotionRate               uint64   `json:"promotion_rate"`
	QrCodeImageUrl              string   `json:"qr_code_image_url"`
	RealtimeSalesTip            string   `json:"realtime_sales_tip"`
	SalesTip                    string   `json:"sales_tip"`
	SearchId                    string   `json:"search_id"`
	ServTxt                     string   `json:"serv_txt"`
	ShareDesc                   string   `json:"share_desc"`
	ShareRate                   uint16   `json:"share_rate"`
	SubsidyAmount               uint64   `json:"subsidy_amount"`
	SubsidyGoodsType            int      `json:"subsidy_goods_type"`
	SubsidyDuoAmountTenMillion  uint16   `json:"subsidy_duo_amount_ten_million"`
	UnifiedTags                 []string `json:"unified_tags"`
}

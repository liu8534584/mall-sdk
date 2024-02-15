package jtt

type Response struct {
	Return      int64       `json:"requestId"`
	Result      interface{} `json:"result"`
	Second      interface{} `json:"second"`
	CurrentPage string      `json:"current_page"`
	TotalCount  int64       `json:"total_count"`
	TotalPage   int64       `json:"total_page"`
}

type RankListResponse struct {
	Response
	Data []RankList `json:"data"`
}

type RankList struct {
	JState             int         `json:"J_state"`
	JID                int         `json:"JID"`
	GoodsId            int64       `json:"goods_id"`
	ShortName          string      `json:"short_name"`
	GoodsName          string      `json:"goods_name"`
	GoodsContent       string      `json:"goods_content"`
	LinkContent        interface{} `json:"link_content"`
	CircleContent      interface{} `json:"circle_content"`
	GoodsLink          string      `json:"goods_link"`
	CommissionShare    float64     `json:"commissionShare"`
	BrandId            interface{} `json:"brand_id"`
	InOrderCount30Days int         `json:"inOrderCount30Days"`
	InOrderComm30Days  int         `json:"inOrderComm30Days"`
	BrandCode          int         `json:"brandCode"`
	BrandName          string      `json:"brandName"`
	IsHot              int         `json:"isHot"`
	Comments           int         `json:"comments"`
	GoodCommentsShare  float64     `json:"goodCommentsShare"`
	GoodsPrice         float64     `json:"goods_price"`
	FinalPrice         float64     `json:"final_price"`
	HistoryPriceDay    interface{} `json:"historyPriceDay"`
	GoodsImg           string      `json:"goods_img"`
	ImageList          string      `json:"imageList"`
	TodayNum           int         `json:"today_num"`
	DiscountLink       string      `json:"discount_link"`
	DiscountPrice      float64     `json:"discount_price"`
	GetStartTime       int64       `json:"get_start_time"`
	GetEndTime         int64       `json:"get_end_time"`
	ShopId             int         `json:"shop_id"`
	ShopName           string      `json:"shop_name"`
	Score              string      `json:"score"`
	Owner              string      `json:"owner"`
	GoodsType          int         `json:"goods_type"`
	GoodsSecondType    int         `json:"goods_second_type"`
	SpuId              int64       `json:"spu_id"`
	Cid1               int         `json:"cid1"`
	Cid1Name           string      `json:"cid1Name"`
	Cid2               int         `json:"cid2"`
	Cid2Name           string      `json:"cid2Name"`
	Cid3               int         `json:"cid3"`
	Cid3Name           string      `json:"cid3Name"`
	GoodsLevel         int         `json:"goods_level"`
	DeliveryType       int         `json:"deliveryType"`
	JdType             int         `json:"jd_type"`
	PingouPrice        interface{} `json:"pingouPrice"`
	PingouTmCount      interface{} `json:"pingouTmCount"`
	PingouUrl          interface{} `json:"pingouUrl"`
	PingouStartTime    interface{} `json:"pingouStartTime"`
	PingouEndTime      interface{} `json:"pingouEndTime"`
	JdVideoList        interface{} `json:"jd_videoList"`
	PcWareStyle        interface{} `json:"pc_ware_style"`
	DiscountCount      int         `json:"discount_count"`
	IsGoodShop         int         `json:"is_good_shop"`
	IsFlagshipStore    int         `json:"is_flagship_store"`
	IsSubsidyGoods     int         `json:"is_subsidy_goods"`
	HasCoupon          int         `json:"has_coupon"`
	ProType            int         `json:"pro_type"`
	ProCom             interface{} `json:"pro_com"`
	ProEndTime         interface{} `json:"pro_end_time"`
}

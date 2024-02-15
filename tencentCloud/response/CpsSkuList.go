package response

type CpsSkuList struct {
	Data []struct {
		MediaInfo struct {
			PrimaryImgs struct {
				ImgURL string `json:"img_url"`
			} `json:"primary_imgs"`
			Imgs []struct {
				ImgURL      string      `json:"img_url"`
				RichTextURL interface{} `json:"rich_text_url"`
			} `json:"imgs"`
			PropImgs []interface{} `json:"prop_imgs"`
		} `json:"media_info"`
		CategoryInfo []struct {
			CategoryType       float64 `json:"category_type"`
			CategoryLevel1ID   string  `json:"category_level_1_id"`
			CategoryLevel1Name string  `json:"category_level_1_name"`
			CategoryLevel2ID   string  `json:"category_level_2_id"`
			CategoryLevel2Name string  `json:"category_level_2_name"`
			CategoryLevel3ID   string  `json:"category_level_3_id"`
			CategoryLevel3Name string  `json:"category_level_3_name"`
		} `json:"category_info"`
		UniqID         string      `json:"uniq_id"`
		CatalogID      int         `json:"catalog_id"`
		ExternalSkuID  string      `json:"external_sku_id"`
		ExternalSpuID  string      `json:"external_spu_id"`
		UniSpuID       interface{} `json:"uni_spu_id"`
		SkuNameChinese string      `json:"sku_name_chinese"`
		IsAvailable    bool        `json:"is_available"`
		SalesInfo      []struct {
			IsAvailable         bool        `json:"is_available"`
			SkuStockStatus      int         `json:"sku_stock_status"`
			CurrentPrice        float64     `json:"current_price"`
			SkuPrice            float64     `json:"sku_price"`
			URLMiniprogram      string      `json:"url_miniprogram"`
			MiniprogramAppid    string      `json:"miniprogram_appid"`
			MiniprogramUsername string      `json:"miniprogram_username"`
			ExternalStoreID     string      `json:"external_store_id"`
			StoreName           string      `json:"store_name"`
			Logo                string      `json:"logo"`
			AfterCouponPrice    interface{} `json:"after_coupon_price"`
			URLStoreH5          string      `json:"url_store_h5"`
			URLStoreMiniprogram string      `json:"url_store_miniprogram"`
			URLH5               string      `json:"url_h5"`
			ShopOperationModel  interface{} `json:"shop_operation_model"`
		} `json:"sales_info"`
		CouponInfo         interface{} `json:"coupon_info"`
		ThirdPromotionInfo struct {
			CommissionRate              float64     `json:"commission_rate"`
			PromoteStatus               int         `json:"promote_status"`
			CommissionAmount            float64     `json:"commission_amount"`
			CommissionAmountAfterCoupon interface{} `json:"commission_amount_after_coupon"`
			StartTime                   string      `json:"start_time"`
			EndTime                     string      `json:"end_time"`
			InnerCommissionType         int         `json:"inner_commission_type"`
		} `json:"third_promotion_info"`
		AfterSaleInfo struct {
			ServiceType int `json:"service_type"`
		} `json:"after_sale_info"`
		ProductSource string `json:"product_source"`
		BrandInfo     struct {
			BrandWord string `json:"brand_word"`
		} `json:"brand_info"`
		WarehouseID         string      `json:"warehouse_id"`
		CustomTags          interface{} `json:"custom_tags"`
		BuyPrice            float64     `json:"buy_price"`
		BuyCommissionAmount float64     `json:"buy_commission_amount"`
		InsertTime          int64       `json:"insert_time"`
		UpdateTime          int64       `json:"update_time"`
		MerchantID          int         `json:"merchant_id"`
		MerchantName        string      `json:"merchant_name"`
		ProductSubSource    string      `json:"product_sub_source"`
		SaleNumTotal        int         `json:"sale_num_total"`
		UncouponInfoList    interface{} `json:"uncoupon_info_list"`
		PromotionInfo       interface{} `json:"promotion_info"`
		PlanSku             int         `json:"plan_sku"`
		SaleNumLast7Days    int         `json:"sale_num_last_7days"`
		SaleNumLast30Days   int         `json:"sale_num_last_30days"`
		PromotionInfos      interface{} `json:"promotion_infos"`
	} `json:"data"`
	Retcode  int         `json:"retcode"`
	Errmsg   string      `json:"errmsg"`
	MorePage bool        `json:"more_page"`
	ScrollID interface{} `json:"scroll_id"`
}

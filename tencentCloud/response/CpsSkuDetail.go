package response

type CpsSkuDetail struct {
	Data struct {
		UniqId         string `json:"uniq_id"`
		ExternalSkuId  string `json:"external_sku_id"`
		ExternalSpuId  string `json:"external_spu_id"`
		SkuNameChinese string `json:"sku_name_chinese"`
		IsAvailable    bool   `json:"is_available"`
		CatalogId      int    `json:"catalog_id"`
		MediaInfo      struct {
			PrimaryImgs struct {
				ImgUrl string `json:"img_url"`
			} `json:"primary_imgs"`
			Imgs []struct {
				ImgUrl string `json:"img_url"`
			} `json:"imgs"`
			PropImgs []struct {
				ImgUrl string `json:"img_url"`
			} `json:"prop_imgs"`
			DetailImgs []struct {
				ImgUrl string `json:"img_url"`
			} `json:"detail_imgs"`
			Videos []struct {
				VideoUrl string `json:"video_url,omitempty"`
				ImgUrl   string `json:"img_url,omitempty"`
			} `json:"videos"`
		} `json:"media_info"`
		ProductInfo struct {
			Color struct {
				ColorRgb  string `json:"color_rgb"`
				ColorName string `json:"color_name"`
			} `json:"color"`
			PropInfo []struct {
				PropId    int    `json:"prop_id"`
				PropName  string `json:"prop_name"`
				ValueId   int    `json:"value_id"`
				ValueName string `json:"value_name"`
			} `json:"prop_info"`
		} `json:"product_info"`
		CategoryInfo []struct {
			CategoryType       int    `json:"category_type"`
			CategoryLevel1Id   string `json:"category_level1_id"`
			CategoryLevel1Name string `json:"category_level1_name"`
			CategoryLevel2Id   string `json:"category_level2_id"`
			CategoryLevel2Name string `json:"category_level2_name"`
			CategoryLevel3Id   string `json:"category_level3_id"`
			CategoryLevel3Name string `json:"category_level3_name"`
		} `json:"category_info"`
		SalesInfo []struct {
			IsAvailable         bool        `json:"is_available"`
			SkuStockStatus      int         `json:"sku_stock_status"`
			CurrentPrice        float64     `json:"current_price"`
			SkuPrice            float64     `json:"sku_price"`
			UrlMiniprogram      string      `json:"url_miniprogram"`
			MiniprogramAppid    string      `json:"miniprogram_appid"`
			MiniprogramUsername string      `json:"miniprogram_username"`
			ExternalStoreId     string      `json:"external_store_id"`
			StoreName           string      `json:"store_name"`
			Logo                string      `json:"logo"`
			AfterCouponPrice    interface{} `json:"after_coupon_price"`
			UrlStoreH5          string      `json:"url_store_h5"`
			UrlStoreMiniprogram string      `json:"url_store_miniprogram"`
			UrlH5               string      `json:"url_h5"`
			DeliverInfos        []struct {
				FreightAmount float64 `json:"freight_amount"`
				Location      []struct {
					CountryCode  int    `json:"country_code"`
					CityName     string `json:"city_name"`
					CountryName  string `json:"country_name"`
					ProvinceName string `json:"province_name"`
					ProvinceCode int    `json:"province_code"`
					CityCode     int    `json:"city_code"`
				} `json:"location"`
			} `json:"deliver_infos"`
		} `json:"sales_info"`
		CouponInfo []struct {
			CouponId              string `json:"coupon_id"`
			AmountCoupon          int    `json:"amount_coupon"`
			UrlList               string `json:"url_list"`
			AmountMinimum         int    `json:"amount_minimum"`
			ShowStartTime         string `json:"show_start_time"`
			ShowEndTime           string `json:"show_end_time"`
			StartTime             string `json:"start_time"`
			EndTime               string `json:"end_time"`
			AvailableNumPerPerson int    `json:"available_num_per_person"`
			TotalNum              int    `json:"total_num"`
			StoreNum              int    `json:"store_num"`
			Status                int    `json:"status"`
		} `json:"coupon_info"`
		CommentInfo struct {
			CommentNum            int     `json:"comment_num"`
			PositiveCommentRating float64 `json:"positive_comment_rating"`
		} `json:"comment_info"`
		ThirdPromotionInfo struct {
			CommissionRate              float64     `json:"commission_rate"`
			PromoteStatus               int         `json:"promote_status"`
			CommissionAmount            float64     `json:"commission_amount"`
			CommissionAmountAfterCoupon interface{} `json:"commission_amount_after_coupon"`
			StartTime                   string      `json:"start_time"`
			EndTime                     string      `json:"end_time"`
		} `json:"third_promotion_info"`
		AfterSaleInfo struct {
			ServiceType int `json:"service_type"`
		} `json:"after_sale_info"`
		ProductSource string `json:"product_source"`
		WarehouseId   int    `json:"warehouse_id"`
	} `json:"data"`
	Retcode  int    `json:"retcode"`
	Errmsg   string `json:"errmsg"`
	MorePage bool   `json:"more_page"`
}

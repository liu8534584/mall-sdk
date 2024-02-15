package response

type Act struct {
	ActRoleList         []ActRole `json:"actRoleList"`
	ActivityDescription string    `json:"activityDescription"`
	ActivityID          string    `json:"activityId"`
	ActivityTypeID      string    `json:"activityTypeId"`
}

type ActRole struct {
	BaseAddUpType         string      `json:"baseAddUpType"`
	BaseAmount            string      `json:"baseAmount"`
	BaseQuantifierType    string      `json:"baseQuantifierType"`
	CouponAmount          interface{} `json:"couponAmount"`
	Floor                 string      `json:"floor"`
	RewardAmount          string      `json:"rewardAmount"`
	RewardQuantifierType  string      `json:"rewardQuantifierType"`
	ScopeType             string      `json:"scopeType"`
	SingleMaxRewardAmount string      `json:"singleMaxRewardAmount"`
}

type CategoryInfo struct {
	FirstPurchaseCategoryID    string `json:"firstPurchaseCategoryId"`
	FirstPurchaseCategoryName  string `json:"firstPurchaseCategoryName"`
	FirstSaleCategoryID        string `json:"firstSaleCategoryId"`
	FirstSaleCategoryName      string `json:"firstSaleCategoryName"`
	GoodsGroupCategoryID       string `json:"goodsGroupCategoryId"`
	GoodsGroupCategoryName     string `json:"goodsGroupCategoryName"`
	SecondPurchaseCategoryID   string `json:"secondPurchaseCategoryId"`
	SecondPurchaseCategoryName string `json:"secondPurchaseCategoryName"`
	SecondSaleCategoryID       string `json:"secondSaleCategoryId"`
	SecondSaleCategoryName     string `json:"secondSaleCategoryName"`
	ThirdPurchaseCategoryID    string `json:"thirdPurchaseCategoryId"`
	ThirdPurchaseCategoryName  string `json:"thirdPurchaseCategoryName"`
	ThirdSaleCategoryID        string `json:"thirdSaleCategoryId"`
	ThirdSaleCategoryName      string `json:"thirdSaleCategoryName"`
}

type CmmdtyReviewInfo struct {
	GoodRate         string `json:"goodRate"`
	GoodReviewCount  int64  `json:"goodReviewCount"`
	TotalReviewCount int64  `json:"totalReviewCount"`
}

type CommodityInfo struct {
	Baoyou         int64  `json:"baoyou"`
	CommodityCode  string `json:"commodityCode"`
	CommodityName  string `json:"commodityName"`
	CommodityPrice string `json:"commodityPrice"`
	CommodityType  int64  `json:"commodityType"`
	ForbiddenMark  string `json:"forbiddenMark"`
	IsAct          string `json:"isAct"`
	MonthSales     int64  `json:"monthSales"`
	PictureURL     []struct {
		LocationID int64  `json:"locationId"`
		PicURL     string `json:"picUrl"`
	} `json:"pictureUrl"`
	PriceType     string `json:"priceType"`
	PriceTypeCode int64  `json:"priceTypeCode"`
	ProductURL    string `json:"productUrl"`
	Rate          string `json:"rate"`
	SaleStatus    int64  `json:"saleStatus"`
	SellingPoint  string `json:"sellingPoint"`
	SnPrice       string `json:"snPrice"`
	SupplierCode  string `json:"supplierCode"`
	SupplierName  string `json:"supplierName"`
}

type CouponInfo struct {
	ActivityDescription  string      `json:"activityDescription"`
	ActivityID           string      `json:"activityId"`
	ActivitySecretKey    string      `json:"activitySecretKey"`
	AfterCouponPrice     interface{} `json:"afterCouponPrice"` //会返回string 和 float64
	BounsAmount          string      `json:"bounsAmount"`
	BounsLimit           string      `json:"bounsLimit"`
	CouponCount          string      `json:"couponCount"`
	CouponEndTime        string      `json:"couponEndTime"`
	CouponShowType       string      `json:"couponShowType"`
	CouponStartTime      string      `json:"couponStartTime"`
	CouponURL            string      `json:"couponUrl"`
	CouponValue          string      `json:"couponValue"`
	DescriptionForList   string      `json:"descriptionForList"`
	EndTime              string      `json:"endTime"`
	PreferentialDistinct string      `json:"preferentialDistinct"`
	StartTime            string      `json:"startTime"`
}

type Parameters struct {
	CoreFlag          string      `json:"coreFlag"`
	Explain           interface{} `json:"explain"`
	ParameterCode     string      `json:"parameterCode"`
	ParameterDesc     string      `json:"parameterDesc"`
	ParameterSequence string      `json:"parameterSequence"`
	ParameterVal      string      `json:"parameterVal"`
	ParametersCode    string      `json:"parametersCode"`
	ParametersDesc    string      `json:"parametersDesc"`
	Sequence          string      `json:"sequence"`
}

type PgInfo struct {
	MinOrderQuantity string `json:"minOrderQuantity"`
	PgActionID       string `json:"pgActionId"`
	PgNum            string `json:"pgNum"`
	PgPrice          string `json:"pgPrice"`
	PgURL            string `json:"pgUrl"`
}

type AdvanceSale struct {
	DepositAmount      string `json:"depositAmount"`
	DepositEndTime     string `json:"depositEndTime"`
	IsReserveCommodity string `json:"isReserveCommodity"`
}

type OrderActivity struct {
	ActivityRuleList    []ActivityRule `json:"activityRuleList"`
	ActivityDescription string         `json:"activityDescription"`
}

type ActivityRule struct {
	// 层级
	Floor string `json:"floor"`
	// 基数数值
	BaseAmount string `json:"baseAmount"`
	// 奖励数值
	RewardAmount string `json:"rewardAmount"`
}

type Detail struct {
	ActList            []Act            `json:"actList"`
	AdvanceSale        AdvanceSale      `json:"advanceSale"`
	ArrivalPrice       float64          `json:"arrivalPrice"`
	CategoryInfo       CategoryInfo     `json:"categoryInfo"`
	CmmdtyReviewInfo   CmmdtyReviewInfo `json:"cmmdtyReviewInfo"`
	CommodityInfo      CommodityInfo    `json:"commodityInfo"`
	CouponInfo         CouponInfo       `json:"couponInfo"`
	CouponInfoList     []interface{}    `json:"couponInfoList"`
	DiscountCouponList []interface{}    `json:"discountCouponList"`
	OrderActivity      OrderActivity    `json:"orderActivity"`
	ParametersList     []Parameters     `json:"parametersList"`
	PgInfo             PgInfo           `json:"pgInfo"`
	SubCodeList        []SubCode        `json:"subCodeList"`
}

type QueryCommodityDetailResponse struct {
	SnResponseContent struct {
		SnBody struct {
			QueryCommoditydetail []Detail `json:"queryCommoditydetail"`
		} `json:"sn_body"`
	} `json:"sn_responseContent"`
}

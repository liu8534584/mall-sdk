package qianzhuResponse

type AccessTokenResponse struct {
	Success bool   `json:"success"`
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		UserID              int         `json:"userId"`
		AccessToken         string      `json:"accessToken"`
		AccessTokenExpireIn int         `json:"accessTokenExpireIn"`
		RefreshToken        string      `json:"refreshToken"`
		Roles               []string    `json:"roles"`
		NewUser             bool        `json:"newUser"`
		CityID              interface{} `json:"cityId"`
		CityName            interface{} `json:"cityName"`
	} `json:"data"`
	TraceID string `json:"traceId"`
}

type PayResponse struct {
	Success bool   `json:"success"`
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		Balance int `json:"balance"`
	} `json:"data"`
	TraceID string `json:"traceId"`
}

type OrderDetailResponse struct {
	Success bool   `json:"success"`
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		CreateTime           string      `json:"createTime"`
		UpdateTime           string      `json:"updateTime"`
		UserID               int         `json:"userId"`
		OrderNo              string      `json:"orderNo"`
		Status               int         `json:"status"`
		StatusDesc           string      `json:"statusDesc"`
		UnitPrice            float64     `json:"unitPrice"`
		Quantity             int         `json:"quantity"`
		TotalPrice           float64     `json:"totalPrice"`
		MarketUnitPrice      float64     `json:"marketUnitPrice"`
		PaymentTime          string      `json:"paymentTime"`
		PaymentExpireSeconds interface{} `json:"paymentExpireSeconds"`
		Amount               float64     `json:"amount"`
		DrawTime             string      `json:"drawTime"`
		Ticket               string      `json:"ticket"`
		UserRemark           interface{} `json:"userRemark"`
		CommissionPrice      float64     `json:"commissionPrice"`
		UserMobile           string      `json:"userMobile"`
		CancelTime           interface{} `json:"cancelTime"`
		UserName             string      `json:"userName"`
		RefundPrice          int         `json:"refundPrice"`
		KfcOrderMobileSuffix string      `json:"kfcOrderMobileSuffix"`
		KfcOrderMobileRemark string      `json:"kfcOrderMobileRemark"`
		KfcPlaceOrder        struct {
			CityName     string `json:"cityName"`
			StoreName    string `json:"storeName"`
			StoreAddress string `json:"storeAddress"`
			StoreCode    string `json:"storeCode"`
			EatType      int    `json:"eatType"`
			Items        []struct {
				ProductID   string  `json:"productId"`
				ProductName string  `json:"productName"`
				Quantity    int     `json:"quantity"`
				Price       float64 `json:"price"`
				ImageURL    string  `json:"imageUrl"`
				Canceled    bool    `json:"canceled"`
				OptionItems []struct {
					ProductName string `json:"productName"`
					Quantity    int    `json:"quantity"`
					ImageURL    string `json:"imageUrl"`
				} `json:"optionItems"`
			} `json:"items"`
		} `json:"kfcPlaceOrder"`
		PlatformUniqueID     string  `json:"platformUniqueId"`
		Takeout              bool    `json:"takeout"`
		TakeoutPrice         float64 `json:"takeoutPrice"`
		TakeoutAddress       string  `json:"takeoutAddress"`
		TakeoutProvinceName  string  `json:"takeoutProvinceName"`
		TakeoutProvinceCode  string  `json:"takeoutProvinceCode"`
		TakeoutCityName      string  `json:"takeoutCityName"`
		TakeoutCityCode      string  `json:"takeoutCityCode"`
		TakeoutRegionName    string  `json:"takeoutRegionName"`
		TakeoutRegionCode    string  `json:"takeoutRegionCode"`
		TakeoutMainAddress   string  `json:"takeoutMainAddress"`
		TakeoutDetailAddress string  `json:"takeoutDetailAddress"`
		TakeoutLon           string  `json:"takeoutLon"`
		TakeoutLat           string  `json:"takeoutLat"`
		TakeoutDeliveryTime  string  `json:"takeoutDeliveryTime"`
	} `json:"data"`
}

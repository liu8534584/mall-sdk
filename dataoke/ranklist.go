package dataoke

type Response struct {
	RequestID string      `json:"requestId"`
	Time      int64       `json:"time"`
	Code      int         `json:"code"`
	Msg       string      `json:"msg"`
	Data      interface{} `json:"data"`
}

type RankListResponse struct {
	Response
	Data []RankList `json:"data"`
}

type RankList struct {
	ID                     int           `json:"id"`
	GoodsID                string        `json:"goodsId"`
	Ranking                int           `json:"ranking"`
	Dtitle                 string        `json:"dtitle"`
	ActualPrice            float64       `json:"actualPrice"`
	CommissionRate         float64       `json:"commissionRate"`
	CouponPrice            int           `json:"couponPrice"`
	CouponReceiveNum       int           `json:"couponReceiveNum"`
	CouponTotalNum         int           `json:"couponTotalNum"`
	MonthSales             int           `json:"monthSales"`
	TwoHoursSales          int           `json:"twoHoursSales"`
	DailySales             int           `json:"dailySales"`
	HotPush                int           `json:"hotPush"`
	MainPic                string        `json:"mainPic"`
	Title                  string        `json:"title"`
	Desc                   string        `json:"desc"`
	OriginalPrice          float64       `json:"originalPrice"`
	CouponLink             string        `json:"couponLink"`
	CouponStartTime        string        `json:"couponStartTime"`
	CouponEndTime          string        `json:"couponEndTime"`
	CommissionType         int           `json:"commissionType"`
	CreateTime             string        `json:"createTime"`
	ActivityType           int           `json:"activityType"`
	GuideName              string        `json:"guideName"`
	ShopType               int           `json:"shopType"`
	CouponConditions       interface{}   `json:"couponConditions"`
	NewRankingGoods        int           `json:"newRankingGoods"`
	SellerID               string        `json:"sellerId"`
	QuanMLink              int           `json:"quanMLink"`
	HzQuanOver             int           `json:"hzQuanOver"`
	Yunfeixian             int           `json:"yunfeixian"`
	EstimateAmount         int           `json:"estimateAmount"`
	FreeshipRemoteDistrict int           `json:"freeshipRemoteDistrict"`
	Haitao                 int           `json:"haitao"`
	Tchaoshi               int           `json:"tchaoshi"`
	Lowest                 int           `json:"lowest"`
	Fresh                  int           `json:"fresh"`
	Num                    int           `json:"num"`
	SpecialText            []interface{} `json:"specialText"`
	MarketGroup            []interface{} `json:"marketGroup"`
}

package response

type GetOrderListResp struct {
	SnResponseContent struct {
		SnHead struct {
			TotalSize     string `json:"totalSize"`
			PageTotal     string `json:"pageTotal"`
			PageNo        string `json:"pageNo"`
			ReturnMessage string `json:"returnMessage"`
		} `json:"sn_head"`
		SnBody struct {
			QueryOrder []QueryOrder `json:"queryOrder"`
		} `json:"sn_body"`
	} `json:"sn_responseContent"`
}

type QueryOrder struct {
	OrderCode   string        `json:"orderCode"`
	OrderDetail []OrderDetail `json:"orderDetail"`
}

type OrderDetail struct {
	GoodsGroupCatalog         string `json:"goodsGroupCatalog"`
	OrderType                 string `json:"orderType"`
	PayTime                   string `json:"payTime"`
	SaleType                  string `json:"saleType"`
	ProductSecondCatalog      string `json:"productSecondCatalog"`
	OrderLineStatusDesc       string `json:"orderLineStatusDesc"`
	ProductName               string `json:"productName"`
	OrderLineFlag             string `json:"orderLineFlag"`
	StatParam                 string `json:"statParam"`
	SellerCode                string `json:"sellerCode"`
	PayAmount                 string `json:"payAmount"`
	OrderLineOrigin           string `json:"orderLineOrigin"`
	ProductFirstCatalog       string `json:"productFirstCatalog"`
	GoodsNum                  string `json:"goodsNum"`
	OrderLineStatus           string `json:"orderLineStatus"`
	CustNo                    string `json:"custNo"`
	SellName                  string `json:"sellName"`
	PictureUrl                string `json:"pictureUrl"`
	ProductThirdCatalog       string `json:"productThirdCatalog"`
	ConfirmTime               string `json:"confirmTime"`
	SaleNum                   string `json:"saleNum"`
	CommissionRatio           string `json:"commissionRatio"`
	ReturnCommission          int    `json:"returnCommission"`
	PositionId                string `json:"positionId"`
	PrePayCommission          string `json:"prePayCommission"`
	Violation                 int    `json:"violation"`
	OrderSubmitTime           string `json:"orderSubmitTime"`
	OrderLineNumber           string `json:"orderLineNumber"`
	OrderLineStatusChangeTime string `json:"orderLineStatusChangeTime"`
	ChildAccountId            string `json:"childAccountId"`
	Promotion                 int    `json:"promotion"`
}

type OrderInfoResp struct {
	SnResponseContent struct {
		SnBody struct {
			GetOrder struct {
				OrderInfo []OrderDetail `json:"orderInfo"`
				OrderCode string        `json:"orderCode"`
			} `json:"getOrder"`
		} `json:"sn_body"`
	} `json:"sn_responseContent"`
}

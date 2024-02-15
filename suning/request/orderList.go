package request

type OrderListReq struct {
	StartTime       string `json:"startTime"` //y-m-d h:i:s
	EndTime         string `json:"endTime"`   //y-m-d h:i:s
	PageNo          int    `json:"pageNo"`
	PageSize        int    `json:"pageSize"`
	OrderLineStatus int    `json:"orderLineStatus"`
}

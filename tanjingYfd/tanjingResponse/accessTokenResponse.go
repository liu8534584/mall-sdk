package tanjingResponse

type AccessTokenResp struct {
	Code        int    `json:"code"`
	Token       string `json:"token"`
	Merchant    string `json:"merchant"`
	Result      string `json:"result"`
	SerialNo    string `json:"serial_no"`
	ServiceTime string `json:"service_time"`
}

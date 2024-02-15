package kuaishouapi

type LinkParseRequest struct {
	CpsLink string `json:"cpsLink"` //口令
}

type LinkParseResponse struct {
	Code     string                `json:"code"`      // 1 主返回码
	Msg      string                `json:"msg"`       // success 主返回信息
	SubCode  string                `json:"sub_code"`  //1 子返回码
	SubMsg   string                `json:"sub_msg"`   // SUCCESS 子返回信息
	Result   int                   `json:"result"`    // 1 返回码
	ErrorMsg string                `json:"error_msg"` // SUCCESS 错误信息
	Data     LinkParseResponseData `json:"data"`      //list  商品列表
}

type LinkParseResponseData struct {
	ItemId int64 `json:"itemId"`
}

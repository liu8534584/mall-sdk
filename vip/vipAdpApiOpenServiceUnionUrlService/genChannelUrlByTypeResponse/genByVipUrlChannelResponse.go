package genChennelUrlResponse

type VipGenByChannelTypeResponse struct {
	ReturnCode string                            `json:"returnCode"`
	Result     VipGenByChannelTypeResponseResult `json:"result"`
}

type VipGenByChannelTypeResponseResult struct {
	Msg  string `json:"msg"`
	Code int64  `json:"code"`
	Data struct {
		ShortUrl string `json:"shortUrl"`
	} `json:"data"`
}

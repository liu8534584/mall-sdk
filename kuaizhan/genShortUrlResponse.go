package Kuaizhan

type GenShortUrlResp struct {
	Msg       string `json:"msg"`
	Path      string `json:"path"`
	Code      int    `json:"code"`
	RequestId string `json:"requestId"`
	Domain    string `json:"domain"`
	Url       string `json:"url"`
}

package response

type GenerateUrl struct {
	Data struct {
		H5Url    string `json:"h5Url"`
		ShortUrl string `json:"shortUrl"`
		Path     string `json:"path"`
		Query    string `json:"query"`
		AppId    string `json:"appId"`
		UserName string `json:"userName"`
	} `json:"data"`
	Code      string      `json:"code"`
	Msg       string      `json:"msg"`
	ErrorMsg  string      `json:"errorMsg"`
	TraceId   string      `json:"traceId"`
	RequestId string      `json:"requestId"`
	Rt        int         `json:"rt"`
	ClientIp  interface{} `json:"clientIp"`
	Success   bool        `json:"success"`
}

package tanjingYfd

type TanjingBaseConfig struct {
	Merchant string `json:"merchant"`
	Secret   string `json:"secret"`
	Timeout  int64  `json:"timeout"`
	BaseUrl  string `json:"baseUrl"`
}

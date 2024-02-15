package doudian_sdk

type DoudianOpConfig struct {
	AppKey          string
	AppSecret       string
	HttpReadTimeout int64
	OpenRequestUrl  string
}

func NewDoudianOpConfig() *DoudianOpConfig {
	config := &DoudianOpConfig{
		HttpReadTimeout: 5000, //5s超时
		OpenRequestUrl:  "https://openapi-fxg.jinritemai.com",
	}
	return config
}

var GlobalConfig *DoudianOpConfig = NewDoudianOpConfig()

var GlobalLiveConfig *DoudianOpConfig = NewDoudianOpConfig()

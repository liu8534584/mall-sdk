package czb

type ApiRequest interface {
	GetParams() map[string]interface{}
	GetUrlPath() string
}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Result  interface{} `json:"result"`
}

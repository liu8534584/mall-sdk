package douyin_life

type ApiRequest interface {
	GetParams() interface{}
	GetUrlPath() string
}

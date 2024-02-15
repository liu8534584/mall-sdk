package kaola

type KaolaBaseRequest interface {
	GetApiParams() map[string]interface{}
	GetApiMethodName() string
}

package pdd

type Request interface {
	GetApiParams() map[string]interface{}
	GetApiMethodName() string
}

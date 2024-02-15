package base

type Code int

func (code Code) String() string {
	return codeText[code]
}

const (
	OK                  Code = 10001
	NeedLogin           Code = 10002
	InternalServerError Code = 50000
)

var codeText = map[Code]string{
	OK:                  "OK",
	InternalServerError: "内部错误",
	NeedLogin:           "未登陆",
}

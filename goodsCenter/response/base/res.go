package base

type Demotion struct {
	Cache  int `json:"cache"`
	Time   int `json:"time"`
	Forbid int `json:"forbid"`
}

type Body struct {
	Code     Code        `json:"code"`
	Message  string      `json:"msg"`
	Demotion Demotion    `json:"demotion"`
	Data     interface{} `json:"data"`
}

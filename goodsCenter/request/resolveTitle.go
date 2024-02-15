package request

type ResolveTitle struct {
	Content       string        `json:"content"`
	TaobaoPidInfo TaobaoPidInfo `json:"taobaoPid"`
	CustomParams  string        `json:"customParams"`
	PddPidInfo    PddPidInfo    `json:"pddPidInfo"`
	JdChannelId   int           `json:"jdChannelId"`
}

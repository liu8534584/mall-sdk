package tanjingResponse

type GetTagInFriendsResp struct {
	Data       FriendData `json:"Data"`
	NResult    int        `json:"nResult"`
	VcResult   string     `json:"vcResult"`
	VcSerialNo string     `json:"vcSerialNo"`
}

type FriendData struct {
	NTagId          int          `json:"nTagId"`
	VcTagName       string       `json:"vcTagName"`
	NFriendCount    int          `json:"nFriendCount"`
	VcBase64TagName string       `json:"vcBase64TagName"`
	Friends         []FriendInfo `json:"Friends"`
}

type FriendInfo struct {
	VcFriendWxId     string `json:"vcFriendWxId"`
	VcFriendSerialNo string `json:"vcFriendSerialNo"`
	NSex             int    `json:"nSex"`
	VcNickName       string `json:"vcNickName"`
	VcBase64NickName string `json:"vcBase64NickName"`
	VcHeadImgUrl     string `json:"vcHeadImgUrl"`
}

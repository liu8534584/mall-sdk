package tanjingResponse

type GetChatRoomListResp struct {
	Data []struct {
		VcChatRoomSerialNo    string  `json:"vcChatRoomSerialNo"`
		VcChatRoomId          string  `json:"vcChatRoomId"`
		NUserCount            int     `json:"nUserCount"`
		VcName                *string `json:"vcName"`
		VcHeadImg             string  `json:"vcHeadImg"`
		VcBase64Name          *string `json:"vcBase64Name"`
		VcAdminWxId           string  `json:"vcAdminWxId"`
		VcAdminWxUserSerialNo string  `json:"vcAdminWxUserSerialNo"`
		VcChatRoomQRCode      string  `json:"vcChatRoomQRCode"`
		DtCreateDate          string  `json:"dtCreateDate"`
		IsOpenMessage         int     `json:"isOpenMessage"`
		NIsInContacts         int     `json:"nIsInContacts"`
		NRobotInStatus        int     `json:"nRobotInStatus"`
		IsEnterpriseChatroom  bool    `json:"IsEnterpriseChatroom"`
	} `json:"Data"`
	NResult    int    `json:"nResult"`
	VcResult   string `json:"vcResult"`
	VcSerialNo string `json:"vcSerialNo"`
}

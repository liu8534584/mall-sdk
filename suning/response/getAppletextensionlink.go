package response

type GetAppletextensionlinkResponse struct {
	SnResponseContent struct {
		SnBody struct {
			GetAppletextensionlink struct {
				AppID           string `json:"appId"`
				AppletExtendURL string `json:"appletExtendUrl"`
			} `json:"getAppletextensionlink"`
		} `json:"sn_body"`
	} `json:"sn_responseContent"`
}

package response

type GetExtensionlinkResponse struct {
	SnResponseContent struct {
		SnBody struct {
			GetExtensionlink struct {
				ShortLink string `json:"shortLink"`
			} `json:"getExtensionlink"`
		} `json:"sn_body"`
	} `json:"sn_responseContent"`
}

package response

type QueryBacthcustomlinkResponse struct {
	SnResponseContent struct {
		SnBody struct {
			QueryBacthcustomlink struct {
				ShortLink string `json:"shortLink"`
			} `json:"queryBacthcustomlink"`
		} `json:"sn_body"`
	} `json:"sn_responseContent"`
}

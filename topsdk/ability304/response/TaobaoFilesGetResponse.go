package response

import (
	"github.com/liu8534584/mall-sdk/topsdk/ability304/domain"
)

type TaobaoFilesGetResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   results
	*/
	Results []domain.TaobaoFilesGetTopDownloadRecordDo `json:"results,omitempty" `
}

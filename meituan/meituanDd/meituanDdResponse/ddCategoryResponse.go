package meituanDdResponse

type DdCategoryResponse struct {
	Code int64             `json:"code"`
	Msg  AduCategoryInfoVO `json:"msg"`
}

type AduCategoryInfoVO struct {
	ParentId   int64             `json:"parentId"`
	Categories map[string]string `json:"categories"`
}

package request


type TaobaoTbkScMaterialRecommendRequest struct {
    /*
        页大小，默认20，1~100 defalutValue��20    */
    PageSize  *int64 `json:"page_size,omitempty" required:"false" `
    /*
        第几页，默认：1 defalutValue��1    */
    PageNo  *int64 `json:"page_no,omitempty" required:"false" `
    /*
        官方提供的物料Id     */
    MaterialId  *int64 `json:"material_id" required:"true" `
    /*
        mm_xxx_xxx_xxx的第3段数字     */
    AdzoneId  *int64 `json:"adzone_id" required:"true" `
    /*
        mm_xxx_xxx_xxx的第2段数字     */
    SiteId  *int64 `json:"site_id" required:"true" `
}

func (s *TaobaoTbkScMaterialRecommendRequest) SetPageSize(v int64) *TaobaoTbkScMaterialRecommendRequest {
    s.PageSize = &v
    return s
}
func (s *TaobaoTbkScMaterialRecommendRequest) SetPageNo(v int64) *TaobaoTbkScMaterialRecommendRequest {
    s.PageNo = &v
    return s
}
func (s *TaobaoTbkScMaterialRecommendRequest) SetMaterialId(v int64) *TaobaoTbkScMaterialRecommendRequest {
    s.MaterialId = &v
    return s
}
func (s *TaobaoTbkScMaterialRecommendRequest) SetAdzoneId(v int64) *TaobaoTbkScMaterialRecommendRequest {
    s.AdzoneId = &v
    return s
}
func (s *TaobaoTbkScMaterialRecommendRequest) SetSiteId(v int64) *TaobaoTbkScMaterialRecommendRequest {
    s.SiteId = &v
    return s
}

func (req *TaobaoTbkScMaterialRecommendRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.PageSize != nil) {
        paramMap["page_size"] = *req.PageSize
    }
    if(req.PageNo != nil) {
        paramMap["page_no"] = *req.PageNo
    }
    if(req.MaterialId != nil) {
        paramMap["material_id"] = *req.MaterialId
    }
    if(req.AdzoneId != nil) {
        paramMap["adzone_id"] = *req.AdzoneId
    }
    if(req.SiteId != nil) {
        paramMap["site_id"] = *req.SiteId
    }
    return paramMap
}

func (req *TaobaoTbkScMaterialRecommendRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}
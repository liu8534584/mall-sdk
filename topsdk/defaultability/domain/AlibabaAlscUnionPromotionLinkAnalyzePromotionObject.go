package domain


type AlibabaAlscUnionPromotionLinkAnalyzePromotionObject struct {
    /*
        推广对象类型（1-商品；2-店铺；3-活动；4-卡券；5-SKU；6-营销活动）     */
    ObjectType  *int64 `json:"object_type,omitempty" `

    /*
        推广对象ID     */
    ObjectId  *string `json:"object_id,omitempty" `

}

func (s *AlibabaAlscUnionPromotionLinkAnalyzePromotionObject) SetObjectType(v int64) *AlibabaAlscUnionPromotionLinkAnalyzePromotionObject {
    s.ObjectType = &v
    return s
}
func (s *AlibabaAlscUnionPromotionLinkAnalyzePromotionObject) SetObjectId(v string) *AlibabaAlscUnionPromotionLinkAnalyzePromotionObject {
    s.ObjectId = &v
    return s
}

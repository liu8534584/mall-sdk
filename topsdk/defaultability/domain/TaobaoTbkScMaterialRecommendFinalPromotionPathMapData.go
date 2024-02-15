package domain


type TaobaoTbkScMaterialRecommendFinalPromotionPathMapData struct {
    /*
        优惠名称，如“商品券”、“跨店满减”、“单品直降”等     */
    PromotionTitle  *string `json:"promotion_title,omitempty" `

    /*
        优惠利益点文案，如“1件7.92折”、“每200减20”等     */
    PromotionDesc  *string `json:"promotion_desc,omitempty" `

    /*
        实际优惠金额（元）     */
    PromotionFee  *string `json:"promotion_fee,omitempty" `

    /*
        优惠开始时间     */
    PromotionStartTime  *string `json:"promotion_start_time,omitempty" `

    /*
        优惠结束时间     */
    PromotionEndTime  *string `json:"promotion_end_time,omitempty" `

}

func (s *TaobaoTbkScMaterialRecommendFinalPromotionPathMapData) SetPromotionTitle(v string) *TaobaoTbkScMaterialRecommendFinalPromotionPathMapData {
    s.PromotionTitle = &v
    return s
}
func (s *TaobaoTbkScMaterialRecommendFinalPromotionPathMapData) SetPromotionDesc(v string) *TaobaoTbkScMaterialRecommendFinalPromotionPathMapData {
    s.PromotionDesc = &v
    return s
}
func (s *TaobaoTbkScMaterialRecommendFinalPromotionPathMapData) SetPromotionFee(v string) *TaobaoTbkScMaterialRecommendFinalPromotionPathMapData {
    s.PromotionFee = &v
    return s
}
func (s *TaobaoTbkScMaterialRecommendFinalPromotionPathMapData) SetPromotionStartTime(v string) *TaobaoTbkScMaterialRecommendFinalPromotionPathMapData {
    s.PromotionStartTime = &v
    return s
}
func (s *TaobaoTbkScMaterialRecommendFinalPromotionPathMapData) SetPromotionEndTime(v string) *TaobaoTbkScMaterialRecommendFinalPromotionPathMapData {
    s.PromotionEndTime = &v
    return s
}

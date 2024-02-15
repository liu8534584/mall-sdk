package domain


type TaobaoTbkScMaterialRecommendMorePromotionMapData struct {
    /*
        预热优惠名称，如“商品券”、“跨店满减”、“单品直降”等     */
    PromotionTitle  *string `json:"promotion_title,omitempty" `

    /*
        预热优惠利益点文案，如“1件7.92折”、“每200减20”等     */
    PromotionDesc  *string `json:"promotion_desc,omitempty" `

    /*
        优惠开始时间     */
    PromotionStartTime  *string `json:"promotion_start_time,omitempty" `

    /*
        优惠结束时间     */
    PromotionEndTime  *string `json:"promotion_end_time,omitempty" `

}

func (s *TaobaoTbkScMaterialRecommendMorePromotionMapData) SetPromotionTitle(v string) *TaobaoTbkScMaterialRecommendMorePromotionMapData {
    s.PromotionTitle = &v
    return s
}
func (s *TaobaoTbkScMaterialRecommendMorePromotionMapData) SetPromotionDesc(v string) *TaobaoTbkScMaterialRecommendMorePromotionMapData {
    s.PromotionDesc = &v
    return s
}
func (s *TaobaoTbkScMaterialRecommendMorePromotionMapData) SetPromotionStartTime(v string) *TaobaoTbkScMaterialRecommendMorePromotionMapData {
    s.PromotionStartTime = &v
    return s
}
func (s *TaobaoTbkScMaterialRecommendMorePromotionMapData) SetPromotionEndTime(v string) *TaobaoTbkScMaterialRecommendMorePromotionMapData {
    s.PromotionEndTime = &v
    return s
}

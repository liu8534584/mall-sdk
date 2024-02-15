package domain


type TaobaoTbkScMaterialOptionalUpgradeMorePromotionMapData struct {
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

func (s *TaobaoTbkScMaterialOptionalUpgradeMorePromotionMapData) SetPromotionTitle(v string) *TaobaoTbkScMaterialOptionalUpgradeMorePromotionMapData {
    s.PromotionTitle = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalUpgradeMorePromotionMapData) SetPromotionDesc(v string) *TaobaoTbkScMaterialOptionalUpgradeMorePromotionMapData {
    s.PromotionDesc = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalUpgradeMorePromotionMapData) SetPromotionStartTime(v string) *TaobaoTbkScMaterialOptionalUpgradeMorePromotionMapData {
    s.PromotionStartTime = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalUpgradeMorePromotionMapData) SetPromotionEndTime(v string) *TaobaoTbkScMaterialOptionalUpgradeMorePromotionMapData {
    s.PromotionEndTime = &v
    return s
}

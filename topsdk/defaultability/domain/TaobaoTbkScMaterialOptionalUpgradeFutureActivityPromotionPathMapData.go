package domain


type TaobaoTbkScMaterialOptionalUpgradeFutureActivityPromotionPathMapData struct {
    /*
        预热优惠名称，如“商品券”、“跨店满减”、“单品直降”等     */
    PromotionTitle  *string `json:"promotion_title,omitempty" `

    /*
        预热优惠利益点文案，如“1件7.92折”、“每200减20”等     */
    PromotionDesc  *string `json:"promotion_desc,omitempty" `

    /*
        预热实际优惠金额（元）     */
    PromotionFee  *string `json:"promotion_fee,omitempty" `

    /*
        优惠开始时间     */
    PromotionStartTime  *string `json:"promotion_start_time,omitempty" `

    /*
        优惠结束时间     */
    PromotionEndTime  *string `json:"promotion_end_time,omitempty" `

}

func (s *TaobaoTbkScMaterialOptionalUpgradeFutureActivityPromotionPathMapData) SetPromotionTitle(v string) *TaobaoTbkScMaterialOptionalUpgradeFutureActivityPromotionPathMapData {
    s.PromotionTitle = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalUpgradeFutureActivityPromotionPathMapData) SetPromotionDesc(v string) *TaobaoTbkScMaterialOptionalUpgradeFutureActivityPromotionPathMapData {
    s.PromotionDesc = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalUpgradeFutureActivityPromotionPathMapData) SetPromotionFee(v string) *TaobaoTbkScMaterialOptionalUpgradeFutureActivityPromotionPathMapData {
    s.PromotionFee = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalUpgradeFutureActivityPromotionPathMapData) SetPromotionStartTime(v string) *TaobaoTbkScMaterialOptionalUpgradeFutureActivityPromotionPathMapData {
    s.PromotionStartTime = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalUpgradeFutureActivityPromotionPathMapData) SetPromotionEndTime(v string) *TaobaoTbkScMaterialOptionalUpgradeFutureActivityPromotionPathMapData {
    s.PromotionEndTime = &v
    return s
}

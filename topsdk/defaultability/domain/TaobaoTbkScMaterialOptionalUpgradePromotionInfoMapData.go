package domain


type TaobaoTbkScMaterialOptionalUpgradePromotionInfoMapData struct {
    /*
        到手价优惠路径列表     */
    FinalPromotionPathList  *[]TaobaoTbkScMaterialOptionalUpgradeFinalPromotionPathMapData `json:"final_promotion_path_list,omitempty" `

    /*
        促销信息-预估凑单价（元）。预估凑单叠加优惠后的商品单价     */
    PredictRoundingUpPrice  *string `json:"predict_rounding_up_price,omitempty" `

    /*
        促销信息-凑单价说明，描述凑单价的实现说明。如 “可凑单”或“需买X件”     */
    PredictRoundingUpPriceDesc  *string `json:"predict_rounding_up_price_desc,omitempty" `

    /*
        更多活动优惠     */
    MorePromotionList  *[]TaobaoTbkScMaterialOptionalUpgradeMorePromotionMapData `json:"more_promotion_list,omitempty" `

    /*
        商品信息-一口价通常显示为划线价     */
    ReservePrice  *string `json:"reserve_price,omitempty" `

    /*
        促销信息-销售价格，无促销时等于一口价，有促销时为促销价。若属于预售商品，付定金时间内，在线售卖价=预售价     */
    ZkFinalPrice  *string `json:"zk_final_price,omitempty" `

    /*
        促销信息-预估到手价(元)。若属于预售商品，付定金时间内，预估到手价价=定金+尾款的预估到手价     */
    FinalPromotionPrice  *string `json:"final_promotion_price,omitempty" `

    /*
        预热预估到手价（元）     */
    FutureActivityPromotionPrice  *string `json:"future_activity_promotion_price,omitempty" `

    /*
        预热到手价优惠路径列表     */
    FutureActivityPromotionPathList  *[]TaobaoTbkScMaterialOptionalUpgradeFutureActivityPromotionPathMapData `json:"future_activity_promotion_path_list,omitempty" `

    /*
        标签信息列表     */
    PromotionTagList  *[]TaobaoTbkScMaterialOptionalUpgradePromotionTagMapData `json:"promotion_tag_list,omitempty" `

}

func (s *TaobaoTbkScMaterialOptionalUpgradePromotionInfoMapData) SetFinalPromotionPathList(v []TaobaoTbkScMaterialOptionalUpgradeFinalPromotionPathMapData) *TaobaoTbkScMaterialOptionalUpgradePromotionInfoMapData {
    s.FinalPromotionPathList = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalUpgradePromotionInfoMapData) SetPredictRoundingUpPrice(v string) *TaobaoTbkScMaterialOptionalUpgradePromotionInfoMapData {
    s.PredictRoundingUpPrice = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalUpgradePromotionInfoMapData) SetPredictRoundingUpPriceDesc(v string) *TaobaoTbkScMaterialOptionalUpgradePromotionInfoMapData {
    s.PredictRoundingUpPriceDesc = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalUpgradePromotionInfoMapData) SetMorePromotionList(v []TaobaoTbkScMaterialOptionalUpgradeMorePromotionMapData) *TaobaoTbkScMaterialOptionalUpgradePromotionInfoMapData {
    s.MorePromotionList = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalUpgradePromotionInfoMapData) SetReservePrice(v string) *TaobaoTbkScMaterialOptionalUpgradePromotionInfoMapData {
    s.ReservePrice = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalUpgradePromotionInfoMapData) SetZkFinalPrice(v string) *TaobaoTbkScMaterialOptionalUpgradePromotionInfoMapData {
    s.ZkFinalPrice = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalUpgradePromotionInfoMapData) SetFinalPromotionPrice(v string) *TaobaoTbkScMaterialOptionalUpgradePromotionInfoMapData {
    s.FinalPromotionPrice = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalUpgradePromotionInfoMapData) SetFutureActivityPromotionPrice(v string) *TaobaoTbkScMaterialOptionalUpgradePromotionInfoMapData {
    s.FutureActivityPromotionPrice = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalUpgradePromotionInfoMapData) SetFutureActivityPromotionPathList(v []TaobaoTbkScMaterialOptionalUpgradeFutureActivityPromotionPathMapData) *TaobaoTbkScMaterialOptionalUpgradePromotionInfoMapData {
    s.FutureActivityPromotionPathList = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalUpgradePromotionInfoMapData) SetPromotionTagList(v []TaobaoTbkScMaterialOptionalUpgradePromotionTagMapData) *TaobaoTbkScMaterialOptionalUpgradePromotionInfoMapData {
    s.PromotionTagList = &v
    return s
}

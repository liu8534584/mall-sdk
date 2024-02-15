package domain


type TaobaoTbkScMaterialOptionalUpgradeMgcInfo struct {
    /*
        价格     */
    Price  *string `json:"price,omitempty" `

    /*
        价格描述     */
    PriceDesc  *string `json:"price_desc,omitempty" `

    /*
        文案     */
    PromotionSummary  *string `json:"promotion_summary,omitempty" `

    /*
        发布时间，13位毫秒时间戳     */
    PublishTime  *string `json:"publish_time,omitempty" `

    /*
        生效时间，实时线报为0，未来线报为13位毫秒时间戳     */
    ValidTime  *string `json:"valid_time,omitempty" `

}

func (s *TaobaoTbkScMaterialOptionalUpgradeMgcInfo) SetPrice(v string) *TaobaoTbkScMaterialOptionalUpgradeMgcInfo {
    s.Price = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalUpgradeMgcInfo) SetPriceDesc(v string) *TaobaoTbkScMaterialOptionalUpgradeMgcInfo {
    s.PriceDesc = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalUpgradeMgcInfo) SetPromotionSummary(v string) *TaobaoTbkScMaterialOptionalUpgradeMgcInfo {
    s.PromotionSummary = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalUpgradeMgcInfo) SetPublishTime(v string) *TaobaoTbkScMaterialOptionalUpgradeMgcInfo {
    s.PublishTime = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalUpgradeMgcInfo) SetValidTime(v string) *TaobaoTbkScMaterialOptionalUpgradeMgcInfo {
    s.ValidTime = &v
    return s
}

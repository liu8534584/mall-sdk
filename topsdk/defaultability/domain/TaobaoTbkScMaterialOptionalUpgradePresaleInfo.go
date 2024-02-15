package domain


type TaobaoTbkScMaterialOptionalUpgradePresaleInfo struct {
    /*
        预售商品-付定金开始时间（毫秒）     */
    PresaleStartTime  *int64 `json:"presale_start_time,omitempty" `

    /*
        预售商品-付定金结束时间（毫秒）     */
    PresaleEndTime  *int64 `json:"presale_end_time,omitempty" `

    /*
        预售商品-付尾款开始时间（毫秒）     */
    PresaleTailStartTime  *int64 `json:"presale_tail_start_time,omitempty" `

    /*
        预售商品-付尾款结束时间（毫秒）     */
    PresaleTailEndTime  *int64 `json:"presale_tail_end_time,omitempty" `

    /*
        预售商品-定金（元）     */
    PresaleDeposit  *string `json:"presale_deposit,omitempty" `

    /*
        预售商品-优惠信息     */
    PresaleDiscountFeeText  *string `json:"presale_discount_fee_text,omitempty" `

}

func (s *TaobaoTbkScMaterialOptionalUpgradePresaleInfo) SetPresaleStartTime(v int64) *TaobaoTbkScMaterialOptionalUpgradePresaleInfo {
    s.PresaleStartTime = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalUpgradePresaleInfo) SetPresaleEndTime(v int64) *TaobaoTbkScMaterialOptionalUpgradePresaleInfo {
    s.PresaleEndTime = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalUpgradePresaleInfo) SetPresaleTailStartTime(v int64) *TaobaoTbkScMaterialOptionalUpgradePresaleInfo {
    s.PresaleTailStartTime = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalUpgradePresaleInfo) SetPresaleTailEndTime(v int64) *TaobaoTbkScMaterialOptionalUpgradePresaleInfo {
    s.PresaleTailEndTime = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalUpgradePresaleInfo) SetPresaleDeposit(v string) *TaobaoTbkScMaterialOptionalUpgradePresaleInfo {
    s.PresaleDeposit = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalUpgradePresaleInfo) SetPresaleDiscountFeeText(v string) *TaobaoTbkScMaterialOptionalUpgradePresaleInfo {
    s.PresaleDiscountFeeText = &v
    return s
}

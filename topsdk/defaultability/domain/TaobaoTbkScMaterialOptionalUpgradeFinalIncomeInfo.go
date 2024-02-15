package domain


type TaobaoTbkScMaterialOptionalUpgradeFinalIncomeInfo struct {
    /*
        商品佣金比率     */
    CommissionRate  *string `json:"commission_rate,omitempty" `

    /*
        商品佣金金额     */
    CommissionAmount  *string `json:"commission_amount,omitempty" `

    /*
        补贴比率     */
    SubsidyRate  *string `json:"subsidy_rate,omitempty" `

    /*
        补贴金额     */
    SubsidyAmount  *string `json:"subsidy_amount,omitempty" `

    /*
        补贴上限；仅在单笔订单命中补贴上限时返回结果否则出参为空     */
    SubsidyUpperLimit  *string `json:"subsidy_upper_limit,omitempty" `

}

func (s *TaobaoTbkScMaterialOptionalUpgradeFinalIncomeInfo) SetCommissionRate(v string) *TaobaoTbkScMaterialOptionalUpgradeFinalIncomeInfo {
    s.CommissionRate = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalUpgradeFinalIncomeInfo) SetCommissionAmount(v string) *TaobaoTbkScMaterialOptionalUpgradeFinalIncomeInfo {
    s.CommissionAmount = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalUpgradeFinalIncomeInfo) SetSubsidyRate(v string) *TaobaoTbkScMaterialOptionalUpgradeFinalIncomeInfo {
    s.SubsidyRate = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalUpgradeFinalIncomeInfo) SetSubsidyAmount(v string) *TaobaoTbkScMaterialOptionalUpgradeFinalIncomeInfo {
    s.SubsidyAmount = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalUpgradeFinalIncomeInfo) SetSubsidyUpperLimit(v string) *TaobaoTbkScMaterialOptionalUpgradeFinalIncomeInfo {
    s.SubsidyUpperLimit = &v
    return s
}

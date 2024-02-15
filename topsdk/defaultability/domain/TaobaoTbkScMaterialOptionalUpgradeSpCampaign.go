package domain


type TaobaoTbkScMaterialOptionalUpgradeSpCampaign struct {
    /*
        定向计划活动ID     */
    SpCid  *string `json:"sp_cid,omitempty" `

    /*
        定向计划名称     */
    SpName  *string `json:"sp_name,omitempty" `

    /*
        定向佣金率     */
    SpRate  *string `json:"sp_rate,omitempty" `

    /*
        定向是否锁佣，0=不锁佣 1=锁佣     */
    SpLockStatus  *string `json:"sp_lock_status,omitempty" `

    /*
        定向计划申请链接     */
    SpApplyLink  *string `json:"sp_apply_link,omitempty" `

    /*
        定向计划是否可用 1-可用 0-不可用     */
    SpStatus  *string `json:"sp_status,omitempty" `

}

func (s *TaobaoTbkScMaterialOptionalUpgradeSpCampaign) SetSpCid(v string) *TaobaoTbkScMaterialOptionalUpgradeSpCampaign {
    s.SpCid = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalUpgradeSpCampaign) SetSpName(v string) *TaobaoTbkScMaterialOptionalUpgradeSpCampaign {
    s.SpName = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalUpgradeSpCampaign) SetSpRate(v string) *TaobaoTbkScMaterialOptionalUpgradeSpCampaign {
    s.SpRate = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalUpgradeSpCampaign) SetSpLockStatus(v string) *TaobaoTbkScMaterialOptionalUpgradeSpCampaign {
    s.SpLockStatus = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalUpgradeSpCampaign) SetSpApplyLink(v string) *TaobaoTbkScMaterialOptionalUpgradeSpCampaign {
    s.SpApplyLink = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalUpgradeSpCampaign) SetSpStatus(v string) *TaobaoTbkScMaterialOptionalUpgradeSpCampaign {
    s.SpStatus = &v
    return s
}

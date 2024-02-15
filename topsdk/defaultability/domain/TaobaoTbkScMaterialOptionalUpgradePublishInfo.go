package domain


type TaobaoTbkScMaterialOptionalUpgradePublishInfo struct {
    /*
        商品信息-收入比率(%)；商品佣金比率+补贴比率     */
    IncomeRate  *string `json:"income_rate,omitempty" `

    /*
        前N件佣金信息-前N件佣金生效或预热时透出以下字段     */
    TopnInfo  *TaobaoTbkScMaterialOptionalUpgradeTopNInfoDTO `json:"topn_info,omitempty" `

    /*
        链接-宝贝推广链接     */
    ClickUrl  *string `json:"click_url,omitempty" `

    /*
        链接-宝贝+券二合一页面链接     */
    CouponShareUrl  *string `json:"coupon_share_url,omitempty" `

    /*
        额外奖励活动类型，如果一个商品有多个奖励类型，返回结果使用空格分割，0=预售单单奖励，1=618超级U选单单补     */
    CpaRewardType  *string `json:"cpa_reward_type,omitempty" `

    /*
        额外奖励活动金额，活动奖励金额的类型与cpa_reward_type字段对应，如果一个商品有多个奖励类型，返回结果使用空格分割     */
    CpaRewardAmount  *string `json:"cpa_reward_amount,omitempty" `

    /*
        预热活动到手价对应的佣金比率     */
    FutureActivityCommissionRate  *string `json:"future_activity_commission_rate,omitempty" `

    /*
        预热价活动开始时间     */
    FutureActivityTime  *string `json:"future_activity_time,omitempty" `

    /*
        定向计划集合     */
    SpCampaignList  *[]TaobaoTbkScMaterialOptionalUpgradeSpCampaign `json:"sp_campaign_list,omitempty" `

    /*
        榜单url     */
    RankPageUrl  *string `json:"rank_page_url,omitempty" `

    /*
        推广信息-商品信息-佣金类型。MKT表示营销计划，SP表示定向计划，COMMON表示通用计划     */
    CommissionType  *string `json:"commission_type,omitempty" `

    /*
        商品佣金信息     */
    IncomeInfo  *TaobaoTbkScMaterialOptionalUpgradeFinalIncomeInfo `json:"income_info,omitempty" `

}

func (s *TaobaoTbkScMaterialOptionalUpgradePublishInfo) SetIncomeRate(v string) *TaobaoTbkScMaterialOptionalUpgradePublishInfo {
    s.IncomeRate = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalUpgradePublishInfo) SetTopnInfo(v TaobaoTbkScMaterialOptionalUpgradeTopNInfoDTO) *TaobaoTbkScMaterialOptionalUpgradePublishInfo {
    s.TopnInfo = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalUpgradePublishInfo) SetClickUrl(v string) *TaobaoTbkScMaterialOptionalUpgradePublishInfo {
    s.ClickUrl = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalUpgradePublishInfo) SetCouponShareUrl(v string) *TaobaoTbkScMaterialOptionalUpgradePublishInfo {
    s.CouponShareUrl = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalUpgradePublishInfo) SetCpaRewardType(v string) *TaobaoTbkScMaterialOptionalUpgradePublishInfo {
    s.CpaRewardType = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalUpgradePublishInfo) SetCpaRewardAmount(v string) *TaobaoTbkScMaterialOptionalUpgradePublishInfo {
    s.CpaRewardAmount = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalUpgradePublishInfo) SetFutureActivityCommissionRate(v string) *TaobaoTbkScMaterialOptionalUpgradePublishInfo {
    s.FutureActivityCommissionRate = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalUpgradePublishInfo) SetFutureActivityTime(v string) *TaobaoTbkScMaterialOptionalUpgradePublishInfo {
    s.FutureActivityTime = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalUpgradePublishInfo) SetSpCampaignList(v []TaobaoTbkScMaterialOptionalUpgradeSpCampaign) *TaobaoTbkScMaterialOptionalUpgradePublishInfo {
    s.SpCampaignList = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalUpgradePublishInfo) SetRankPageUrl(v string) *TaobaoTbkScMaterialOptionalUpgradePublishInfo {
    s.RankPageUrl = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalUpgradePublishInfo) SetCommissionType(v string) *TaobaoTbkScMaterialOptionalUpgradePublishInfo {
    s.CommissionType = &v
    return s
}
func (s *TaobaoTbkScMaterialOptionalUpgradePublishInfo) SetIncomeInfo(v TaobaoTbkScMaterialOptionalUpgradeFinalIncomeInfo) *TaobaoTbkScMaterialOptionalUpgradePublishInfo {
    s.IncomeInfo = &v
    return s
}

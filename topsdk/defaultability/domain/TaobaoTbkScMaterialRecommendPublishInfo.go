package domain


type TaobaoTbkScMaterialRecommendPublishInfo struct {
    /*
        商品信息-佣金比率(%)     */
    CommissionRate  *string `json:"commission_rate,omitempty" `

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
        定向计划集合     */
    SpCampaignList  *[]TaobaoTbkScMaterialRecommendSpCampaign `json:"sp_campaign_list,omitempty" `

    /*
        预热活动到手价对应的佣金比率     */
    FutureActivityCommissionRate  *string `json:"future_activity_commission_rate,omitempty" `

    /*
        预热价活动开始时间     */
    FutureActivityTime  *string `json:"future_activity_time,omitempty" `

}

func (s *TaobaoTbkScMaterialRecommendPublishInfo) SetCommissionRate(v string) *TaobaoTbkScMaterialRecommendPublishInfo {
    s.CommissionRate = &v
    return s
}
func (s *TaobaoTbkScMaterialRecommendPublishInfo) SetClickUrl(v string) *TaobaoTbkScMaterialRecommendPublishInfo {
    s.ClickUrl = &v
    return s
}
func (s *TaobaoTbkScMaterialRecommendPublishInfo) SetCouponShareUrl(v string) *TaobaoTbkScMaterialRecommendPublishInfo {
    s.CouponShareUrl = &v
    return s
}
func (s *TaobaoTbkScMaterialRecommendPublishInfo) SetCpaRewardType(v string) *TaobaoTbkScMaterialRecommendPublishInfo {
    s.CpaRewardType = &v
    return s
}
func (s *TaobaoTbkScMaterialRecommendPublishInfo) SetCpaRewardAmount(v string) *TaobaoTbkScMaterialRecommendPublishInfo {
    s.CpaRewardAmount = &v
    return s
}
func (s *TaobaoTbkScMaterialRecommendPublishInfo) SetSpCampaignList(v []TaobaoTbkScMaterialRecommendSpCampaign) *TaobaoTbkScMaterialRecommendPublishInfo {
    s.SpCampaignList = &v
    return s
}
func (s *TaobaoTbkScMaterialRecommendPublishInfo) SetFutureActivityCommissionRate(v string) *TaobaoTbkScMaterialRecommendPublishInfo {
    s.FutureActivityCommissionRate = &v
    return s
}
func (s *TaobaoTbkScMaterialRecommendPublishInfo) SetFutureActivityTime(v string) *TaobaoTbkScMaterialRecommendPublishInfo {
    s.FutureActivityTime = &v
    return s
}

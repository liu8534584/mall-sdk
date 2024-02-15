package domain


type AlibabaAlscUnionKbBbtItemDetailGetPurchaseLimit struct {
    /*
        商品限购-每人每天限购（-1表示不限购）     */
    ItemDailyLimitPerUser  *int64 `json:"item_daily_limit_per_user,omitempty" `

    /*
        商品限购-每人终身限购（-1表示不限购）     */
    ItemLimitPerUser  *int64 `json:"item_limit_per_user,omitempty" `

    /*
        商品限购-每人每订单限购（-1表示不限购）     */
    ItemLimitPerUserOrder  *int64 `json:"item_limit_per_user_order,omitempty" `

    /*
        活动限购-每人每天限购（-1表示不限购）     */
    ActivityDailyLimitPerUser  *int64 `json:"activity_daily_limit_per_user,omitempty" `

    /*
        活动限购-每人每活动限购（-1表示不限购）     */
    ActivityLimitPerUser  *int64 `json:"activity_limit_per_user,omitempty" `

}

func (s *AlibabaAlscUnionKbBbtItemDetailGetPurchaseLimit) SetItemDailyLimitPerUser(v int64) *AlibabaAlscUnionKbBbtItemDetailGetPurchaseLimit {
    s.ItemDailyLimitPerUser = &v
    return s
}
func (s *AlibabaAlscUnionKbBbtItemDetailGetPurchaseLimit) SetItemLimitPerUser(v int64) *AlibabaAlscUnionKbBbtItemDetailGetPurchaseLimit {
    s.ItemLimitPerUser = &v
    return s
}
func (s *AlibabaAlscUnionKbBbtItemDetailGetPurchaseLimit) SetItemLimitPerUserOrder(v int64) *AlibabaAlscUnionKbBbtItemDetailGetPurchaseLimit {
    s.ItemLimitPerUserOrder = &v
    return s
}
func (s *AlibabaAlscUnionKbBbtItemDetailGetPurchaseLimit) SetActivityDailyLimitPerUser(v int64) *AlibabaAlscUnionKbBbtItemDetailGetPurchaseLimit {
    s.ActivityDailyLimitPerUser = &v
    return s
}
func (s *AlibabaAlscUnionKbBbtItemDetailGetPurchaseLimit) SetActivityLimitPerUser(v int64) *AlibabaAlscUnionKbBbtItemDetailGetPurchaseLimit {
    s.ActivityLimitPerUser = &v
    return s
}

package domain

type TaobaoTbkScOrderDetailsGetServiceFeeDTO struct {
	/*
	   专项服务费率（字段已废弃）     */
	ShareRelativeRate *string `json:"share_relative_rate,omitempty" `

	/*
	   结算专项服务费（字段已废弃）     */
	ShareFee *string `json:"share_fee,omitempty" `

	/*
	   预估专项服务费（字段已废弃）     */
	SharePreFee *string `json:"share_pre_fee,omitempty" `

	/*
	   专项服务费来源，122-渠道（字段已废弃）     */
	TkShareRoleType *int64 `json:"tk_share_role_type,omitempty" `
}

func (s *TaobaoTbkScOrderDetailsGetServiceFeeDTO) SetShareRelativeRate(v string) *TaobaoTbkScOrderDetailsGetServiceFeeDTO {
	s.ShareRelativeRate = &v
	return s
}
func (s *TaobaoTbkScOrderDetailsGetServiceFeeDTO) SetShareFee(v string) *TaobaoTbkScOrderDetailsGetServiceFeeDTO {
	s.ShareFee = &v
	return s
}
func (s *TaobaoTbkScOrderDetailsGetServiceFeeDTO) SetSharePreFee(v string) *TaobaoTbkScOrderDetailsGetServiceFeeDTO {
	s.SharePreFee = &v
	return s
}
func (s *TaobaoTbkScOrderDetailsGetServiceFeeDTO) SetTkShareRoleType(v int64) *TaobaoTbkScOrderDetailsGetServiceFeeDTO {
	s.TkShareRoleType = &v
	return s
}

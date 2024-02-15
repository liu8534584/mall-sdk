package domain

type TaobaoTbkScOrderDetailsGetSubsidyDetailDTO struct {
	/*
	   该笔订单包含的补贴类型     */
	SubsidyType *string `json:"subsidy_type,omitempty" `

	/*
	   补贴比率     */
	SubsidyRate *float64 `json:"subsidy_rate,omitempty" `

	/*
	   对应补贴类型的补贴金额     */
	SubsidyFee *float64 `json:"subsidy_fee,omitempty" `

	/*
	   单笔订单补贴上限。对应补贴类型的单笔订单补贴上限     */
	SubsidyUpperLimit *float64 `json:"subsidy_upper_limit,omitempty" `

	/*
	   补贴分成比率     */
	SubsidyShareRate *float64 `json:"subsidy_share_rate,omitempty" `
}

func (s *TaobaoTbkScOrderDetailsGetSubsidyDetailDTO) SetSubsidyType(v string) *TaobaoTbkScOrderDetailsGetSubsidyDetailDTO {
	s.SubsidyType = &v
	return s
}
func (s *TaobaoTbkScOrderDetailsGetSubsidyDetailDTO) SetSubsidyRate(v float64) *TaobaoTbkScOrderDetailsGetSubsidyDetailDTO {
	s.SubsidyRate = &v
	return s
}
func (s *TaobaoTbkScOrderDetailsGetSubsidyDetailDTO) SetSubsidyFee(v float64) *TaobaoTbkScOrderDetailsGetSubsidyDetailDTO {
	s.SubsidyFee = &v
	return s
}
func (s *TaobaoTbkScOrderDetailsGetSubsidyDetailDTO) SetSubsidyUpperLimit(v float64) *TaobaoTbkScOrderDetailsGetSubsidyDetailDTO {
	s.SubsidyUpperLimit = &v
	return s
}
func (s *TaobaoTbkScOrderDetailsGetSubsidyDetailDTO) SetSubsidyShareRate(v float64) *TaobaoTbkScOrderDetailsGetSubsidyDetailDTO {
	s.SubsidyShareRate = &v
	return s
}

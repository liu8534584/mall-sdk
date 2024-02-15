package domain

type TaobaoTbkScOrderDetailsGetAlimmShareInfoDTO struct {
	/*
	   技术服务费比率     */
	AlimmTechServiceRate *float64 `json:"alimm_tech_service_rate,omitempty" `

	/*
	   预估技术服务费     */
	AlimmTechServicePreFee *float64 `json:"alimm_tech_service_pre_fee,omitempty" `

	/*
	   结算技术服务费     */
	AlimmTechServiceFee *float64 `json:"alimm_tech_service_fee,omitempty" `

	/*
	   渠道专项服务费比率     */
	AlimmAgentServiceRate *float64 `json:"alimm_agent_service_rate,omitempty" `

	/*
	   预估渠道专项服务费     */
	AlimmAgentServicePreFee *float64 `json:"alimm_agent_service_pre_fee,omitempty" `

	/*
	   结算渠道专项服务费     */
	AlimmAgentServiceFee *float64 `json:"alimm_agent_service_fee,omitempty" `
}

func (s *TaobaoTbkScOrderDetailsGetAlimmShareInfoDTO) SetAlimmTechServiceRate(v float64) *TaobaoTbkScOrderDetailsGetAlimmShareInfoDTO {
	s.AlimmTechServiceRate = &v
	return s
}
func (s *TaobaoTbkScOrderDetailsGetAlimmShareInfoDTO) SetAlimmTechServicePreFee(v float64) *TaobaoTbkScOrderDetailsGetAlimmShareInfoDTO {
	s.AlimmTechServicePreFee = &v
	return s
}
func (s *TaobaoTbkScOrderDetailsGetAlimmShareInfoDTO) SetAlimmTechServiceFee(v float64) *TaobaoTbkScOrderDetailsGetAlimmShareInfoDTO {
	s.AlimmTechServiceFee = &v
	return s
}
func (s *TaobaoTbkScOrderDetailsGetAlimmShareInfoDTO) SetAlimmAgentServiceRate(v float64) *TaobaoTbkScOrderDetailsGetAlimmShareInfoDTO {
	s.AlimmAgentServiceRate = &v
	return s
}
func (s *TaobaoTbkScOrderDetailsGetAlimmShareInfoDTO) SetAlimmAgentServicePreFee(v float64) *TaobaoTbkScOrderDetailsGetAlimmShareInfoDTO {
	s.AlimmAgentServicePreFee = &v
	return s
}
func (s *TaobaoTbkScOrderDetailsGetAlimmShareInfoDTO) SetAlimmAgentServiceFee(v float64) *TaobaoTbkScOrderDetailsGetAlimmShareInfoDTO {
	s.AlimmAgentServiceFee = &v
	return s
}

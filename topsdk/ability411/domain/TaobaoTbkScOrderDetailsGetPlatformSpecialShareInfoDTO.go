package domain

type TaobaoTbkScOrderDetailsGetPlatformSpecialShareInfoDTO struct {
	/*
	   内容专项服务费比率     */
	ContentTechServiceRate *float64 `json:"content_tech_service_rate,omitempty" `

	/*
	   预估内容专项服务费     */
	ContentTechServicePreFee *float64 `json:"content_tech_service_pre_fee,omitempty" `

	/*
	   结算内容专项服务费     */
	ContentTechServiceFee *float64 `json:"content_tech_service_fee,omitempty" `

	/*
	   流量专项服务费比率（默认无，限定开放）     */
	TrafficTechServiceRate *float64 `json:"traffic_tech_service_rate,omitempty" `

	/*
	   预估流量专项服务费（默认无，限定开放）     */
	TrafficTechServicePreFee *float64 `json:"traffic_tech_service_pre_fee,omitempty" `

	/*
	   结算流量专项服务费（默认无，限定开放）     */
	TrafficTechServiceFee *float64 `json:"traffic_tech_service_fee,omitempty" `
}

func (s *TaobaoTbkScOrderDetailsGetPlatformSpecialShareInfoDTO) SetContentTechServiceRate(v float64) *TaobaoTbkScOrderDetailsGetPlatformSpecialShareInfoDTO {
	s.ContentTechServiceRate = &v
	return s
}
func (s *TaobaoTbkScOrderDetailsGetPlatformSpecialShareInfoDTO) SetContentTechServicePreFee(v float64) *TaobaoTbkScOrderDetailsGetPlatformSpecialShareInfoDTO {
	s.ContentTechServicePreFee = &v
	return s
}
func (s *TaobaoTbkScOrderDetailsGetPlatformSpecialShareInfoDTO) SetContentTechServiceFee(v float64) *TaobaoTbkScOrderDetailsGetPlatformSpecialShareInfoDTO {
	s.ContentTechServiceFee = &v
	return s
}
func (s *TaobaoTbkScOrderDetailsGetPlatformSpecialShareInfoDTO) SetTrafficTechServiceRate(v float64) *TaobaoTbkScOrderDetailsGetPlatformSpecialShareInfoDTO {
	s.TrafficTechServiceRate = &v
	return s
}
func (s *TaobaoTbkScOrderDetailsGetPlatformSpecialShareInfoDTO) SetTrafficTechServicePreFee(v float64) *TaobaoTbkScOrderDetailsGetPlatformSpecialShareInfoDTO {
	s.TrafficTechServicePreFee = &v
	return s
}
func (s *TaobaoTbkScOrderDetailsGetPlatformSpecialShareInfoDTO) SetTrafficTechServiceFee(v float64) *TaobaoTbkScOrderDetailsGetPlatformSpecialShareInfoDTO {
	s.TrafficTechServiceFee = &v
	return s
}

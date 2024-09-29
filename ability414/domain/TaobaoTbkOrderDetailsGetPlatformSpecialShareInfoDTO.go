package domain


type TaobaoTbkOrderDetailsGetPlatformSpecialShareInfoDTO struct {
    /*
        内容专项服务费比率     */
    ContentTechServiceRate  *string `json:"content_tech_service_rate,omitempty" `

    /*
        预估内容专项服务费     */
    ContentTechServicePreFee  *string `json:"content_tech_service_pre_fee,omitempty" `

    /*
        结算内容专项服务费     */
    ContentTechServiceFee  *string `json:"content_tech_service_fee,omitempty" `

    /*
        流量专项服务费比率（默认无，限定开放）     */
    TrafficTechServiceRate  *string `json:"traffic_tech_service_rate,omitempty" `

    /*
        预估流量专项服务费（默认无，限定开放）     */
    TrafficTechServicePreFee  *string `json:"traffic_tech_service_pre_fee,omitempty" `

    /*
        结算流量专项服务费（默认无，限定开放）     */
    TrafficTechServiceFee  *string `json:"traffic_tech_service_fee,omitempty" `

}

func (s *TaobaoTbkOrderDetailsGetPlatformSpecialShareInfoDTO) SetContentTechServiceRate(v string) *TaobaoTbkOrderDetailsGetPlatformSpecialShareInfoDTO {
    s.ContentTechServiceRate = &v
    return s
}
func (s *TaobaoTbkOrderDetailsGetPlatformSpecialShareInfoDTO) SetContentTechServicePreFee(v string) *TaobaoTbkOrderDetailsGetPlatformSpecialShareInfoDTO {
    s.ContentTechServicePreFee = &v
    return s
}
func (s *TaobaoTbkOrderDetailsGetPlatformSpecialShareInfoDTO) SetContentTechServiceFee(v string) *TaobaoTbkOrderDetailsGetPlatformSpecialShareInfoDTO {
    s.ContentTechServiceFee = &v
    return s
}
func (s *TaobaoTbkOrderDetailsGetPlatformSpecialShareInfoDTO) SetTrafficTechServiceRate(v string) *TaobaoTbkOrderDetailsGetPlatformSpecialShareInfoDTO {
    s.TrafficTechServiceRate = &v
    return s
}
func (s *TaobaoTbkOrderDetailsGetPlatformSpecialShareInfoDTO) SetTrafficTechServicePreFee(v string) *TaobaoTbkOrderDetailsGetPlatformSpecialShareInfoDTO {
    s.TrafficTechServicePreFee = &v
    return s
}
func (s *TaobaoTbkOrderDetailsGetPlatformSpecialShareInfoDTO) SetTrafficTechServiceFee(v string) *TaobaoTbkOrderDetailsGetPlatformSpecialShareInfoDTO {
    s.TrafficTechServiceFee = &v
    return s
}

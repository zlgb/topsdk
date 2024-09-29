package domain


type TaobaoTbkOrderDetailsGetAlimmShareInfoDTO struct {
    /*
        技术服务费比率     */
    AlimmTechServiceRate  *string `json:"alimm_tech_service_rate,omitempty" `

    /*
        预估技术服务费     */
    AlimmTechServicePreFee  *string `json:"alimm_tech_service_pre_fee,omitempty" `

    /*
        结算技术服务费     */
    AlimmTechServiceFee  *string `json:"alimm_tech_service_fee,omitempty" `

    /*
        渠道专项服务费比率     */
    AlimmAgentServiceRate  *string `json:"alimm_agent_service_rate,omitempty" `

    /*
        预估渠道专项服务费     */
    AlimmAgentServicePreFee  *string `json:"alimm_agent_service_pre_fee,omitempty" `

    /*
        结算渠道专项服务费     */
    AlimmAgentServiceFee  *string `json:"alimm_agent_service_fee,omitempty" `

}

func (s *TaobaoTbkOrderDetailsGetAlimmShareInfoDTO) SetAlimmTechServiceRate(v string) *TaobaoTbkOrderDetailsGetAlimmShareInfoDTO {
    s.AlimmTechServiceRate = &v
    return s
}
func (s *TaobaoTbkOrderDetailsGetAlimmShareInfoDTO) SetAlimmTechServicePreFee(v string) *TaobaoTbkOrderDetailsGetAlimmShareInfoDTO {
    s.AlimmTechServicePreFee = &v
    return s
}
func (s *TaobaoTbkOrderDetailsGetAlimmShareInfoDTO) SetAlimmTechServiceFee(v string) *TaobaoTbkOrderDetailsGetAlimmShareInfoDTO {
    s.AlimmTechServiceFee = &v
    return s
}
func (s *TaobaoTbkOrderDetailsGetAlimmShareInfoDTO) SetAlimmAgentServiceRate(v string) *TaobaoTbkOrderDetailsGetAlimmShareInfoDTO {
    s.AlimmAgentServiceRate = &v
    return s
}
func (s *TaobaoTbkOrderDetailsGetAlimmShareInfoDTO) SetAlimmAgentServicePreFee(v string) *TaobaoTbkOrderDetailsGetAlimmShareInfoDTO {
    s.AlimmAgentServicePreFee = &v
    return s
}
func (s *TaobaoTbkOrderDetailsGetAlimmShareInfoDTO) SetAlimmAgentServiceFee(v string) *TaobaoTbkOrderDetailsGetAlimmShareInfoDTO {
    s.AlimmAgentServiceFee = &v
    return s
}

package domain


type TaobaoTbkDgGeneralLinkConvertPromotionInfoDTO struct {
    /*
        商品收入比率(%)：商品佣金比率+补贴比率。同物料类接口income_rate     */
    CommissionRate  *string `json:"commission_rate,omitempty" `

    /*
        佣金类型。MKT表示营销计划，SP表示定向计划，COMMON表示通用计划，ZX表示自选计划     */
    CommissionType  *string `json:"commission_type,omitempty" `

}

func (s *TaobaoTbkDgGeneralLinkConvertPromotionInfoDTO) SetCommissionRate(v string) *TaobaoTbkDgGeneralLinkConvertPromotionInfoDTO {
    s.CommissionRate = &v
    return s
}
func (s *TaobaoTbkDgGeneralLinkConvertPromotionInfoDTO) SetCommissionType(v string) *TaobaoTbkDgGeneralLinkConvertPromotionInfoDTO {
    s.CommissionType = &v
    return s
}

package domain


type TaobaoTbkDgMaterialTemporaryOptionalStepRateDTO struct {
    /*
        前N件剩余库存     */
    TopnQuantity  *int64 `json:"topn_quantity,omitempty" `

    /*
        前N件初始总库存     */
    TopnTotalCount  *int64 `json:"topn_total_count,omitempty" `

    /*
        前N件佣金结束时间     */
    TopnEndTime  *string `json:"topn_end_time,omitempty" `

    /*
        前N件佣金开始时间     */
    TopnStartTime  *string `json:"topn_start_time,omitempty" `

    /*
        前N件佣金率     */
    TopnRate  *string `json:"topn_rate,omitempty" `

}

func (s *TaobaoTbkDgMaterialTemporaryOptionalStepRateDTO) SetTopnQuantity(v int64) *TaobaoTbkDgMaterialTemporaryOptionalStepRateDTO {
    s.TopnQuantity = &v
    return s
}
func (s *TaobaoTbkDgMaterialTemporaryOptionalStepRateDTO) SetTopnTotalCount(v int64) *TaobaoTbkDgMaterialTemporaryOptionalStepRateDTO {
    s.TopnTotalCount = &v
    return s
}
func (s *TaobaoTbkDgMaterialTemporaryOptionalStepRateDTO) SetTopnEndTime(v string) *TaobaoTbkDgMaterialTemporaryOptionalStepRateDTO {
    s.TopnEndTime = &v
    return s
}
func (s *TaobaoTbkDgMaterialTemporaryOptionalStepRateDTO) SetTopnStartTime(v string) *TaobaoTbkDgMaterialTemporaryOptionalStepRateDTO {
    s.TopnStartTime = &v
    return s
}
func (s *TaobaoTbkDgMaterialTemporaryOptionalStepRateDTO) SetTopnRate(v string) *TaobaoTbkDgMaterialTemporaryOptionalStepRateDTO {
    s.TopnRate = &v
    return s
}

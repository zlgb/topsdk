package domain


type TaobaoTbkDgMaterialOptionalUpgradeTopNInfoDTO struct {
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

func (s *TaobaoTbkDgMaterialOptionalUpgradeTopNInfoDTO) SetTopnQuantity(v int64) *TaobaoTbkDgMaterialOptionalUpgradeTopNInfoDTO {
    s.TopnQuantity = &v
    return s
}
func (s *TaobaoTbkDgMaterialOptionalUpgradeTopNInfoDTO) SetTopnTotalCount(v int64) *TaobaoTbkDgMaterialOptionalUpgradeTopNInfoDTO {
    s.TopnTotalCount = &v
    return s
}
func (s *TaobaoTbkDgMaterialOptionalUpgradeTopNInfoDTO) SetTopnEndTime(v string) *TaobaoTbkDgMaterialOptionalUpgradeTopNInfoDTO {
    s.TopnEndTime = &v
    return s
}
func (s *TaobaoTbkDgMaterialOptionalUpgradeTopNInfoDTO) SetTopnStartTime(v string) *TaobaoTbkDgMaterialOptionalUpgradeTopNInfoDTO {
    s.TopnStartTime = &v
    return s
}
func (s *TaobaoTbkDgMaterialOptionalUpgradeTopNInfoDTO) SetTopnRate(v string) *TaobaoTbkDgMaterialOptionalUpgradeTopNInfoDTO {
    s.TopnRate = &v
    return s
}

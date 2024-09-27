package domain


type TaobaoTbkDgMaterialRecommendTopNInfoDTO struct {
    /*
        前N件剩余库存     */
    TopnQuantity  *int64 `json:"topn_quantity,omitempty" `

    /*
        前N件初始总库存     */
    TopnTotalCount  *int64 `json:"topn_total_count,omitempty" `

    /*
        前N件佣金开始时间     */
    TopnStartTime  *string `json:"topn_start_time,omitempty" `

    /*
        前N件佣金结束时间     */
    TopnEndTime  *string `json:"topn_end_time,omitempty" `

    /*
        前N件佣金率     */
    TopnRate  *string `json:"topn_rate,omitempty" `

}

func (s *TaobaoTbkDgMaterialRecommendTopNInfoDTO) SetTopnQuantity(v int64) *TaobaoTbkDgMaterialRecommendTopNInfoDTO {
    s.TopnQuantity = &v
    return s
}
func (s *TaobaoTbkDgMaterialRecommendTopNInfoDTO) SetTopnTotalCount(v int64) *TaobaoTbkDgMaterialRecommendTopNInfoDTO {
    s.TopnTotalCount = &v
    return s
}
func (s *TaobaoTbkDgMaterialRecommendTopNInfoDTO) SetTopnStartTime(v string) *TaobaoTbkDgMaterialRecommendTopNInfoDTO {
    s.TopnStartTime = &v
    return s
}
func (s *TaobaoTbkDgMaterialRecommendTopNInfoDTO) SetTopnEndTime(v string) *TaobaoTbkDgMaterialRecommendTopNInfoDTO {
    s.TopnEndTime = &v
    return s
}
func (s *TaobaoTbkDgMaterialRecommendTopNInfoDTO) SetTopnRate(v string) *TaobaoTbkDgMaterialRecommendTopNInfoDTO {
    s.TopnRate = &v
    return s
}

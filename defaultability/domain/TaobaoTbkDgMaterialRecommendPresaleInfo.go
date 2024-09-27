package domain


type TaobaoTbkDgMaterialRecommendPresaleInfo struct {
    /*
        预售商品-付定金开始时间（毫秒）     */
    PresaleStartTime  *int64 `json:"presale_start_time,omitempty" `

    /*
        预售商品-付定金结束时间（毫秒）     */
    PresaleEndTime  *int64 `json:"presale_end_time,omitempty" `

    /*
        预售商品-付尾款开始时间（毫秒）     */
    PresaleTailStartTime  *int64 `json:"presale_tail_start_time,omitempty" `

    /*
        预售商品-付尾款结束时间（毫秒）     */
    PresaleTailEndTime  *int64 `json:"presale_tail_end_time,omitempty" `

    /*
        预售商品-定金（元）     */
    PresaleDeposit  *string `json:"presale_deposit,omitempty" `

    /*
        预售商品-优惠信息     */
    PresaleDiscountFeeText  *string `json:"presale_discount_fee_text,omitempty" `

}

func (s *TaobaoTbkDgMaterialRecommendPresaleInfo) SetPresaleStartTime(v int64) *TaobaoTbkDgMaterialRecommendPresaleInfo {
    s.PresaleStartTime = &v
    return s
}
func (s *TaobaoTbkDgMaterialRecommendPresaleInfo) SetPresaleEndTime(v int64) *TaobaoTbkDgMaterialRecommendPresaleInfo {
    s.PresaleEndTime = &v
    return s
}
func (s *TaobaoTbkDgMaterialRecommendPresaleInfo) SetPresaleTailStartTime(v int64) *TaobaoTbkDgMaterialRecommendPresaleInfo {
    s.PresaleTailStartTime = &v
    return s
}
func (s *TaobaoTbkDgMaterialRecommendPresaleInfo) SetPresaleTailEndTime(v int64) *TaobaoTbkDgMaterialRecommendPresaleInfo {
    s.PresaleTailEndTime = &v
    return s
}
func (s *TaobaoTbkDgMaterialRecommendPresaleInfo) SetPresaleDeposit(v string) *TaobaoTbkDgMaterialRecommendPresaleInfo {
    s.PresaleDeposit = &v
    return s
}
func (s *TaobaoTbkDgMaterialRecommendPresaleInfo) SetPresaleDiscountFeeText(v string) *TaobaoTbkDgMaterialRecommendPresaleInfo {
    s.PresaleDiscountFeeText = &v
    return s
}

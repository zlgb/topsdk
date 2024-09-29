package domain


type TaobaoTbkDgGeneralLinkConvertCouponInfoDTO struct {
    /*
        优惠券结束时间     */
    CouponEndTime  *string `json:"coupon_end_time,omitempty" `

    /*
        券ID     */
    ActivityId  *string `json:"activity_id,omitempty" `

    /*
        优惠券剩余数量     */
    CouponRemainCount  *int64 `json:"coupon_remain_count,omitempty" `

    /*
        券面额     */
    CouponAmount  *string `json:"coupon_amount,omitempty" `

    /*
        优惠券开始时间     */
    CouponStartTime  *string `json:"coupon_start_time,omitempty" `

    /*
        优惠券信息，满XX元减XX元     */
    CouponDesc  *string `json:"coupon_desc,omitempty" `

    /*
         0-店铺 1-单品     */
    CouponType  *int64 `json:"coupon_type,omitempty" `

}

func (s *TaobaoTbkDgGeneralLinkConvertCouponInfoDTO) SetCouponEndTime(v string) *TaobaoTbkDgGeneralLinkConvertCouponInfoDTO {
    s.CouponEndTime = &v
    return s
}
func (s *TaobaoTbkDgGeneralLinkConvertCouponInfoDTO) SetActivityId(v string) *TaobaoTbkDgGeneralLinkConvertCouponInfoDTO {
    s.ActivityId = &v
    return s
}
func (s *TaobaoTbkDgGeneralLinkConvertCouponInfoDTO) SetCouponRemainCount(v int64) *TaobaoTbkDgGeneralLinkConvertCouponInfoDTO {
    s.CouponRemainCount = &v
    return s
}
func (s *TaobaoTbkDgGeneralLinkConvertCouponInfoDTO) SetCouponAmount(v string) *TaobaoTbkDgGeneralLinkConvertCouponInfoDTO {
    s.CouponAmount = &v
    return s
}
func (s *TaobaoTbkDgGeneralLinkConvertCouponInfoDTO) SetCouponStartTime(v string) *TaobaoTbkDgGeneralLinkConvertCouponInfoDTO {
    s.CouponStartTime = &v
    return s
}
func (s *TaobaoTbkDgGeneralLinkConvertCouponInfoDTO) SetCouponDesc(v string) *TaobaoTbkDgGeneralLinkConvertCouponInfoDTO {
    s.CouponDesc = &v
    return s
}
func (s *TaobaoTbkDgGeneralLinkConvertCouponInfoDTO) SetCouponType(v int64) *TaobaoTbkDgGeneralLinkConvertCouponInfoDTO {
    s.CouponType = &v
    return s
}

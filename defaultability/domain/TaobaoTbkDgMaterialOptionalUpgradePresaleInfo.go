package domain


type TaobaoTbkDgMaterialOptionalUpgradePresaleInfo struct {
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

func (s *TaobaoTbkDgMaterialOptionalUpgradePresaleInfo) SetPresaleStartTime(v int64) *TaobaoTbkDgMaterialOptionalUpgradePresaleInfo {
    s.PresaleStartTime = &v
    return s
}
func (s *TaobaoTbkDgMaterialOptionalUpgradePresaleInfo) SetPresaleEndTime(v int64) *TaobaoTbkDgMaterialOptionalUpgradePresaleInfo {
    s.PresaleEndTime = &v
    return s
}
func (s *TaobaoTbkDgMaterialOptionalUpgradePresaleInfo) SetPresaleTailStartTime(v int64) *TaobaoTbkDgMaterialOptionalUpgradePresaleInfo {
    s.PresaleTailStartTime = &v
    return s
}
func (s *TaobaoTbkDgMaterialOptionalUpgradePresaleInfo) SetPresaleTailEndTime(v int64) *TaobaoTbkDgMaterialOptionalUpgradePresaleInfo {
    s.PresaleTailEndTime = &v
    return s
}
func (s *TaobaoTbkDgMaterialOptionalUpgradePresaleInfo) SetPresaleDeposit(v string) *TaobaoTbkDgMaterialOptionalUpgradePresaleInfo {
    s.PresaleDeposit = &v
    return s
}
func (s *TaobaoTbkDgMaterialOptionalUpgradePresaleInfo) SetPresaleDiscountFeeText(v string) *TaobaoTbkDgMaterialOptionalUpgradePresaleInfo {
    s.PresaleDiscountFeeText = &v
    return s
}

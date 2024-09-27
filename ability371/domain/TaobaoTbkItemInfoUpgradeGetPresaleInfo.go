package domain


type TaobaoTbkItemInfoUpgradeGetPresaleInfo struct {
    /*
        预售商品-付尾款结束时间（毫秒）     */
    PresaleTailEndTime  *int64 `json:"presale_tail_end_time,omitempty" `

    /*
        预售商品-付尾款开始时间（毫秒）     */
    PresaleTailStartTime  *int64 `json:"presale_tail_start_time,omitempty" `

    /*
        预售商品-付定金结束时间（毫秒）     */
    PresaleEndTime  *int64 `json:"presale_end_time,omitempty" `

    /*
        预售商品-付定金开始时间（毫秒）     */
    PresaleStartTime  *int64 `json:"presale_start_time,omitempty" `

    /*
        预售商品-定金（元）     */
    PresaleDeposit  *string `json:"presale_deposit,omitempty" `

    /*
        预售商品-优惠信息     */
    PresaleDiscountFeeText  *string `json:"presale_discount_fee_text,omitempty" `

}

func (s *TaobaoTbkItemInfoUpgradeGetPresaleInfo) SetPresaleTailEndTime(v int64) *TaobaoTbkItemInfoUpgradeGetPresaleInfo {
    s.PresaleTailEndTime = &v
    return s
}
func (s *TaobaoTbkItemInfoUpgradeGetPresaleInfo) SetPresaleTailStartTime(v int64) *TaobaoTbkItemInfoUpgradeGetPresaleInfo {
    s.PresaleTailStartTime = &v
    return s
}
func (s *TaobaoTbkItemInfoUpgradeGetPresaleInfo) SetPresaleEndTime(v int64) *TaobaoTbkItemInfoUpgradeGetPresaleInfo {
    s.PresaleEndTime = &v
    return s
}
func (s *TaobaoTbkItemInfoUpgradeGetPresaleInfo) SetPresaleStartTime(v int64) *TaobaoTbkItemInfoUpgradeGetPresaleInfo {
    s.PresaleStartTime = &v
    return s
}
func (s *TaobaoTbkItemInfoUpgradeGetPresaleInfo) SetPresaleDeposit(v string) *TaobaoTbkItemInfoUpgradeGetPresaleInfo {
    s.PresaleDeposit = &v
    return s
}
func (s *TaobaoTbkItemInfoUpgradeGetPresaleInfo) SetPresaleDiscountFeeText(v string) *TaobaoTbkItemInfoUpgradeGetPresaleInfo {
    s.PresaleDiscountFeeText = &v
    return s
}

package domain


type TaobaoTbkOrderDetailsGetSubsidyDetailDTO struct {
    /*
        该笔订单包含的补贴类型     */
    SubsidyType  *string `json:"subsidy_type,omitempty" `

    /*
        补贴比率     */
    SubsidyRate  *string `json:"subsidy_rate,omitempty" `

    /*
        对应补贴类型的补贴金额     */
    SubsidyFee  *string `json:"subsidy_fee,omitempty" `

    /*
        单笔订单补贴上限。对应补贴类型的单笔订单补贴上限     */
    SubsidyUpperLimit  *string `json:"subsidy_upper_limit,omitempty" `

    /*
        补贴分成比率     */
    SubsidyShareRate  *string `json:"subsidy_share_rate,omitempty" `

}

func (s *TaobaoTbkOrderDetailsGetSubsidyDetailDTO) SetSubsidyType(v string) *TaobaoTbkOrderDetailsGetSubsidyDetailDTO {
    s.SubsidyType = &v
    return s
}
func (s *TaobaoTbkOrderDetailsGetSubsidyDetailDTO) SetSubsidyRate(v string) *TaobaoTbkOrderDetailsGetSubsidyDetailDTO {
    s.SubsidyRate = &v
    return s
}
func (s *TaobaoTbkOrderDetailsGetSubsidyDetailDTO) SetSubsidyFee(v string) *TaobaoTbkOrderDetailsGetSubsidyDetailDTO {
    s.SubsidyFee = &v
    return s
}
func (s *TaobaoTbkOrderDetailsGetSubsidyDetailDTO) SetSubsidyUpperLimit(v string) *TaobaoTbkOrderDetailsGetSubsidyDetailDTO {
    s.SubsidyUpperLimit = &v
    return s
}
func (s *TaobaoTbkOrderDetailsGetSubsidyDetailDTO) SetSubsidyShareRate(v string) *TaobaoTbkOrderDetailsGetSubsidyDetailDTO {
    s.SubsidyShareRate = &v
    return s
}

package domain


type TaobaoTbkOrderDetailsTemporaryGetServiceFeeDto struct {
    /*
        专项服务费来源，122-渠道     */
    TkShareRoleType  *int64 `json:"tk_share_role_type,omitempty" `

    /*
        专项服务费率     */
    ShareRelativeRate  *string `json:"share_relative_rate,omitempty" `

    /*
        结算专项服务费     */
    ShareFee  *string `json:"share_fee,omitempty" `

    /*
        预估专项服务费     */
    SharePreFee  *string `json:"share_pre_fee,omitempty" `

}

func (s *TaobaoTbkOrderDetailsTemporaryGetServiceFeeDto) SetTkShareRoleType(v int64) *TaobaoTbkOrderDetailsTemporaryGetServiceFeeDto {
    s.TkShareRoleType = &v
    return s
}
func (s *TaobaoTbkOrderDetailsTemporaryGetServiceFeeDto) SetShareRelativeRate(v string) *TaobaoTbkOrderDetailsTemporaryGetServiceFeeDto {
    s.ShareRelativeRate = &v
    return s
}
func (s *TaobaoTbkOrderDetailsTemporaryGetServiceFeeDto) SetShareFee(v string) *TaobaoTbkOrderDetailsTemporaryGetServiceFeeDto {
    s.ShareFee = &v
    return s
}
func (s *TaobaoTbkOrderDetailsTemporaryGetServiceFeeDto) SetSharePreFee(v string) *TaobaoTbkOrderDetailsTemporaryGetServiceFeeDto {
    s.SharePreFee = &v
    return s
}

package domain


type TaobaoTbkOrderDetailsGetServiceFeeDTO struct {
    /*
        专项服务费来源，122-渠道（字段已废弃）     */
    TkShareRoleType  *int64 `json:"tk_share_role_type,omitempty" `

    /*
        专项服务费率（字段已废弃）     */
    ShareRelativeRate  *string `json:"share_relative_rate,omitempty" `

    /*
        结算专项服务费（字段已废弃）     */
    ShareFee  *string `json:"share_fee,omitempty" `

    /*
        预估专项服务费（字段已废弃）     */
    SharePreFee  *string `json:"share_pre_fee,omitempty" `

}

func (s *TaobaoTbkOrderDetailsGetServiceFeeDTO) SetTkShareRoleType(v int64) *TaobaoTbkOrderDetailsGetServiceFeeDTO {
    s.TkShareRoleType = &v
    return s
}
func (s *TaobaoTbkOrderDetailsGetServiceFeeDTO) SetShareRelativeRate(v string) *TaobaoTbkOrderDetailsGetServiceFeeDTO {
    s.ShareRelativeRate = &v
    return s
}
func (s *TaobaoTbkOrderDetailsGetServiceFeeDTO) SetShareFee(v string) *TaobaoTbkOrderDetailsGetServiceFeeDTO {
    s.ShareFee = &v
    return s
}
func (s *TaobaoTbkOrderDetailsGetServiceFeeDTO) SetSharePreFee(v string) *TaobaoTbkOrderDetailsGetServiceFeeDTO {
    s.SharePreFee = &v
    return s
}

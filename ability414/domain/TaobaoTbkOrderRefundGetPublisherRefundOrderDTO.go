package domain


type TaobaoTbkOrderRefundGetPublisherRefundOrderDTO struct {
    /*
        维权编号，是当前订单发生维权退款的编号，非淘宝订单编号，如订单发生多次维权，则会产生多个维权编号     */
    RefundSuitId  *string `json:"refund_suit_id,omitempty" `

    /*
        淘宝父订单编号(买家在淘宝后台显示的订单编号)     */
    TbTradeParentId  *string `json:"tb_trade_parent_id,omitempty" `

    /*
        淘宝子订单编号     */
    TbTradeId  *string `json:"tb_trade_id,omitempty" `

    /*
        订单创建时间     */
    TbTradeCreateTime  *string `json:"tb_trade_create_time,omitempty" `

    /*
        订单结算时间     */
    EarningTime  *string `json:"earning_time,omitempty" `

    /*
        维权创建时间     */
    TkRefundTime  *string `json:"tk_refund_time,omitempty" `

    /*
        维权完成时间     */
    TkRefundSuitTime  *string `json:"tk_refund_suit_time,omitempty" `

    /*
        订单更新时间     */
    ModifiedTime  *string `json:"modified_time,omitempty" `

    /*
        商品标题     */
    ItemTitle  *string `json:"item_title,omitempty" `

    /*
        推广者角色(二方、三方)     */
    TkOrderRole  *string `json:"tk_order_role,omitempty" `

    /*
        维权创建(淘客结算回执) 4, 维权成功(淘客结算回执) 2, 维权失败(淘客结算回执) 3, 发生多次维权，待处理 11, 从淘客处补扣（钱已结给淘客） 等待扣款 12, 从淘客处补扣（钱已结给淘客） 扣款成功 13, 从卖家处补扣（钱已结给卖家） 等待扣款 14, 从卖家处补扣（钱已结给卖家） 扣款成功 15     */
    RefundStatus  *int64 `json:"refund_status,omitempty" `

    /*
        结算金额(订单确认收货后的成交金额）     */
    TbTradeFinishPrice  *string `json:"tb_trade_finish_price,omitempty" `

    /*
        维权金额(买家申请维权退款的金额)     */
    RefundFee  *string `json:"refund_fee,omitempty" `

    /*
        结算预估收入=结算金额*提成。以订单确认收货后的成交金额为基数，预估您可能获得的收入。     */
    PubShareFee  *string `json:"pub_share_fee,omitempty" `

    /*
        应退还佣金(不含技术服务费和渠道专项服务费)     */
    TkCommissionFeeRefund  *string `json:"tk_commission_fee_refund,omitempty" `

    /*
        应退还补贴(不含技术服务费和渠道专项服务费)     */
    TkSubsidyFeeRefund  *string `json:"tk_subsidy_fee_refund,omitempty" `

    /*
        应退还佣金对应的技术服务费     */
    TkCommissionAlimmRefundFee  *string `json:"tk_commission_alimm_refund_fee,omitempty" `

    /*
        应退还补贴对应的技术服务费     */
    TkSubsidyAlimmRefundFee  *string `json:"tk_subsidy_alimm_refund_fee,omitempty" `

    /*
        应退还佣金对应的渠道专项服务费     */
    TkCommissionAgentRefundFee  *string `json:"tk_commission_agent_refund_fee,omitempty" `

    /*
        应退还补贴对应的渠道专项服务费     */
    TkSubsidyAgentRefundFee  *string `json:"tk_subsidy_agent_refund_fee,omitempty" `

    /*
        应退还预估收入：订单发生维权退款应退还的预估收入（佣金+补贴，含技术服务费和渠道专项服务费）     */
    ShowReturnFee  *string `json:"show_return_fee,omitempty" `

    /*
        渠道关系id     */
    RelationId  *int64 `json:"relation_id,omitempty" `

    /*
        会员关系id     */
    SpecialId  *int64 `json:"special_id,omitempty" `

}

func (s *TaobaoTbkOrderRefundGetPublisherRefundOrderDTO) SetRefundSuitId(v string) *TaobaoTbkOrderRefundGetPublisherRefundOrderDTO {
    s.RefundSuitId = &v
    return s
}
func (s *TaobaoTbkOrderRefundGetPublisherRefundOrderDTO) SetTbTradeParentId(v string) *TaobaoTbkOrderRefundGetPublisherRefundOrderDTO {
    s.TbTradeParentId = &v
    return s
}
func (s *TaobaoTbkOrderRefundGetPublisherRefundOrderDTO) SetTbTradeId(v string) *TaobaoTbkOrderRefundGetPublisherRefundOrderDTO {
    s.TbTradeId = &v
    return s
}
func (s *TaobaoTbkOrderRefundGetPublisherRefundOrderDTO) SetTbTradeCreateTime(v string) *TaobaoTbkOrderRefundGetPublisherRefundOrderDTO {
    s.TbTradeCreateTime = &v
    return s
}
func (s *TaobaoTbkOrderRefundGetPublisherRefundOrderDTO) SetEarningTime(v string) *TaobaoTbkOrderRefundGetPublisherRefundOrderDTO {
    s.EarningTime = &v
    return s
}
func (s *TaobaoTbkOrderRefundGetPublisherRefundOrderDTO) SetTkRefundTime(v string) *TaobaoTbkOrderRefundGetPublisherRefundOrderDTO {
    s.TkRefundTime = &v
    return s
}
func (s *TaobaoTbkOrderRefundGetPublisherRefundOrderDTO) SetTkRefundSuitTime(v string) *TaobaoTbkOrderRefundGetPublisherRefundOrderDTO {
    s.TkRefundSuitTime = &v
    return s
}
func (s *TaobaoTbkOrderRefundGetPublisherRefundOrderDTO) SetModifiedTime(v string) *TaobaoTbkOrderRefundGetPublisherRefundOrderDTO {
    s.ModifiedTime = &v
    return s
}
func (s *TaobaoTbkOrderRefundGetPublisherRefundOrderDTO) SetItemTitle(v string) *TaobaoTbkOrderRefundGetPublisherRefundOrderDTO {
    s.ItemTitle = &v
    return s
}
func (s *TaobaoTbkOrderRefundGetPublisherRefundOrderDTO) SetTkOrderRole(v string) *TaobaoTbkOrderRefundGetPublisherRefundOrderDTO {
    s.TkOrderRole = &v
    return s
}
func (s *TaobaoTbkOrderRefundGetPublisherRefundOrderDTO) SetRefundStatus(v int64) *TaobaoTbkOrderRefundGetPublisherRefundOrderDTO {
    s.RefundStatus = &v
    return s
}
func (s *TaobaoTbkOrderRefundGetPublisherRefundOrderDTO) SetTbTradeFinishPrice(v string) *TaobaoTbkOrderRefundGetPublisherRefundOrderDTO {
    s.TbTradeFinishPrice = &v
    return s
}
func (s *TaobaoTbkOrderRefundGetPublisherRefundOrderDTO) SetRefundFee(v string) *TaobaoTbkOrderRefundGetPublisherRefundOrderDTO {
    s.RefundFee = &v
    return s
}
func (s *TaobaoTbkOrderRefundGetPublisherRefundOrderDTO) SetPubShareFee(v string) *TaobaoTbkOrderRefundGetPublisherRefundOrderDTO {
    s.PubShareFee = &v
    return s
}
func (s *TaobaoTbkOrderRefundGetPublisherRefundOrderDTO) SetTkCommissionFeeRefund(v string) *TaobaoTbkOrderRefundGetPublisherRefundOrderDTO {
    s.TkCommissionFeeRefund = &v
    return s
}
func (s *TaobaoTbkOrderRefundGetPublisherRefundOrderDTO) SetTkSubsidyFeeRefund(v string) *TaobaoTbkOrderRefundGetPublisherRefundOrderDTO {
    s.TkSubsidyFeeRefund = &v
    return s
}
func (s *TaobaoTbkOrderRefundGetPublisherRefundOrderDTO) SetTkCommissionAlimmRefundFee(v string) *TaobaoTbkOrderRefundGetPublisherRefundOrderDTO {
    s.TkCommissionAlimmRefundFee = &v
    return s
}
func (s *TaobaoTbkOrderRefundGetPublisherRefundOrderDTO) SetTkSubsidyAlimmRefundFee(v string) *TaobaoTbkOrderRefundGetPublisherRefundOrderDTO {
    s.TkSubsidyAlimmRefundFee = &v
    return s
}
func (s *TaobaoTbkOrderRefundGetPublisherRefundOrderDTO) SetTkCommissionAgentRefundFee(v string) *TaobaoTbkOrderRefundGetPublisherRefundOrderDTO {
    s.TkCommissionAgentRefundFee = &v
    return s
}
func (s *TaobaoTbkOrderRefundGetPublisherRefundOrderDTO) SetTkSubsidyAgentRefundFee(v string) *TaobaoTbkOrderRefundGetPublisherRefundOrderDTO {
    s.TkSubsidyAgentRefundFee = &v
    return s
}
func (s *TaobaoTbkOrderRefundGetPublisherRefundOrderDTO) SetShowReturnFee(v string) *TaobaoTbkOrderRefundGetPublisherRefundOrderDTO {
    s.ShowReturnFee = &v
    return s
}
func (s *TaobaoTbkOrderRefundGetPublisherRefundOrderDTO) SetRelationId(v int64) *TaobaoTbkOrderRefundGetPublisherRefundOrderDTO {
    s.RelationId = &v
    return s
}
func (s *TaobaoTbkOrderRefundGetPublisherRefundOrderDTO) SetSpecialId(v int64) *TaobaoTbkOrderRefundGetPublisherRefundOrderDTO {
    s.SpecialId = &v
    return s
}

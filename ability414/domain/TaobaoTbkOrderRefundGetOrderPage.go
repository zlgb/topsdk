package domain


type TaobaoTbkOrderRefundGetOrderPage struct {
    /*
        真正的业务数据结构     */
    Result  *[]TaobaoTbkOrderRefundGetPublisherRefundOrderDTO `json:"result,omitempty" `

    /*
        上一页     */
    PrePage  *int64 `json:"pre_page,omitempty" `

    /*
        下一页     */
    NextPage  *int64 `json:"next_page,omitempty" `

    /*
        页码     */
    PageNo  *int64 `json:"page_no,omitempty" `

    /*
        每页订单量     */
    PageSize  *int64 `json:"page_size,omitempty" `

    /*
        是否有下一页     */
    HasNext  *bool `json:"has_next,omitempty" `

    /*
        位点     */
    PositionIndex  *string `json:"position_index,omitempty" `

    /*
        是否有上一页     */
    HasPre  *bool `json:"has_pre,omitempty" `

}

func (s *TaobaoTbkOrderRefundGetOrderPage) SetResult(v []TaobaoTbkOrderRefundGetPublisherRefundOrderDTO) *TaobaoTbkOrderRefundGetOrderPage {
    s.Result = &v
    return s
}
func (s *TaobaoTbkOrderRefundGetOrderPage) SetPrePage(v int64) *TaobaoTbkOrderRefundGetOrderPage {
    s.PrePage = &v
    return s
}
func (s *TaobaoTbkOrderRefundGetOrderPage) SetNextPage(v int64) *TaobaoTbkOrderRefundGetOrderPage {
    s.NextPage = &v
    return s
}
func (s *TaobaoTbkOrderRefundGetOrderPage) SetPageNo(v int64) *TaobaoTbkOrderRefundGetOrderPage {
    s.PageNo = &v
    return s
}
func (s *TaobaoTbkOrderRefundGetOrderPage) SetPageSize(v int64) *TaobaoTbkOrderRefundGetOrderPage {
    s.PageSize = &v
    return s
}
func (s *TaobaoTbkOrderRefundGetOrderPage) SetHasNext(v bool) *TaobaoTbkOrderRefundGetOrderPage {
    s.HasNext = &v
    return s
}
func (s *TaobaoTbkOrderRefundGetOrderPage) SetPositionIndex(v string) *TaobaoTbkOrderRefundGetOrderPage {
    s.PositionIndex = &v
    return s
}
func (s *TaobaoTbkOrderRefundGetOrderPage) SetHasPre(v bool) *TaobaoTbkOrderRefundGetOrderPage {
    s.HasPre = &v
    return s
}

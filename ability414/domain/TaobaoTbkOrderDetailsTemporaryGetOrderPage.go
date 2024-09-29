package domain


type TaobaoTbkOrderDetailsTemporaryGetOrderPage struct {
    /*
        PublisherOrderDto     */
    Results  *[]TaobaoTbkOrderDetailsTemporaryGetPublisherOrderDto `json:"results,omitempty" `

    /*
        是否还有上一页     */
    HasPre  *bool `json:"has_pre,omitempty" `

    /*
        位点字段，由调用方原样传递     */
    PositionIndex  *string `json:"position_index,omitempty" `

    /*
        是否还有下一页     */
    HasNext  *bool `json:"has_next,omitempty" `

    /*
        页码     */
    PageNo  *int64 `json:"page_no,omitempty" `

    /*
        页大小     */
    PageSize  *int64 `json:"page_size,omitempty" `

}

func (s *TaobaoTbkOrderDetailsTemporaryGetOrderPage) SetResults(v []TaobaoTbkOrderDetailsTemporaryGetPublisherOrderDto) *TaobaoTbkOrderDetailsTemporaryGetOrderPage {
    s.Results = &v
    return s
}
func (s *TaobaoTbkOrderDetailsTemporaryGetOrderPage) SetHasPre(v bool) *TaobaoTbkOrderDetailsTemporaryGetOrderPage {
    s.HasPre = &v
    return s
}
func (s *TaobaoTbkOrderDetailsTemporaryGetOrderPage) SetPositionIndex(v string) *TaobaoTbkOrderDetailsTemporaryGetOrderPage {
    s.PositionIndex = &v
    return s
}
func (s *TaobaoTbkOrderDetailsTemporaryGetOrderPage) SetHasNext(v bool) *TaobaoTbkOrderDetailsTemporaryGetOrderPage {
    s.HasNext = &v
    return s
}
func (s *TaobaoTbkOrderDetailsTemporaryGetOrderPage) SetPageNo(v int64) *TaobaoTbkOrderDetailsTemporaryGetOrderPage {
    s.PageNo = &v
    return s
}
func (s *TaobaoTbkOrderDetailsTemporaryGetOrderPage) SetPageSize(v int64) *TaobaoTbkOrderDetailsTemporaryGetOrderPage {
    s.PageSize = &v
    return s
}

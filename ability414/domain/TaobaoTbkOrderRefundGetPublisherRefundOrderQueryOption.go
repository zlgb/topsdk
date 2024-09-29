package domain


type TaobaoTbkOrderRefundGetPublisherRefundOrderQueryOption struct {
    /*
        跳转类型，当向前或者向后翻页必须提供,-1: 向前翻页,1：向后翻页     */
    JumpType  *int64 `json:"jump_type,omitempty" `

    /*
        1全部订单     */
    OrderScene  *int64 `json:"order_scene,omitempty" `

    /*
        单页订单量，1-100，上限100条     */
    PageSize  *int64 `json:"page_size,omitempty" `

    /*
        1-维权发起时间，2-订单结算时间（正向订单），3-维权完成时间，4-订单创建时间，5-订单更新时间 （注意：入参startTime小于2022年12月20日0点时，query_type只能为1、2、3；入参startTime大于等于2022年12月20日0点时，query_type可以为1、2、3、4、5)     */
    QueryType  *int64 `json:"query_type,omitempty" `

    /*
        查询页数     */
    PageNo  *int64 `json:"page_no,omitempty" `

    /*
        位点，除第一页之外，都需要传递；前端原样返回。     */
    PositionIndex  *string `json:"position_index,omitempty" `

    /*
        开始时间     */
    StartTime  *string `json:"start_time,omitempty" `

    /*
        1 表示2方，2表示3方，4表示不限     */
    MemberType  *int64 `json:"member_type,omitempty" `

    /*
        结束时间     */
    EndTime  *string `json:"end_time,omitempty" `

}

func (s *TaobaoTbkOrderRefundGetPublisherRefundOrderQueryOption) SetJumpType(v int64) *TaobaoTbkOrderRefundGetPublisherRefundOrderQueryOption {
    s.JumpType = &v
    return s
}
func (s *TaobaoTbkOrderRefundGetPublisherRefundOrderQueryOption) SetOrderScene(v int64) *TaobaoTbkOrderRefundGetPublisherRefundOrderQueryOption {
    s.OrderScene = &v
    return s
}
func (s *TaobaoTbkOrderRefundGetPublisherRefundOrderQueryOption) SetPageSize(v int64) *TaobaoTbkOrderRefundGetPublisherRefundOrderQueryOption {
    s.PageSize = &v
    return s
}
func (s *TaobaoTbkOrderRefundGetPublisherRefundOrderQueryOption) SetQueryType(v int64) *TaobaoTbkOrderRefundGetPublisherRefundOrderQueryOption {
    s.QueryType = &v
    return s
}
func (s *TaobaoTbkOrderRefundGetPublisherRefundOrderQueryOption) SetPageNo(v int64) *TaobaoTbkOrderRefundGetPublisherRefundOrderQueryOption {
    s.PageNo = &v
    return s
}
func (s *TaobaoTbkOrderRefundGetPublisherRefundOrderQueryOption) SetPositionIndex(v string) *TaobaoTbkOrderRefundGetPublisherRefundOrderQueryOption {
    s.PositionIndex = &v
    return s
}
func (s *TaobaoTbkOrderRefundGetPublisherRefundOrderQueryOption) SetStartTime(v string) *TaobaoTbkOrderRefundGetPublisherRefundOrderQueryOption {
    s.StartTime = &v
    return s
}
func (s *TaobaoTbkOrderRefundGetPublisherRefundOrderQueryOption) SetMemberType(v int64) *TaobaoTbkOrderRefundGetPublisherRefundOrderQueryOption {
    s.MemberType = &v
    return s
}
func (s *TaobaoTbkOrderRefundGetPublisherRefundOrderQueryOption) SetEndTime(v string) *TaobaoTbkOrderRefundGetPublisherRefundOrderQueryOption {
    s.EndTime = &v
    return s
}

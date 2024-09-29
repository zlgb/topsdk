package domain


type TaobaoTbkDgGeneralLinkConvertItemUrlList struct {
    /*
        物料对应错误描述     */
    Msg  *string `json:"msg,omitempty" `

    /*
        物料对应错误码     */
    Code  *int64 `json:"code,omitempty" `

    /*
        营销信息     */
    PromotionInfoDto  *TaobaoTbkDgGeneralLinkConvertPromotionInfoDTO `json:"promotion_info_dto,omitempty" `

    /*
        券信息     */
    CouponInfoDto  *TaobaoTbkDgGeneralLinkConvertCouponInfoDTO `json:"coupon_info_dto,omitempty" `

    /*
        链接信息     */
    LinkInfoDto  *TaobaoTbkDgGeneralLinkConvertItemLinkInfoDTO `json:"link_info_dto,omitempty" `

    /*
        入参的商品ID     */
    InputItemId  *string `json:"input_item_id,omitempty" `

    /*
        转链成功的场景下，需要补充说明的信息     */
    ExtraInfo  *string `json:"extra_info,omitempty" `

}

func (s *TaobaoTbkDgGeneralLinkConvertItemUrlList) SetMsg(v string) *TaobaoTbkDgGeneralLinkConvertItemUrlList {
    s.Msg = &v
    return s
}
func (s *TaobaoTbkDgGeneralLinkConvertItemUrlList) SetCode(v int64) *TaobaoTbkDgGeneralLinkConvertItemUrlList {
    s.Code = &v
    return s
}
func (s *TaobaoTbkDgGeneralLinkConvertItemUrlList) SetPromotionInfoDto(v TaobaoTbkDgGeneralLinkConvertPromotionInfoDTO) *TaobaoTbkDgGeneralLinkConvertItemUrlList {
    s.PromotionInfoDto = &v
    return s
}
func (s *TaobaoTbkDgGeneralLinkConvertItemUrlList) SetCouponInfoDto(v TaobaoTbkDgGeneralLinkConvertCouponInfoDTO) *TaobaoTbkDgGeneralLinkConvertItemUrlList {
    s.CouponInfoDto = &v
    return s
}
func (s *TaobaoTbkDgGeneralLinkConvertItemUrlList) SetLinkInfoDto(v TaobaoTbkDgGeneralLinkConvertItemLinkInfoDTO) *TaobaoTbkDgGeneralLinkConvertItemUrlList {
    s.LinkInfoDto = &v
    return s
}
func (s *TaobaoTbkDgGeneralLinkConvertItemUrlList) SetInputItemId(v string) *TaobaoTbkDgGeneralLinkConvertItemUrlList {
    s.InputItemId = &v
    return s
}
func (s *TaobaoTbkDgGeneralLinkConvertItemUrlList) SetExtraInfo(v string) *TaobaoTbkDgGeneralLinkConvertItemUrlList {
    s.ExtraInfo = &v
    return s
}

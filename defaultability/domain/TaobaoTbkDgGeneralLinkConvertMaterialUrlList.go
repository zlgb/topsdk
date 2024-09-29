package domain


type TaobaoTbkDgGeneralLinkConvertMaterialUrlList struct {
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
    LinkInfoDto  *TaobaoTbkDgGeneralLinkConvertLinkInfoDTO `json:"link_info_dto,omitempty" `

    /*
        入参物料链接     */
    InputMaterialUrl  *string `json:"input_material_url,omitempty" `

    /*
        转链成功的场景下，需要补充说明的信息     */
    ExtraInfo  *string `json:"extra_info,omitempty" `

}

func (s *TaobaoTbkDgGeneralLinkConvertMaterialUrlList) SetMsg(v string) *TaobaoTbkDgGeneralLinkConvertMaterialUrlList {
    s.Msg = &v
    return s
}
func (s *TaobaoTbkDgGeneralLinkConvertMaterialUrlList) SetCode(v int64) *TaobaoTbkDgGeneralLinkConvertMaterialUrlList {
    s.Code = &v
    return s
}
func (s *TaobaoTbkDgGeneralLinkConvertMaterialUrlList) SetPromotionInfoDto(v TaobaoTbkDgGeneralLinkConvertPromotionInfoDTO) *TaobaoTbkDgGeneralLinkConvertMaterialUrlList {
    s.PromotionInfoDto = &v
    return s
}
func (s *TaobaoTbkDgGeneralLinkConvertMaterialUrlList) SetCouponInfoDto(v TaobaoTbkDgGeneralLinkConvertCouponInfoDTO) *TaobaoTbkDgGeneralLinkConvertMaterialUrlList {
    s.CouponInfoDto = &v
    return s
}
func (s *TaobaoTbkDgGeneralLinkConvertMaterialUrlList) SetLinkInfoDto(v TaobaoTbkDgGeneralLinkConvertLinkInfoDTO) *TaobaoTbkDgGeneralLinkConvertMaterialUrlList {
    s.LinkInfoDto = &v
    return s
}
func (s *TaobaoTbkDgGeneralLinkConvertMaterialUrlList) SetInputMaterialUrl(v string) *TaobaoTbkDgGeneralLinkConvertMaterialUrlList {
    s.InputMaterialUrl = &v
    return s
}
func (s *TaobaoTbkDgGeneralLinkConvertMaterialUrlList) SetExtraInfo(v string) *TaobaoTbkDgGeneralLinkConvertMaterialUrlList {
    s.ExtraInfo = &v
    return s
}

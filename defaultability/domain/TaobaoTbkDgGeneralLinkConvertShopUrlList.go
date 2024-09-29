package domain


type TaobaoTbkDgGeneralLinkConvertShopUrlList struct {
    /*
        物料对应错误描述     */
    Msg  *string `json:"msg,omitempty" `

    /*
        物料对应错误码     */
    Code  *int64 `json:"code,omitempty" `

    /*
        链接信息     */
    LinkInfoDto  *TaobaoTbkDgGeneralLinkConvertLinkInfoDTO `json:"link_info_dto,omitempty" `

    /*
        入参的店铺ID     */
    InputShopId  *string `json:"input_shop_id,omitempty" `

}

func (s *TaobaoTbkDgGeneralLinkConvertShopUrlList) SetMsg(v string) *TaobaoTbkDgGeneralLinkConvertShopUrlList {
    s.Msg = &v
    return s
}
func (s *TaobaoTbkDgGeneralLinkConvertShopUrlList) SetCode(v int64) *TaobaoTbkDgGeneralLinkConvertShopUrlList {
    s.Code = &v
    return s
}
func (s *TaobaoTbkDgGeneralLinkConvertShopUrlList) SetLinkInfoDto(v TaobaoTbkDgGeneralLinkConvertLinkInfoDTO) *TaobaoTbkDgGeneralLinkConvertShopUrlList {
    s.LinkInfoDto = &v
    return s
}
func (s *TaobaoTbkDgGeneralLinkConvertShopUrlList) SetInputShopId(v string) *TaobaoTbkDgGeneralLinkConvertShopUrlList {
    s.InputShopId = &v
    return s
}

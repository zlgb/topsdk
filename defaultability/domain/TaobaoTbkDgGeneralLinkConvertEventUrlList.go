package domain


type TaobaoTbkDgGeneralLinkConvertEventUrlList struct {
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
        入参的会场ID     */
    InputPageId  *string `json:"input_page_id,omitempty" `

}

func (s *TaobaoTbkDgGeneralLinkConvertEventUrlList) SetMsg(v string) *TaobaoTbkDgGeneralLinkConvertEventUrlList {
    s.Msg = &v
    return s
}
func (s *TaobaoTbkDgGeneralLinkConvertEventUrlList) SetCode(v int64) *TaobaoTbkDgGeneralLinkConvertEventUrlList {
    s.Code = &v
    return s
}
func (s *TaobaoTbkDgGeneralLinkConvertEventUrlList) SetLinkInfoDto(v TaobaoTbkDgGeneralLinkConvertLinkInfoDTO) *TaobaoTbkDgGeneralLinkConvertEventUrlList {
    s.LinkInfoDto = &v
    return s
}
func (s *TaobaoTbkDgGeneralLinkConvertEventUrlList) SetInputPageId(v string) *TaobaoTbkDgGeneralLinkConvertEventUrlList {
    s.InputPageId = &v
    return s
}

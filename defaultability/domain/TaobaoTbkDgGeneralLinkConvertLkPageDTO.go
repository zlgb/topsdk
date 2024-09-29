package domain


type TaobaoTbkDgGeneralLinkConvertLkPageDTO struct {
    /*
        会场ID     */
    PageId  *string `json:"page_id,omitempty" `

    /*
        页面内定坑商品ID，用于素材-坑位还原     */
    TargetItemList  *[]string `json:"target_item_list,omitempty" `

}

func (s *TaobaoTbkDgGeneralLinkConvertLkPageDTO) SetPageId(v string) *TaobaoTbkDgGeneralLinkConvertLkPageDTO {
    s.PageId = &v
    return s
}
func (s *TaobaoTbkDgGeneralLinkConvertLkPageDTO) SetTargetItemList(v []string) *TaobaoTbkDgGeneralLinkConvertLkPageDTO {
    s.TargetItemList = &v
    return s
}

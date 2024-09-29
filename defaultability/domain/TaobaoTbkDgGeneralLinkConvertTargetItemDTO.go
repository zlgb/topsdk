package domain


type TaobaoTbkDgGeneralLinkConvertTargetItemDTO struct {
    /*
        页面内定坑商品ID，用于素材-坑位还原     */
    ItemIdList  *[]string `json:"item_id_list,omitempty" `

}

func (s *TaobaoTbkDgGeneralLinkConvertTargetItemDTO) SetItemIdList(v []string) *TaobaoTbkDgGeneralLinkConvertTargetItemDTO {
    s.ItemIdList = &v
    return s
}

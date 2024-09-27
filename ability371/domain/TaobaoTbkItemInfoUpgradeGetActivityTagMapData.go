package domain


type TaobaoTbkItemInfoUpgradeGetActivityTagMapData struct {
    /*
        货品展示标识，展示在商品标题前的商品活动标     */
    TagName  *string `json:"tag_name,omitempty" `

}

func (s *TaobaoTbkItemInfoUpgradeGetActivityTagMapData) SetTagName(v string) *TaobaoTbkItemInfoUpgradeGetActivityTagMapData {
    s.TagName = &v
    return s
}

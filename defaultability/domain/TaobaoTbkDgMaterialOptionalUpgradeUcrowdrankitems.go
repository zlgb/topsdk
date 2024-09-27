package domain


type TaobaoTbkDgMaterialOptionalUpgradeUcrowdrankitems struct {
    /*
        物料评估-商品佣金率，如：1234表示12.34%，material_id=41377时选填     */
    Commirate  *int64 `json:"commirate,omitempty" `

    /*
        物料评估-商品价格，单位：元，material_id=41377时选填     */
    Price  *string `json:"price,omitempty" `

    /*
        物料评估-商品ID，material_id=41377时必填     */
    ItemId  *string `json:"item_id,omitempty" `

}

func (s *TaobaoTbkDgMaterialOptionalUpgradeUcrowdrankitems) SetCommirate(v int64) *TaobaoTbkDgMaterialOptionalUpgradeUcrowdrankitems {
    s.Commirate = &v
    return s
}
func (s *TaobaoTbkDgMaterialOptionalUpgradeUcrowdrankitems) SetPrice(v string) *TaobaoTbkDgMaterialOptionalUpgradeUcrowdrankitems {
    s.Price = &v
    return s
}
func (s *TaobaoTbkDgMaterialOptionalUpgradeUcrowdrankitems) SetItemId(v string) *TaobaoTbkDgMaterialOptionalUpgradeUcrowdrankitems {
    s.ItemId = &v
    return s
}

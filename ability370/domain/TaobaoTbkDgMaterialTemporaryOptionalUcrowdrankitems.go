package domain


type TaobaoTbkDgMaterialTemporaryOptionalUcrowdrankitems struct {
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

func (s *TaobaoTbkDgMaterialTemporaryOptionalUcrowdrankitems) SetCommirate(v int64) *TaobaoTbkDgMaterialTemporaryOptionalUcrowdrankitems {
    s.Commirate = &v
    return s
}
func (s *TaobaoTbkDgMaterialTemporaryOptionalUcrowdrankitems) SetPrice(v string) *TaobaoTbkDgMaterialTemporaryOptionalUcrowdrankitems {
    s.Price = &v
    return s
}
func (s *TaobaoTbkDgMaterialTemporaryOptionalUcrowdrankitems) SetItemId(v string) *TaobaoTbkDgMaterialTemporaryOptionalUcrowdrankitems {
    s.ItemId = &v
    return s
}

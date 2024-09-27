package domain


type TaobaoTbkOptimusTouMaterialIdsGetMaterialQuery struct {
    /*
        页码，默认1，取值范围1~100 defalutValue:1    */
    PageNo  *int64 `json:"page_no,omitempty" `

    /*
        物料主题类型, 1促销活动;2热门主题;3精选榜单;4行业频道等;5其他 defalutValue:1    */
    Subject  *int64 `json:"subject,omitempty" `

    /*
        物料类型，1: 商品；2:权益 defalutValue:1    */
    MaterialType  *int64 `json:"material_type,omitempty" `

    /*
        每页物料id数量，默认20，取值范围1~100 defalutValue:20    */
    PageSize  *int64 `json:"page_size,omitempty" `

}

func (s *TaobaoTbkOptimusTouMaterialIdsGetMaterialQuery) SetPageNo(v int64) *TaobaoTbkOptimusTouMaterialIdsGetMaterialQuery {
    s.PageNo = &v
    return s
}
func (s *TaobaoTbkOptimusTouMaterialIdsGetMaterialQuery) SetSubject(v int64) *TaobaoTbkOptimusTouMaterialIdsGetMaterialQuery {
    s.Subject = &v
    return s
}
func (s *TaobaoTbkOptimusTouMaterialIdsGetMaterialQuery) SetMaterialType(v int64) *TaobaoTbkOptimusTouMaterialIdsGetMaterialQuery {
    s.MaterialType = &v
    return s
}
func (s *TaobaoTbkOptimusTouMaterialIdsGetMaterialQuery) SetPageSize(v int64) *TaobaoTbkOptimusTouMaterialIdsGetMaterialQuery {
    s.PageSize = &v
    return s
}

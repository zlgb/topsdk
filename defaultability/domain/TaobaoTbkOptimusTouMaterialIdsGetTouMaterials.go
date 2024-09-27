package domain


type TaobaoTbkOptimusTouMaterialIdsGetTouMaterials struct {
    /*
        物料集合名称     */
    MaterialName  *string `json:"material_name,omitempty" `

    /*
        物料id     */
    MaterialId  *int64 `json:"material_id,omitempty" `

    /*
        物料类型，1: 商品；2:权益     */
    MaterialType  *int64 `json:"material_type,omitempty" `

    /*
        物料主题类型, 1促销活动;2热门主题;3精选榜单;4行业频道等;5其他     */
    Subject  *int64 `json:"subject,omitempty" `

    /*
        开始时间，Unix时间戳     */
    StartTime  *int64 `json:"start_time,omitempty" `

    /*
        结束时间，Unix时间戳     */
    EndTime  *int64 `json:"end_time,omitempty" `

}

func (s *TaobaoTbkOptimusTouMaterialIdsGetTouMaterials) SetMaterialName(v string) *TaobaoTbkOptimusTouMaterialIdsGetTouMaterials {
    s.MaterialName = &v
    return s
}
func (s *TaobaoTbkOptimusTouMaterialIdsGetTouMaterials) SetMaterialId(v int64) *TaobaoTbkOptimusTouMaterialIdsGetTouMaterials {
    s.MaterialId = &v
    return s
}
func (s *TaobaoTbkOptimusTouMaterialIdsGetTouMaterials) SetMaterialType(v int64) *TaobaoTbkOptimusTouMaterialIdsGetTouMaterials {
    s.MaterialType = &v
    return s
}
func (s *TaobaoTbkOptimusTouMaterialIdsGetTouMaterials) SetSubject(v int64) *TaobaoTbkOptimusTouMaterialIdsGetTouMaterials {
    s.Subject = &v
    return s
}
func (s *TaobaoTbkOptimusTouMaterialIdsGetTouMaterials) SetStartTime(v int64) *TaobaoTbkOptimusTouMaterialIdsGetTouMaterials {
    s.StartTime = &v
    return s
}
func (s *TaobaoTbkOptimusTouMaterialIdsGetTouMaterials) SetEndTime(v int64) *TaobaoTbkOptimusTouMaterialIdsGetTouMaterials {
    s.EndTime = &v
    return s
}

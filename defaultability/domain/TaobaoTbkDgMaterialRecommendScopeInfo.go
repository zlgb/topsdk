package domain


type TaobaoTbkDgMaterialRecommendScopeInfo struct {
    /*
        是否品牌精选，0不是，1是     */
    SuperiorBrand  *string `json:"superior_brand,omitempty" `

}

func (s *TaobaoTbkDgMaterialRecommendScopeInfo) SetSuperiorBrand(v string) *TaobaoTbkDgMaterialRecommendScopeInfo {
    s.SuperiorBrand = &v
    return s
}

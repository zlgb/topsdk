package domain


type TaobaoTbkDgMaterialOptionalUpgradeScopeInfo struct {
    /*
        是否品牌精选，0不是，1是     */
    SuperiorBrand  *string `json:"superior_brand,omitempty" `

}

func (s *TaobaoTbkDgMaterialOptionalUpgradeScopeInfo) SetSuperiorBrand(v string) *TaobaoTbkDgMaterialOptionalUpgradeScopeInfo {
    s.SuperiorBrand = &v
    return s
}

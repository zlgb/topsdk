package domain


type TaobaoTbkDgMaterialRecommendIncomeInfo struct {
    /*
        商品佣金比率     */
    CommissionRate  *string `json:"commission_rate,omitempty" `

    /*
        商品佣金金额     */
    CommissionAmount  *string `json:"commission_amount,omitempty" `

    /*
        补贴比率     */
    SubsidyRate  *string `json:"subsidy_rate,omitempty" `

    /*
        补贴金额     */
    SubsidyAmount  *string `json:"subsidy_amount,omitempty" `

    /*
        补贴上限；仅在单笔订单命中补贴上限时返回结果否则出参为空     */
    SubsidyUpperLimit  *string `json:"subsidy_upper_limit,omitempty" `

    /*
        补贴类型     */
    SubsidyType  *string `json:"subsidy_type,omitempty" `

}

func (s *TaobaoTbkDgMaterialRecommendIncomeInfo) SetCommissionRate(v string) *TaobaoTbkDgMaterialRecommendIncomeInfo {
    s.CommissionRate = &v
    return s
}
func (s *TaobaoTbkDgMaterialRecommendIncomeInfo) SetCommissionAmount(v string) *TaobaoTbkDgMaterialRecommendIncomeInfo {
    s.CommissionAmount = &v
    return s
}
func (s *TaobaoTbkDgMaterialRecommendIncomeInfo) SetSubsidyRate(v string) *TaobaoTbkDgMaterialRecommendIncomeInfo {
    s.SubsidyRate = &v
    return s
}
func (s *TaobaoTbkDgMaterialRecommendIncomeInfo) SetSubsidyAmount(v string) *TaobaoTbkDgMaterialRecommendIncomeInfo {
    s.SubsidyAmount = &v
    return s
}
func (s *TaobaoTbkDgMaterialRecommendIncomeInfo) SetSubsidyUpperLimit(v string) *TaobaoTbkDgMaterialRecommendIncomeInfo {
    s.SubsidyUpperLimit = &v
    return s
}
func (s *TaobaoTbkDgMaterialRecommendIncomeInfo) SetSubsidyType(v string) *TaobaoTbkDgMaterialRecommendIncomeInfo {
    s.SubsidyType = &v
    return s
}

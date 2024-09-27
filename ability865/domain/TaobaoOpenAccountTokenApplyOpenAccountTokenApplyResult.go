package domain


type TaobaoOpenAccountTokenApplyOpenAccountTokenApplyResult struct {
    /*
        错误码     */
    Code  *int64 `json:"code,omitempty" `

    /*
        token     */
    Data  *string `json:"data,omitempty" `

    /*
        错误信息     */
    Message  *string `json:"message,omitempty" `

    /*
        是否成功     */
    Successful  *bool `json:"successful,omitempty" `

}

func (s *TaobaoOpenAccountTokenApplyOpenAccountTokenApplyResult) SetCode(v int64) *TaobaoOpenAccountTokenApplyOpenAccountTokenApplyResult {
    s.Code = &v
    return s
}
func (s *TaobaoOpenAccountTokenApplyOpenAccountTokenApplyResult) SetData(v string) *TaobaoOpenAccountTokenApplyOpenAccountTokenApplyResult {
    s.Data = &v
    return s
}
func (s *TaobaoOpenAccountTokenApplyOpenAccountTokenApplyResult) SetMessage(v string) *TaobaoOpenAccountTokenApplyOpenAccountTokenApplyResult {
    s.Message = &v
    return s
}
func (s *TaobaoOpenAccountTokenApplyOpenAccountTokenApplyResult) SetSuccessful(v bool) *TaobaoOpenAccountTokenApplyOpenAccountTokenApplyResult {
    s.Successful = &v
    return s
}

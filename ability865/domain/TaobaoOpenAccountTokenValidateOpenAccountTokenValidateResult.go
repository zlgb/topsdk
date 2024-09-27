package domain


type TaobaoOpenAccountTokenValidateOpenAccountTokenValidateResult struct {
    /*
        错误码     */
    Code  *int64 `json:"code,omitempty" `

    /*
        token中的数据     */
    Data  *TaobaoOpenAccountTokenValidateTokenInfo `json:"data,omitempty" `

    /*
        错误信息     */
    Message  *string `json:"message,omitempty" `

    /*
        是否成功     */
    Successful  *bool `json:"successful,omitempty" `

}

func (s *TaobaoOpenAccountTokenValidateOpenAccountTokenValidateResult) SetCode(v int64) *TaobaoOpenAccountTokenValidateOpenAccountTokenValidateResult {
    s.Code = &v
    return s
}
func (s *TaobaoOpenAccountTokenValidateOpenAccountTokenValidateResult) SetData(v TaobaoOpenAccountTokenValidateTokenInfo) *TaobaoOpenAccountTokenValidateOpenAccountTokenValidateResult {
    s.Data = &v
    return s
}
func (s *TaobaoOpenAccountTokenValidateOpenAccountTokenValidateResult) SetMessage(v string) *TaobaoOpenAccountTokenValidateOpenAccountTokenValidateResult {
    s.Message = &v
    return s
}
func (s *TaobaoOpenAccountTokenValidateOpenAccountTokenValidateResult) SetSuccessful(v bool) *TaobaoOpenAccountTokenValidateOpenAccountTokenValidateResult {
    s.Successful = &v
    return s
}

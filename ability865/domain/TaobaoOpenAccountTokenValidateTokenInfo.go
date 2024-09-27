package domain


type TaobaoOpenAccountTokenValidateTokenInfo struct {
    /*
        token info扩展信息     */
    Ext  *TaobaoOpenAccountTokenValidateTokenInfoExt `json:"ext,omitempty" `

    /*
        isv自己账号的唯一id     */
    IsvAccountId  *string `json:"isv_account_id,omitempty" `

    /*
        ISV APP的登录态时长     */
    LoginStateExpireIn  *int64 `json:"login_state_expire_in,omitempty" `

    /*
        open account id     */
    OpenAccountId  *int64 `json:"open_account_id,omitempty" `

    /*
        时间戳     */
    Timestamp  *int64 `json:"timestamp,omitempty" `

    /*
        用于防重放的唯一id     */
    Uuid  *string `json:"uuid,omitempty" `

}

func (s *TaobaoOpenAccountTokenValidateTokenInfo) SetExt(v TaobaoOpenAccountTokenValidateTokenInfoExt) *TaobaoOpenAccountTokenValidateTokenInfo {
    s.Ext = &v
    return s
}
func (s *TaobaoOpenAccountTokenValidateTokenInfo) SetIsvAccountId(v string) *TaobaoOpenAccountTokenValidateTokenInfo {
    s.IsvAccountId = &v
    return s
}
func (s *TaobaoOpenAccountTokenValidateTokenInfo) SetLoginStateExpireIn(v int64) *TaobaoOpenAccountTokenValidateTokenInfo {
    s.LoginStateExpireIn = &v
    return s
}
func (s *TaobaoOpenAccountTokenValidateTokenInfo) SetOpenAccountId(v int64) *TaobaoOpenAccountTokenValidateTokenInfo {
    s.OpenAccountId = &v
    return s
}
func (s *TaobaoOpenAccountTokenValidateTokenInfo) SetTimestamp(v int64) *TaobaoOpenAccountTokenValidateTokenInfo {
    s.Timestamp = &v
    return s
}
func (s *TaobaoOpenAccountTokenValidateTokenInfo) SetUuid(v string) *TaobaoOpenAccountTokenValidateTokenInfo {
    s.Uuid = &v
    return s
}

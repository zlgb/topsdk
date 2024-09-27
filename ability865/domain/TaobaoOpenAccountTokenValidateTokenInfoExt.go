package domain


type TaobaoOpenAccountTokenValidateTokenInfoExt struct {
    /*
        open account当前token info中open account id对应的open account信息     */
    OpenAccount  *TaobaoOpenAccountTokenValidateOpenAccount `json:"open_account,omitempty" `

    /*
        oauthOtherInfo     */
    OauthOtherInfo  *TaobaoOpenAccountTokenValidateOAuthOtherInfo `json:"oauth_other_info,omitempty" `

}

func (s *TaobaoOpenAccountTokenValidateTokenInfoExt) SetOpenAccount(v TaobaoOpenAccountTokenValidateOpenAccount) *TaobaoOpenAccountTokenValidateTokenInfoExt {
    s.OpenAccount = &v
    return s
}
func (s *TaobaoOpenAccountTokenValidateTokenInfoExt) SetOauthOtherInfo(v TaobaoOpenAccountTokenValidateOAuthOtherInfo) *TaobaoOpenAccountTokenValidateTokenInfoExt {
    s.OauthOtherInfo = &v
    return s
}

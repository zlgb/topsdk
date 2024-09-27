package domain


type TaobaoOpenAccountTokenValidateOAuthOtherInfo struct {
    /*
        access_token     */
    AccessToken  *string `json:"access_token,omitempty" `

    /*
        nick     */
    Nick  *string `json:"nick,omitempty" `

    /*
        avatarUrl     */
    AvatarUrl  *string `json:"avatar_url,omitempty" `

    /*
        id     */
    Id  *string `json:"id,omitempty" `

    /*
        三方平台类型     */
    PlatformType  *int64 `json:"platform_type,omitempty" `

    /*
        三方平台的openId     */
    OpenId  *string `json:"open_id,omitempty" `

    /*
        三方平台的unionId     */
    UnionId  *string `json:"union_id,omitempty" `

}

func (s *TaobaoOpenAccountTokenValidateOAuthOtherInfo) SetAccessToken(v string) *TaobaoOpenAccountTokenValidateOAuthOtherInfo {
    s.AccessToken = &v
    return s
}
func (s *TaobaoOpenAccountTokenValidateOAuthOtherInfo) SetNick(v string) *TaobaoOpenAccountTokenValidateOAuthOtherInfo {
    s.Nick = &v
    return s
}
func (s *TaobaoOpenAccountTokenValidateOAuthOtherInfo) SetAvatarUrl(v string) *TaobaoOpenAccountTokenValidateOAuthOtherInfo {
    s.AvatarUrl = &v
    return s
}
func (s *TaobaoOpenAccountTokenValidateOAuthOtherInfo) SetId(v string) *TaobaoOpenAccountTokenValidateOAuthOtherInfo {
    s.Id = &v
    return s
}
func (s *TaobaoOpenAccountTokenValidateOAuthOtherInfo) SetPlatformType(v int64) *TaobaoOpenAccountTokenValidateOAuthOtherInfo {
    s.PlatformType = &v
    return s
}
func (s *TaobaoOpenAccountTokenValidateOAuthOtherInfo) SetOpenId(v string) *TaobaoOpenAccountTokenValidateOAuthOtherInfo {
    s.OpenId = &v
    return s
}
func (s *TaobaoOpenAccountTokenValidateOAuthOtherInfo) SetUnionId(v string) *TaobaoOpenAccountTokenValidateOAuthOtherInfo {
    s.UnionId = &v
    return s
}

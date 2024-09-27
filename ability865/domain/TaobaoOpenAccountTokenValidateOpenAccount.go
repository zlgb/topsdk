package domain


type TaobaoOpenAccountTokenValidateOpenAccount struct {
    /*
        登录名     */
    LoginId  *string `json:"login_id,omitempty" `

    /*
        帐号创建的设备的ID     */
    CreateDeviceId  *string `json:"create_device_id,omitempty" `

    /*
        支付宝的帐号标识     */
    AlipayId  *string `json:"alipay_id,omitempty" `

    /*
        创建帐号的App Key     */
    CreateAppKey  *string `json:"create_app_key,omitempty" `

    /*
        地区     */
    Locale  *string `json:"locale,omitempty" `

    /*
        银行卡号     */
    BankCardNo  *string `json:"bank_card_no,omitempty" `

    /*
        开发者自定义账号id     */
    IsvAccountId  *string `json:"isv_account_id,omitempty" `

    /*
        邮箱     */
    Email  *string `json:"email,omitempty" `

    /*
        头像url     */
    AvatarUrl  *string `json:"avatar_url,omitempty" `

    /*
        数据域     */
    DomainId  *int64 `json:"domain_id,omitempty" `

    /*
        银行卡的拥有者姓名     */
    BankCardOwnerName  *string `json:"bank_card_owner_name,omitempty" `

    /*
        展示名     */
    DisplayName  *string `json:"display_name,omitempty" `

    /*
        密码salt     */
    LoginPwdSalt  *string `json:"login_pwd_salt,omitempty" `

    /*
        密码     */
    LoginPwd  *string `json:"login_pwd,omitempty" `

    /*
        第三方oauth openid     */
    OpenId  *string `json:"open_id,omitempty" `

    /*
        手机     */
    Mobile  *string `json:"mobile,omitempty" `

    /*
        账号创建时的位置     */
    CreateLocation  *string `json:"create_location,omitempty" `

    /*
        自定义扩展信息Map的Json格式     */
    ExtInfos  *string `json:"ext_infos,omitempty" `

    /*
        密码加密强度     */
    LoginPwdIntensity  *int64 `json:"login_pwd_intensity,omitempty" `

    /*
        Open Account Id     */
    Id  *int64 `json:"id,omitempty" `

    /*
        账号创建类型：1、通过短信创建，2、ISV批量导入，3、ISV OAuth创建     */
    Type  *int64 `json:"type,omitempty" `

    /*
        账号状态：1、启用，2、删除，3、禁用     */
    Status  *int64 `json:"status,omitempty" `

    /*
        记录的版本号     */
    Version  *int64 `json:"version,omitempty" `

    /*
        加密算法类型：1、代表单纯MD5，2：代表单一Salt的MD5，3、代表根据记录不同后的MD5     */
    LoginPwdEncryption  *int64 `json:"login_pwd_encryption,omitempty" `

    /*
        1男 2女     */
    Gender  *int64 `json:"gender,omitempty" `

    /*
        姓名     */
    Name  *string `json:"name,omitempty" `

    /*
        出生日期     */
    Birthday  *string `json:"birthday,omitempty" `

    /*
        旺旺     */
    Wangwang  *string `json:"wangwang,omitempty" `

    /*
        微信     */
    Weixin  *string `json:"weixin,omitempty" `

    /*
        TAOBAO = 1;WEIXIN = 2;WEIBO = 3;QQ = 4;     */
    OauthPlateform  *int64 `json:"oauth_plateform,omitempty" `

}

func (s *TaobaoOpenAccountTokenValidateOpenAccount) SetLoginId(v string) *TaobaoOpenAccountTokenValidateOpenAccount {
    s.LoginId = &v
    return s
}
func (s *TaobaoOpenAccountTokenValidateOpenAccount) SetCreateDeviceId(v string) *TaobaoOpenAccountTokenValidateOpenAccount {
    s.CreateDeviceId = &v
    return s
}
func (s *TaobaoOpenAccountTokenValidateOpenAccount) SetAlipayId(v string) *TaobaoOpenAccountTokenValidateOpenAccount {
    s.AlipayId = &v
    return s
}
func (s *TaobaoOpenAccountTokenValidateOpenAccount) SetCreateAppKey(v string) *TaobaoOpenAccountTokenValidateOpenAccount {
    s.CreateAppKey = &v
    return s
}
func (s *TaobaoOpenAccountTokenValidateOpenAccount) SetLocale(v string) *TaobaoOpenAccountTokenValidateOpenAccount {
    s.Locale = &v
    return s
}
func (s *TaobaoOpenAccountTokenValidateOpenAccount) SetBankCardNo(v string) *TaobaoOpenAccountTokenValidateOpenAccount {
    s.BankCardNo = &v
    return s
}
func (s *TaobaoOpenAccountTokenValidateOpenAccount) SetIsvAccountId(v string) *TaobaoOpenAccountTokenValidateOpenAccount {
    s.IsvAccountId = &v
    return s
}
func (s *TaobaoOpenAccountTokenValidateOpenAccount) SetEmail(v string) *TaobaoOpenAccountTokenValidateOpenAccount {
    s.Email = &v
    return s
}
func (s *TaobaoOpenAccountTokenValidateOpenAccount) SetAvatarUrl(v string) *TaobaoOpenAccountTokenValidateOpenAccount {
    s.AvatarUrl = &v
    return s
}
func (s *TaobaoOpenAccountTokenValidateOpenAccount) SetDomainId(v int64) *TaobaoOpenAccountTokenValidateOpenAccount {
    s.DomainId = &v
    return s
}
func (s *TaobaoOpenAccountTokenValidateOpenAccount) SetBankCardOwnerName(v string) *TaobaoOpenAccountTokenValidateOpenAccount {
    s.BankCardOwnerName = &v
    return s
}
func (s *TaobaoOpenAccountTokenValidateOpenAccount) SetDisplayName(v string) *TaobaoOpenAccountTokenValidateOpenAccount {
    s.DisplayName = &v
    return s
}
func (s *TaobaoOpenAccountTokenValidateOpenAccount) SetLoginPwdSalt(v string) *TaobaoOpenAccountTokenValidateOpenAccount {
    s.LoginPwdSalt = &v
    return s
}
func (s *TaobaoOpenAccountTokenValidateOpenAccount) SetLoginPwd(v string) *TaobaoOpenAccountTokenValidateOpenAccount {
    s.LoginPwd = &v
    return s
}
func (s *TaobaoOpenAccountTokenValidateOpenAccount) SetOpenId(v string) *TaobaoOpenAccountTokenValidateOpenAccount {
    s.OpenId = &v
    return s
}
func (s *TaobaoOpenAccountTokenValidateOpenAccount) SetMobile(v string) *TaobaoOpenAccountTokenValidateOpenAccount {
    s.Mobile = &v
    return s
}
func (s *TaobaoOpenAccountTokenValidateOpenAccount) SetCreateLocation(v string) *TaobaoOpenAccountTokenValidateOpenAccount {
    s.CreateLocation = &v
    return s
}
func (s *TaobaoOpenAccountTokenValidateOpenAccount) SetExtInfos(v string) *TaobaoOpenAccountTokenValidateOpenAccount {
    s.ExtInfos = &v
    return s
}
func (s *TaobaoOpenAccountTokenValidateOpenAccount) SetLoginPwdIntensity(v int64) *TaobaoOpenAccountTokenValidateOpenAccount {
    s.LoginPwdIntensity = &v
    return s
}
func (s *TaobaoOpenAccountTokenValidateOpenAccount) SetId(v int64) *TaobaoOpenAccountTokenValidateOpenAccount {
    s.Id = &v
    return s
}
func (s *TaobaoOpenAccountTokenValidateOpenAccount) SetType(v int64) *TaobaoOpenAccountTokenValidateOpenAccount {
    s.Type = &v
    return s
}
func (s *TaobaoOpenAccountTokenValidateOpenAccount) SetStatus(v int64) *TaobaoOpenAccountTokenValidateOpenAccount {
    s.Status = &v
    return s
}
func (s *TaobaoOpenAccountTokenValidateOpenAccount) SetVersion(v int64) *TaobaoOpenAccountTokenValidateOpenAccount {
    s.Version = &v
    return s
}
func (s *TaobaoOpenAccountTokenValidateOpenAccount) SetLoginPwdEncryption(v int64) *TaobaoOpenAccountTokenValidateOpenAccount {
    s.LoginPwdEncryption = &v
    return s
}
func (s *TaobaoOpenAccountTokenValidateOpenAccount) SetGender(v int64) *TaobaoOpenAccountTokenValidateOpenAccount {
    s.Gender = &v
    return s
}
func (s *TaobaoOpenAccountTokenValidateOpenAccount) SetName(v string) *TaobaoOpenAccountTokenValidateOpenAccount {
    s.Name = &v
    return s
}
func (s *TaobaoOpenAccountTokenValidateOpenAccount) SetBirthday(v string) *TaobaoOpenAccountTokenValidateOpenAccount {
    s.Birthday = &v
    return s
}
func (s *TaobaoOpenAccountTokenValidateOpenAccount) SetWangwang(v string) *TaobaoOpenAccountTokenValidateOpenAccount {
    s.Wangwang = &v
    return s
}
func (s *TaobaoOpenAccountTokenValidateOpenAccount) SetWeixin(v string) *TaobaoOpenAccountTokenValidateOpenAccount {
    s.Weixin = &v
    return s
}
func (s *TaobaoOpenAccountTokenValidateOpenAccount) SetOauthPlateform(v int64) *TaobaoOpenAccountTokenValidateOpenAccount {
    s.OauthPlateform = &v
    return s
}

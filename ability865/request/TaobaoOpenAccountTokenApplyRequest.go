package request

import (
        "topsdk/util"
    )

type TaobaoOpenAccountTokenApplyRequest struct {
    /*
        isv自己账号的唯一id     */
    IsvAccountId  *string `json:"isv_account_id,omitempty" required:"false" `
    /*
        ISV APP的登录态时长单位是秒     */
    LoginStateExpireIn  *int64 `json:"login_state_expire_in,omitempty" required:"false" `
    /*
        open account id     */
    OpenAccountId  *int64 `json:"open_account_id,omitempty" required:"false" `
    /*
        时间戳单位是毫秒     */
    TokenTimestamp  *int64 `json:"token_timestamp,omitempty" required:"false" `
    /*
        用于防重放的唯一id     */
    Uuid  *string `json:"uuid,omitempty" required:"false" `
    /*
        用于透传一些业务附加参数     */
    Ext  *map[string]interface{} `json:"ext,omitempty" required:"false" `
}

func (s *TaobaoOpenAccountTokenApplyRequest) SetIsvAccountId(v string) *TaobaoOpenAccountTokenApplyRequest {
    s.IsvAccountId = &v
    return s
}
func (s *TaobaoOpenAccountTokenApplyRequest) SetLoginStateExpireIn(v int64) *TaobaoOpenAccountTokenApplyRequest {
    s.LoginStateExpireIn = &v
    return s
}
func (s *TaobaoOpenAccountTokenApplyRequest) SetOpenAccountId(v int64) *TaobaoOpenAccountTokenApplyRequest {
    s.OpenAccountId = &v
    return s
}
func (s *TaobaoOpenAccountTokenApplyRequest) SetTokenTimestamp(v int64) *TaobaoOpenAccountTokenApplyRequest {
    s.TokenTimestamp = &v
    return s
}
func (s *TaobaoOpenAccountTokenApplyRequest) SetUuid(v string) *TaobaoOpenAccountTokenApplyRequest {
    s.Uuid = &v
    return s
}
func (s *TaobaoOpenAccountTokenApplyRequest) SetExt(v map[string]interface{}) *TaobaoOpenAccountTokenApplyRequest {
    s.Ext = &v
    return s
}

func (req *TaobaoOpenAccountTokenApplyRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.IsvAccountId != nil) {
        paramMap["isv_account_id"] = *req.IsvAccountId
    }
    if(req.LoginStateExpireIn != nil) {
        paramMap["login_state_expire_in"] = *req.LoginStateExpireIn
    }
    if(req.OpenAccountId != nil) {
        paramMap["open_account_id"] = *req.OpenAccountId
    }
    if(req.TokenTimestamp != nil) {
        paramMap["token_timestamp"] = *req.TokenTimestamp
    }
    if(req.Uuid != nil) {
        paramMap["uuid"] = *req.Uuid
    }
    if(req.Ext != nil) {
        paramMap["ext"] = util.ConvertStruct(*req.Ext)
    }
    return paramMap
}

func (req *TaobaoOpenAccountTokenApplyRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}
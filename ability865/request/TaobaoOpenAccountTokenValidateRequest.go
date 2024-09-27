package request


type TaobaoOpenAccountTokenValidateRequest struct {
    /*
        token     */
    ParamToken  *string `json:"param_token" required:"true" `
}

func (s *TaobaoOpenAccountTokenValidateRequest) SetParamToken(v string) *TaobaoOpenAccountTokenValidateRequest {
    s.ParamToken = &v
    return s
}

func (req *TaobaoOpenAccountTokenValidateRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.ParamToken != nil) {
        paramMap["param_token"] = *req.ParamToken
    }
    return paramMap
}

func (req *TaobaoOpenAccountTokenValidateRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}
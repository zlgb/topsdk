package ability865

import (
    "log"
    "errors"
    "topsdk"
    "topsdk/ability865/request"
    "topsdk/ability865/response"
    "topsdk/util"
)

type Ability865 struct {
    Client *topsdk.TopClient
}

func NewAbility865(client *topsdk.TopClient) *Ability865{
    return &Ability865{client}
}

/*
    open account token验证
*/
func (ability *Ability865) TaobaoOpenAccountTokenValidate(req *request.TaobaoOpenAccountTokenValidateRequest) (*response.TaobaoOpenAccountTokenValidateResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability865 topClient is nil")
    }
    var jsonStr,err = ability.Client.Execute("taobao.open.account.token.validate",req.ToMap(),req.ToFileMap())
    var respStruct = response.TaobaoOpenAccountTokenValidateResponse{}
    if(err != nil){
        log.Println("taobaoOpenAccountTokenValidate error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    申请免登Open Account Token
*/
func (ability *Ability865) TaobaoOpenAccountTokenApply(req *request.TaobaoOpenAccountTokenApplyRequest) (*response.TaobaoOpenAccountTokenApplyResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability865 topClient is nil")
    }
    var jsonStr,err = ability.Client.Execute("taobao.open.account.token.apply",req.ToMap(),req.ToFileMap())
    var respStruct = response.TaobaoOpenAccountTokenApplyResponse{}
    if(err != nil){
        log.Println("taobaoOpenAccountTokenApply error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}

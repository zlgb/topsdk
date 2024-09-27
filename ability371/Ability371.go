package ability371

import (
    "log"
    "errors"
    "topsdk"
    "topsdk/ability371/request"
    "topsdk/ability371/response"
    "topsdk/util"
)

type Ability371 struct {
    Client *topsdk.TopClient
}

func NewAbility371(client *topsdk.TopClient) *Ability371{
    return &Ability371{client}
}

/*
    淘宝客-公用-阿里妈妈推广券详情查询
*/
func (ability *Ability371) TaobaoTbkCouponGet(req *request.TaobaoTbkCouponGetRequest) (*response.TaobaoTbkCouponGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability371 topClient is nil")
    }
    var jsonStr,err = ability.Client.Execute("taobao.tbk.coupon.get",req.ToMap(),req.ToFileMap())
    var respStruct = response.TaobaoTbkCouponGetResponse{}
    if(err != nil){
        log.Println("taobaoTbkCouponGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    淘宝客-公用-淘宝客商品详情查询(简版)
*/
func (ability *Ability371) TaobaoTbkItemInfoGet(req *request.TaobaoTbkItemInfoGetRequest) (*response.TaobaoTbkItemInfoGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability371 topClient is nil")
    }
    var jsonStr,err = ability.Client.Execute("taobao.tbk.item.info.get",req.ToMap(),req.ToFileMap())
    var respStruct = response.TaobaoTbkItemInfoGetResponse{}
    if(err != nil){
        log.Println("taobaoTbkItemInfoGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    淘宝客商品详情查询（简版）（临时接口）
*/
func (ability *Ability371) TaobaoTbkItemInfoTemporaryGet(req *request.TaobaoTbkItemInfoTemporaryGetRequest) (*response.TaobaoTbkItemInfoTemporaryGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability371 topClient is nil")
    }
    var jsonStr,err = ability.Client.Execute("taobao.tbk.item.info.temporary.get",req.ToMap(),req.ToFileMap())
    var respStruct = response.TaobaoTbkItemInfoTemporaryGetResponse{}
    if(err != nil){
        log.Println("taobaoTbkItemInfoTemporaryGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    淘宝客-公用-淘宝客商品详情查询升级版（简易版）
*/
func (ability *Ability371) TaobaoTbkItemInfoUpgradeGet(req *request.TaobaoTbkItemInfoUpgradeGetRequest) (*response.TaobaoTbkItemInfoUpgradeGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability371 topClient is nil")
    }
    var jsonStr,err = ability.Client.Execute("taobao.tbk.item.info.upgrade.get",req.ToMap(),req.ToFileMap())
    var respStruct = response.TaobaoTbkItemInfoUpgradeGetResponse{}
    if(err != nil){
        log.Println("taobaoTbkItemInfoUpgradeGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}

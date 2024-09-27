package defaultability

import (
    "log"
    "errors"
    "topsdk"
    "topsdk/defaultability/request"
    "topsdk/defaultability/response"
    "topsdk/util"
)

type Defaultability struct {
    Client *topsdk.TopClient
}

func NewDefaultability(client *topsdk.TopClient) *Defaultability{
    return &Defaultability{client}
}

/*
    关键词过滤匹配
*/
func (ability *Defaultability) TaobaoKfcKeywordSearch(req *request.TaobaoKfcKeywordSearchRequest,session string) (*response.TaobaoKfcKeywordSearchResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.kfc.keyword.search",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoKfcKeywordSearchResponse{}
    if(err != nil){
        log.Println("taobaoKfcKeywordSearch error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    淘宝客-推广者-物料id列表查询
*/
func (ability *Defaultability) TaobaoTbkOptimusTouMaterialIdsGet(req *request.TaobaoTbkOptimusTouMaterialIdsGetRequest) (*response.TaobaoTbkOptimusTouMaterialIdsGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.Execute("taobao.tbk.optimus.tou.material.ids.get",req.ToMap(),req.ToFileMap())
    var respStruct = response.TaobaoTbkOptimusTouMaterialIdsGetResponse{}
    if(err != nil){
        log.Println("taobaoTbkOptimusTouMaterialIdsGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    淘宝客-推广者-物料搜索升级版
*/
func (ability *Defaultability) TaobaoTbkDgMaterialOptionalUpgrade(req *request.TaobaoTbkDgMaterialOptionalUpgradeRequest) (*response.TaobaoTbkDgMaterialOptionalUpgradeResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.Execute("taobao.tbk.dg.material.optional.upgrade",req.ToMap(),req.ToFileMap())
    var respStruct = response.TaobaoTbkDgMaterialOptionalUpgradeResponse{}
    if(err != nil){
        log.Println("taobaoTbkDgMaterialOptionalUpgrade error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    淘宝客-推广者-物料精选升级版
*/
func (ability *Defaultability) TaobaoTbkDgMaterialRecommend(req *request.TaobaoTbkDgMaterialRecommendRequest) (*response.TaobaoTbkDgMaterialRecommendResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.Execute("taobao.tbk.dg.material.recommend",req.ToMap(),req.ToFileMap())
    var respStruct = response.TaobaoTbkDgMaterialRecommendResponse{}
    if(err != nil){
        log.Println("taobaoTbkDgMaterialRecommend error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}

package ability372

import (
    "log"
    "errors"
    "topsdk"
    "topsdk/ability372/request"
    "topsdk/ability372/response"
    "topsdk/util"
)

type Ability372 struct {
    Client *topsdk.TopClient
}

func NewAbility372(client *topsdk.TopClient) *Ability372{
    return &Ability372{client}
}

/*
    淘宝客-公用-店铺关联推荐
*/
func (ability *Ability372) TaobaoTbkShopRecommendGet(req *request.TaobaoTbkShopRecommendGetRequest) (*response.TaobaoTbkShopRecommendGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability372 topClient is nil")
    }
    var jsonStr,err = ability.Client.Execute("taobao.tbk.shop.recommend.get",req.ToMap(),req.ToFileMap())
    var respStruct = response.TaobaoTbkShopRecommendGetResponse{}
    if(err != nil){
        log.Println("taobaoTbkShopRecommendGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}

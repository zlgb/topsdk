package response

import (
    "topsdk/defaultability/domain"
)

type TaobaoTbkDgMaterialRecommendResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        resultList
    */
    ResultList  []domain.TaobaoTbkDgMaterialRecommendMapData `json:"result_list,omitempty" `
    /*
        推荐信息-是否抄底
    */
    IsDefault  string `json:"is_default,omitempty" `
    /*
        商品总数-该选品库收藏夹id及官方物料id对应的商品数量
    */
    TotalCount  int64 `json:"total_count,omitempty" `
    /*
        uvid结果信息，传入但未使用uvid时会返回原因
    */
    UvidMsg  string `json:"uvid_msg,omitempty" `
}

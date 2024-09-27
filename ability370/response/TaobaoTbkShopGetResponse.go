package response

import (
    "topsdk/ability370/domain"
)

type TaobaoTbkShopGetResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        淘宝客店铺
    */
    Results  []domain.TaobaoTbkShopGetNTbkShop `json:"results,omitempty" `
    /*
        搜索到符合条件的结果总数
    */
    TotalResults  int64 `json:"total_results,omitempty" `
}

package response

import (
    "topsdk/ability371/domain"
)

type TaobaoTbkItemInfoUpgradeGetResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        仅淘宝客商品，字段值根据API赋权等级输出
    */
    Results  []domain.TaobaoTbkItemInfoUpgradeGetTbkItemDetail `json:"results,omitempty" `
}

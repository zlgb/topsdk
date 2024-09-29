package request

import (
        "topsdk/ability414/domain"
        "topsdk/util"
    )

type TaobaoTbkOrderRefundGetRequest struct {
    /*
        全量维权订单查询入参     */
    PublisherRefundOrderQueryOption  *domain.TaobaoTbkOrderRefundGetPublisherRefundOrderQueryOption `json:"publisher_refund_order_query_option,omitempty" required:"false" `
}

func (s *TaobaoTbkOrderRefundGetRequest) SetPublisherRefundOrderQueryOption(v domain.TaobaoTbkOrderRefundGetPublisherRefundOrderQueryOption) *TaobaoTbkOrderRefundGetRequest {
    s.PublisherRefundOrderQueryOption = &v
    return s
}

func (req *TaobaoTbkOrderRefundGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.PublisherRefundOrderQueryOption != nil) {
        paramMap["publisher_refund_order_query_option"] = util.ConvertStruct(*req.PublisherRefundOrderQueryOption)
    }
    return paramMap
}

func (req *TaobaoTbkOrderRefundGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}
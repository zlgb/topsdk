package domain


type TaobaoTbkOrderDetailsGetPublisherOrderDto struct {
    /*
        淘宝付款时间。解释：订单在淘宝付款的时间     */
    TbPaidTime  *string `json:"tb_paid_time,omitempty" `

    /*
        付款时间。解释：订单付款的时间，该时间同步淘宝，可能会略晚于买家在淘宝的订单付款时间  （预售订单完尾款支付后，付款时间会自动更新为尾款支付的时间）     */
    TkPaidTime  *string `json:"tk_paid_time,omitempty" `

    /*
        结算金额。解释：买家确认收货的付款金额（不包含运费金额）     */
    PayPrice  *string `json:"pay_price,omitempty" `

    /*
        结算预估收入。解释：结算预付收入=结算预估佣金收入+结算预估补贴收入     */
    PubShareFee  *string `json:"pub_share_fee,omitempty" `

    /*
        子订单号。解释：买家通过购物车购买的每个商品对应的订单编号，此订单编号并未在淘宝买家后台透出，若平台类型为淘宝、口碑、饿了么等，则订单编号即为淘宝子订单编号、口碑订单编号、饿了么订单编号，以此类推     */
    TradeId  *string `json:"trade_id,omitempty" `

    /*
        推广者身份。解释：二方：佣金收益的第一归属者；三方：从其他淘宝客佣金中进行分成的推广者     */
    TkOrderRole  *int64 `json:"tk_order_role,omitempty" `

    /*
        结算时间。解释：订单确认收货后且商家完成佣金支付的时间     */
    TkEarningTime  *string `json:"tk_earning_time,omitempty" `

    /*
        推广位ID。解释：“推广管理-推广位管理”中的pid中的最后一段数字，如pid=mm_1_2_3中的“3”这段数字     */
    AdzoneId  *int64 `json:"adzone_id,omitempty" `

    /*
        佣金分成比率。解释：从佣金中分得的收益比率（含平台技术服务费比率）     */
    PubShareRate  *string `json:"pub_share_rate,omitempty" `

    /*
        unid（不对外开放）     */
    Unid  *string `json:"unid,omitempty" `

    /*
        维权标签。解释：若该订单产生了维权退款，则会打上“维权单”的提示。0 含义为非维权、1 含义为维权订单     */
    RefundTag  *int64 `json:"refund_tag,omitempty" `

    /*
        补贴比率。解释：指该笔订单上各类型补贴的补贴比率总和。补贴比率=a补贴比率+b补贴比率+…。举例：天猫补贴1%、飞猪补贴1%等，则该数值显示为2%     */
    SubsidyRate  *string `json:"subsidy_rate,omitempty" `

    /*
        佣金提成。解释：佣金提成=佣金比率*佣金分成比率。指实际获得的佣金收益比率（含平台技术服务费）     */
    TkTotalRate  *string `json:"tk_total_rate,omitempty" `

    /*
        类目名称。解释：商品所属的一级类目名称     */
    ItemCategoryName  *string `json:"item_category_name,omitempty" `

    /*
        掌柜旺旺     */
    SellerNick  *string `json:"seller_nick,omitempty" `

    /*
        推广者的账号id     */
    PubId  *int64 `json:"pub_id,omitempty" `

    /*
        平台技术服务费比率。解释：指该笔订单推广者在需支付给淘宝联盟的所有技术服务费比率总和。平台技术服务费比率=技术服务费比率a+技术服务费比率b+…。（特别说明：补贴从2023-03-10之后都不再收取补贴技术服务费了，故本字段公式中不再提及补贴相关)     */
    AlimamaRate  *string `json:"alimama_rate,omitempty" `

    /*
        补贴类型。解释：各补贴类型的类型名称、补贴比率、补贴金额、单笔补贴上限、补贴分成比率的详细说明，如有多个补贴会一并告知，举例：淘宝特价版（补贴比率：1%，补贴金额1.00元，单笔补贴上限金额1000.00元，补贴分成比率：100.00%）、飞猪补贴（补贴比率：1%，补贴金额1.00元，单笔补贴上限金额1000.00元，补贴分成比率：100.00%）     */
    SubsidyType  *string `json:"subsidy_type,omitempty" `

    /*
        商品图片     */
    ItemImg  *string `json:"item_img,omitempty" `

    /*
        付款预估收入。解释：付款预付收入=付款预估佣金收入+付款预估补贴收入     */
    PubSharePreFee  *string `json:"pub_share_pre_fee,omitempty" `

    /*
        付款金额。解释：买家拍下并付款金额（含预售订单定金金额，但不包含运费金额）；当预售订单完成尾款支付后，付款金额会自动更新为订单总金额     */
    AlipayTotalPrice  *string `json:"alipay_total_price,omitempty" `

    /*
        商品标题     */
    ItemTitle  *string `json:"item_title,omitempty" `

    /*
        媒体名称。解释：“推广管理-媒体备案管理”中的媒体名称     */
    SiteName  *string `json:"site_name,omitempty" `

    /*
        商品数量     */
    ItemNum  *int64 `json:"item_num,omitempty" `

    /*
        补贴金额。解释：指该笔订单上各类补贴的补贴金额总和。补贴金额=a补贴金额+b补贴金额+…=结算金额*a补贴比率+结算金额*b补贴比率+…。若（结算金额*a补贴比率）＞补贴类型a单笔补贴上限，则a补贴金额=补贴类型a单笔补贴上限，b补贴金额等其他补贴金额同理     */
    SubsidyFee  *string `json:"subsidy_fee,omitempty" `

    /*
        平台技术服务费。解释：指该笔订单推广者在需支付给淘宝联盟的所有技术服务费费用总和。目前包含以下两类： 1、技术服务费说明：通过淘宝联盟平台推广所需支付的基础技术服务费；2、渠道专项服务费说明：开通渠道管理权限的推广者进行推广需要支付给淘宝联盟的专项技术服务费用。           平台技术服务费=付款金额*佣金比率*平台技术服务费比率。注意：若订单已结算则会按结算金额计算。（特别说明：补贴从2023-03-10之后都不再收取补贴技术服务费了，故本字段公式中不再提及补贴相关)     */
    AlimamaShareFee  *string `json:"alimama_share_fee,omitempty" `

    /*
        订单编号。解释：买家在购买后台显示的订单编号，若平台类型为淘宝、饿了么等，则订单编号即为淘宝订单编号、饿了么订单编号，以此类推     */
    TradeParentId  *string `json:"trade_parent_id,omitempty" `

    /*
        平台类型。解释：订单所属平台类型，包括天猫、淘宝、聚划算、口碑等     */
    OrderType  *string `json:"order_type,omitempty" `

    /*
        创建时间。解释：订单创建的时间，该时间同步淘宝，可能会略晚于买家在淘宝的订单创建时间     */
    TkCreateTime  *string `json:"tk_create_time,omitempty" `

    /*
        产品类型。解释：若订单属于特定的产品类型，会进行说明，举例：该订单属于“联盟超级活动”、“品牌精选推广”、“自主推广”等，若无说明仅为“--”则为普通类型订单     */
    FlowSource  *string `json:"flow_source,omitempty" `

    /*
        成交平台。解释：成交来自于PC或无线     */
    TerminalType  *string `json:"terminal_type,omitempty" `

    /*
        点击时间。解释：通过推广链接达到商品、店铺详情页的点击时间     */
    ClickTime  *string `json:"click_time,omitempty" `

    /*
        订单状态。状态类型：已付款（指订单已付款，但还未确认收货）、已收货（指订单已确认收货，但商家佣金未支付）、已结算（指订单已确认收货，且商家佣金已支付成功）、已失效（指订单关闭/订单佣金小于0.01元，订单关闭主要有：①买家超时未付款；②买家付款前，买家/商家取消了订单；③订单付款后发起售中退款成功；④仅支付定金，到期未支付尾款的预售订单）。注意：当订单商家有实收货款（含交易成功和交易关闭订单），联盟以商家实收货款进行佣金结算(满返/买返订单等特殊订单除外)。出参数字代表：3:订单结算，12:订单付款，13:订单失效，14:订单成功     */
    TkStatus  *int64 `json:"tk_status,omitempty" `

    /*
        商品单价     */
    ItemPrice  *string `json:"item_price,omitempty" `

    /*
        商品id     */
    ItemId  *string `json:"item_id,omitempty" `

    /*
        推广位名称。解释：“推广管理-推广位管理”中的推广位名称     */
    AdzoneName  *string `json:"adzone_name,omitempty" `

    /*
        佣金比率。解释：商品的佣金比率（含平台技术服务费比率，但不包含平台专项服务费比率）（特别说明：若想知道属于您名下的整体佣金比率，则整体佣金比率=佣金比率+平台专项服务费比率，为了简化展示报表，不再增加整体佣金收入字段，有需要的淘宝客可自行加和，举例：若整体佣金比率为10%，平台专项服务费比率为3%，则此处商品佣金比率显示为7%，即10%-3%=7%）     */
    TotalCommissionRate  *string `json:"total_commission_rate,omitempty" `

    /*
        商品链接     */
    ItemLink  *string `json:"item_link,omitempty" `

    /*
        媒体ID。解释：“推广管理-媒体备案管理”中的媒体ID，也是pid中的第二段数字，如pid=mm_1_2_3中的“2”这段数字     */
    SiteId  *int64 `json:"site_id,omitempty" `

    /*
        店铺名称     */
    SellerShopTitle  *string `json:"seller_shop_title,omitempty" `

    /*
        收入比率。解释：收入比率=佣金比率+补贴比率     */
    IncomeRate  *string `json:"income_rate,omitempty" `

    /*
        佣金金额。解释：如果该订单还未结算，则佣金金额=付款金额*佣金比率；如果该订单已经结算，则佣金金额=结算金额*佣金比率     */
    TotalCommissionFee  *string `json:"total_commission_fee,omitempty" `

    /*
        会员运营ID。解释：会员运营管理功能中的会员运营ID，若该订单来自于会员的购买，则会展示对应会员运营ID     */
    SpecialId  *int64 `json:"special_id,omitempty" `

    /*
        渠道关系id。解释：渠道管理功能中的渠道关系ID，若该订单来自于渠道方的推广，则会展示对应渠道关系ID     */
    RelationId  *int64 `json:"relation_id,omitempty" `

    /*
        定金付款金额。解释：预售时期，预售订单支付的定金金额     */
    DepositPrice  *string `json:"deposit_price,omitempty" `

    /*
        定金淘宝付款时间。解释：预售时期，预售订单在淘宝支付定金的付款时间     */
    TbDepositTime  *string `json:"tb_deposit_time,omitempty" `

    /*
        定金付款时间。解释：预售时期，预售订单支付定金的付款时间，该时间同步淘宝，可能会略晚于买家在淘宝的订单创建时间     */
    TkDepositTime  *string `json:"tk_deposit_time,omitempty" `

    /*
        非电商淘系子订单号     */
    TpOrderId  *string `json:"tp_order_id,omitempty" `

    /*
        服务费信息（字段已废弃）     */
    ServiceFeeDtoList  *[]TaobaoTbkOrderDetailsGetServiceFeeDTO `json:"service_fee_dto_list,omitempty" `

    /*
        类目名称。解释：商品所属的二级类目名称     */
    ItemCategoryLevel2Name  *string `json:"item_category_level2_name,omitempty" `

    /*
        营销类型。解释：该字段中视订单情况有单个或多个类型告知。 举例：联盟超级活动-优品、特价版客户端锁粉、特价版客户端推广等     */
    MarketingType  *string `json:"marketing_type,omitempty" `

    /*
        订单更新时间     */
    ModifiedTime  *string `json:"modified_time,omitempty" `

    /*
        专用（不对外开放）     */
    TalentPid  *string `json:"talent_pid,omitempty" `

    /*
        买家拍下金额（不包含运费金额）     */
    TbGmvTotalPrice  *string `json:"tb_gmv_total_price,omitempty" `

    /*
        管理member新商品ID后段     */
    ExtraMktId  *string `json:"extra_mkt_id,omitempty" `

    /*
        流量通untts（默认无，限定开放）     */
    Untts  *string `json:"untts,omitempty" `

    /*
        付款预估佣金收入。解释：付款预估佣金收入=付款金额*佣金提成。以买家付款金额为基数，预估您可能获得的付款佣金收入，包含平台技术服务费金额（最终发钱时会减去平台技术服务费）。注意：因买家退款等原因，可能与结算预估佣金收入不一致。（特别说明：若想知道属于您名下的整体佣金收入，则整体付款佣金收入=付款预估佣金收入+平台专项服务费，为了简化展示报表，不再增加整体佣金收入字段，有需要的淘宝客可自行加和）     */
    PubSharePreFeeForCommission  *string `json:"pub_share_pre_fee_for_commission,omitempty" `

    /*
        结算预估佣金收入。解释：结算预估佣金收入=付款金额*佣金提成。以买家确认收货的付款金额为基数，预估您可能获得的结算佣金收入，包含技术服务费金额（最终发钱时会减去平台技术服务费）。注意：因买家退款、您违规推广等原因，可能与您最终收入不一致，最终收入以月结后您实际收到的为准。（特别说明：若想知道属于您名下的整体佣金收入，则整体结算佣金收入=结算预估佣金收入+平台专项服务费，为了简化展示报表，不再增加整体佣金收入字段，有需要的淘宝客可自行加和）     */
    PubShareFeeForCommission  *string `json:"pub_share_fee_for_commission,omitempty" `

    /*
        补贴金额明细节点。解释：各补贴类型的类型名称、补贴比率、补贴金额、单笔补贴上限、补贴分成比率的详细说明     */
    SubsidyInfoDtoList  *[]TaobaoTbkOrderDetailsGetSubsidyDetailDTO `json:"subsidy_info_dto_list,omitempty" `

    /*
        补贴分成比率。解释：从补贴中分得的收益比率     */
    PubShareRateForSdy  *string `json:"pub_share_rate_for_sdy,omitempty" `

    /*
        补贴提成。解释：补贴提成=补贴比率*补贴分成比率。指实际获得的补贴收益比率     */
    TkTotalRateForSdy  *string `json:"tk_total_rate_for_sdy,omitempty" `

    /*
        付款预估补贴收入。解释：指以买家付款金额为基数，预估您可能获得的补贴收入。付款预估补贴收入=（付款金额*a补贴比率+付款金额*b补贴比率+……）*补贴分成比率。如果付款金额*a补贴比率＞补贴类型a单笔订单补贴上限，则付款金额*a补贴比率的值取补贴类型a单笔订单补贴上限，b补贴金额等其他类型补贴金额同理。注意：因买家退款等原因，可能与结算预估补贴收入不一致     */
    PubSharePreFeeForSdy  *string `json:"pub_share_pre_fee_for_sdy,omitempty" `

    /*
        结算预估补贴收入。解释：以买家确认收货的付款金额为基数，预估您可能获得的结算补贴收入。结算预估补贴收入=（结算金额*a补贴比率+结算金额*b补贴比率+……）*补贴分成比率。如果结算金额*a补贴比率＞补贴类型a单笔订单补贴上限，则结算金额*a补贴比率的值取补贴类型a单笔订单补贴上限，b补贴金额等其他类型补贴金额同理。注意：因买家退款、您违规推广等原因，可能与您最终收入不一致，最终收入以月结后您实际收到的为准     */
    PubShareFeeForSdy  *string `json:"pub_share_fee_for_sdy,omitempty" `

    /*
        平台技术服务费明细节点。解释：各项平台技术服务费类型的类型名称、扣费比率、扣费金额的详细说明     */
    AlimmShareInfoDto  *TaobaoTbkOrderDetailsGetAlimmShareInfoDTO `json:"alimm_share_info_dto,omitempty" `

    /*
        平台专项服务费比率。解释：指该笔订单推广者在各种特殊场景下需支付给淘宝联盟的所有专项服务费比率总和。平台专项服务费比率=专项服务费比率a+专项服务费比率b+…     */
    PlatformSpecialServiceRate  *string `json:"platform_special_service_rate,omitempty" `

    /*
        平台专项服务费。解释：指该笔订单推广者在各种特殊场景下需支付给淘宝联盟的所有专项服务费用总和。目前包含以下两类：1、内容专项服务费说明：内容场景专项技术服务费，内容推广者在内容场景进行推广需要支付给淘宝联盟的专项技术服务费用；2、流量专项服务费说明：流量通产品合作场景专项技术服务费，推广者在流量通产品能力下，进行流量投放推广需要支付给淘宝联盟的专项技术服务费用。           平台专项服务费=付款金额*平台专项服务费比率。注意：若订单已结算则会按结算金额计算     */
    PlatformSpecialServiceFee  *string `json:"platform_special_service_fee,omitempty" `

    /*
        平台专项服务费明细节点。解释：各项平台专项服务费类型的类型名称、扣费比率、扣费金额的详细说明     */
    PlatformSpecialShareInfoDto  *TaobaoTbkOrderDetailsGetPlatformSpecialShareInfoDTO `json:"platform_special_share_info_dto,omitempty" `

}

func (s *TaobaoTbkOrderDetailsGetPublisherOrderDto) SetTbPaidTime(v string) *TaobaoTbkOrderDetailsGetPublisherOrderDto {
    s.TbPaidTime = &v
    return s
}
func (s *TaobaoTbkOrderDetailsGetPublisherOrderDto) SetTkPaidTime(v string) *TaobaoTbkOrderDetailsGetPublisherOrderDto {
    s.TkPaidTime = &v
    return s
}
func (s *TaobaoTbkOrderDetailsGetPublisherOrderDto) SetPayPrice(v string) *TaobaoTbkOrderDetailsGetPublisherOrderDto {
    s.PayPrice = &v
    return s
}
func (s *TaobaoTbkOrderDetailsGetPublisherOrderDto) SetPubShareFee(v string) *TaobaoTbkOrderDetailsGetPublisherOrderDto {
    s.PubShareFee = &v
    return s
}
func (s *TaobaoTbkOrderDetailsGetPublisherOrderDto) SetTradeId(v string) *TaobaoTbkOrderDetailsGetPublisherOrderDto {
    s.TradeId = &v
    return s
}
func (s *TaobaoTbkOrderDetailsGetPublisherOrderDto) SetTkOrderRole(v int64) *TaobaoTbkOrderDetailsGetPublisherOrderDto {
    s.TkOrderRole = &v
    return s
}
func (s *TaobaoTbkOrderDetailsGetPublisherOrderDto) SetTkEarningTime(v string) *TaobaoTbkOrderDetailsGetPublisherOrderDto {
    s.TkEarningTime = &v
    return s
}
func (s *TaobaoTbkOrderDetailsGetPublisherOrderDto) SetAdzoneId(v int64) *TaobaoTbkOrderDetailsGetPublisherOrderDto {
    s.AdzoneId = &v
    return s
}
func (s *TaobaoTbkOrderDetailsGetPublisherOrderDto) SetPubShareRate(v string) *TaobaoTbkOrderDetailsGetPublisherOrderDto {
    s.PubShareRate = &v
    return s
}
func (s *TaobaoTbkOrderDetailsGetPublisherOrderDto) SetUnid(v string) *TaobaoTbkOrderDetailsGetPublisherOrderDto {
    s.Unid = &v
    return s
}
func (s *TaobaoTbkOrderDetailsGetPublisherOrderDto) SetRefundTag(v int64) *TaobaoTbkOrderDetailsGetPublisherOrderDto {
    s.RefundTag = &v
    return s
}
func (s *TaobaoTbkOrderDetailsGetPublisherOrderDto) SetSubsidyRate(v string) *TaobaoTbkOrderDetailsGetPublisherOrderDto {
    s.SubsidyRate = &v
    return s
}
func (s *TaobaoTbkOrderDetailsGetPublisherOrderDto) SetTkTotalRate(v string) *TaobaoTbkOrderDetailsGetPublisherOrderDto {
    s.TkTotalRate = &v
    return s
}
func (s *TaobaoTbkOrderDetailsGetPublisherOrderDto) SetItemCategoryName(v string) *TaobaoTbkOrderDetailsGetPublisherOrderDto {
    s.ItemCategoryName = &v
    return s
}
func (s *TaobaoTbkOrderDetailsGetPublisherOrderDto) SetSellerNick(v string) *TaobaoTbkOrderDetailsGetPublisherOrderDto {
    s.SellerNick = &v
    return s
}
func (s *TaobaoTbkOrderDetailsGetPublisherOrderDto) SetPubId(v int64) *TaobaoTbkOrderDetailsGetPublisherOrderDto {
    s.PubId = &v
    return s
}
func (s *TaobaoTbkOrderDetailsGetPublisherOrderDto) SetAlimamaRate(v string) *TaobaoTbkOrderDetailsGetPublisherOrderDto {
    s.AlimamaRate = &v
    return s
}
func (s *TaobaoTbkOrderDetailsGetPublisherOrderDto) SetSubsidyType(v string) *TaobaoTbkOrderDetailsGetPublisherOrderDto {
    s.SubsidyType = &v
    return s
}
func (s *TaobaoTbkOrderDetailsGetPublisherOrderDto) SetItemImg(v string) *TaobaoTbkOrderDetailsGetPublisherOrderDto {
    s.ItemImg = &v
    return s
}
func (s *TaobaoTbkOrderDetailsGetPublisherOrderDto) SetPubSharePreFee(v string) *TaobaoTbkOrderDetailsGetPublisherOrderDto {
    s.PubSharePreFee = &v
    return s
}
func (s *TaobaoTbkOrderDetailsGetPublisherOrderDto) SetAlipayTotalPrice(v string) *TaobaoTbkOrderDetailsGetPublisherOrderDto {
    s.AlipayTotalPrice = &v
    return s
}
func (s *TaobaoTbkOrderDetailsGetPublisherOrderDto) SetItemTitle(v string) *TaobaoTbkOrderDetailsGetPublisherOrderDto {
    s.ItemTitle = &v
    return s
}
func (s *TaobaoTbkOrderDetailsGetPublisherOrderDto) SetSiteName(v string) *TaobaoTbkOrderDetailsGetPublisherOrderDto {
    s.SiteName = &v
    return s
}
func (s *TaobaoTbkOrderDetailsGetPublisherOrderDto) SetItemNum(v int64) *TaobaoTbkOrderDetailsGetPublisherOrderDto {
    s.ItemNum = &v
    return s
}
func (s *TaobaoTbkOrderDetailsGetPublisherOrderDto) SetSubsidyFee(v string) *TaobaoTbkOrderDetailsGetPublisherOrderDto {
    s.SubsidyFee = &v
    return s
}
func (s *TaobaoTbkOrderDetailsGetPublisherOrderDto) SetAlimamaShareFee(v string) *TaobaoTbkOrderDetailsGetPublisherOrderDto {
    s.AlimamaShareFee = &v
    return s
}
func (s *TaobaoTbkOrderDetailsGetPublisherOrderDto) SetTradeParentId(v string) *TaobaoTbkOrderDetailsGetPublisherOrderDto {
    s.TradeParentId = &v
    return s
}
func (s *TaobaoTbkOrderDetailsGetPublisherOrderDto) SetOrderType(v string) *TaobaoTbkOrderDetailsGetPublisherOrderDto {
    s.OrderType = &v
    return s
}
func (s *TaobaoTbkOrderDetailsGetPublisherOrderDto) SetTkCreateTime(v string) *TaobaoTbkOrderDetailsGetPublisherOrderDto {
    s.TkCreateTime = &v
    return s
}
func (s *TaobaoTbkOrderDetailsGetPublisherOrderDto) SetFlowSource(v string) *TaobaoTbkOrderDetailsGetPublisherOrderDto {
    s.FlowSource = &v
    return s
}
func (s *TaobaoTbkOrderDetailsGetPublisherOrderDto) SetTerminalType(v string) *TaobaoTbkOrderDetailsGetPublisherOrderDto {
    s.TerminalType = &v
    return s
}
func (s *TaobaoTbkOrderDetailsGetPublisherOrderDto) SetClickTime(v string) *TaobaoTbkOrderDetailsGetPublisherOrderDto {
    s.ClickTime = &v
    return s
}
func (s *TaobaoTbkOrderDetailsGetPublisherOrderDto) SetTkStatus(v int64) *TaobaoTbkOrderDetailsGetPublisherOrderDto {
    s.TkStatus = &v
    return s
}
func (s *TaobaoTbkOrderDetailsGetPublisherOrderDto) SetItemPrice(v string) *TaobaoTbkOrderDetailsGetPublisherOrderDto {
    s.ItemPrice = &v
    return s
}
func (s *TaobaoTbkOrderDetailsGetPublisherOrderDto) SetItemId(v string) *TaobaoTbkOrderDetailsGetPublisherOrderDto {
    s.ItemId = &v
    return s
}
func (s *TaobaoTbkOrderDetailsGetPublisherOrderDto) SetAdzoneName(v string) *TaobaoTbkOrderDetailsGetPublisherOrderDto {
    s.AdzoneName = &v
    return s
}
func (s *TaobaoTbkOrderDetailsGetPublisherOrderDto) SetTotalCommissionRate(v string) *TaobaoTbkOrderDetailsGetPublisherOrderDto {
    s.TotalCommissionRate = &v
    return s
}
func (s *TaobaoTbkOrderDetailsGetPublisherOrderDto) SetItemLink(v string) *TaobaoTbkOrderDetailsGetPublisherOrderDto {
    s.ItemLink = &v
    return s
}
func (s *TaobaoTbkOrderDetailsGetPublisherOrderDto) SetSiteId(v int64) *TaobaoTbkOrderDetailsGetPublisherOrderDto {
    s.SiteId = &v
    return s
}
func (s *TaobaoTbkOrderDetailsGetPublisherOrderDto) SetSellerShopTitle(v string) *TaobaoTbkOrderDetailsGetPublisherOrderDto {
    s.SellerShopTitle = &v
    return s
}
func (s *TaobaoTbkOrderDetailsGetPublisherOrderDto) SetIncomeRate(v string) *TaobaoTbkOrderDetailsGetPublisherOrderDto {
    s.IncomeRate = &v
    return s
}
func (s *TaobaoTbkOrderDetailsGetPublisherOrderDto) SetTotalCommissionFee(v string) *TaobaoTbkOrderDetailsGetPublisherOrderDto {
    s.TotalCommissionFee = &v
    return s
}
func (s *TaobaoTbkOrderDetailsGetPublisherOrderDto) SetSpecialId(v int64) *TaobaoTbkOrderDetailsGetPublisherOrderDto {
    s.SpecialId = &v
    return s
}
func (s *TaobaoTbkOrderDetailsGetPublisherOrderDto) SetRelationId(v int64) *TaobaoTbkOrderDetailsGetPublisherOrderDto {
    s.RelationId = &v
    return s
}
func (s *TaobaoTbkOrderDetailsGetPublisherOrderDto) SetDepositPrice(v string) *TaobaoTbkOrderDetailsGetPublisherOrderDto {
    s.DepositPrice = &v
    return s
}
func (s *TaobaoTbkOrderDetailsGetPublisherOrderDto) SetTbDepositTime(v string) *TaobaoTbkOrderDetailsGetPublisherOrderDto {
    s.TbDepositTime = &v
    return s
}
func (s *TaobaoTbkOrderDetailsGetPublisherOrderDto) SetTkDepositTime(v string) *TaobaoTbkOrderDetailsGetPublisherOrderDto {
    s.TkDepositTime = &v
    return s
}
func (s *TaobaoTbkOrderDetailsGetPublisherOrderDto) SetTpOrderId(v string) *TaobaoTbkOrderDetailsGetPublisherOrderDto {
    s.TpOrderId = &v
    return s
}
func (s *TaobaoTbkOrderDetailsGetPublisherOrderDto) SetServiceFeeDtoList(v []TaobaoTbkOrderDetailsGetServiceFeeDTO) *TaobaoTbkOrderDetailsGetPublisherOrderDto {
    s.ServiceFeeDtoList = &v
    return s
}
func (s *TaobaoTbkOrderDetailsGetPublisherOrderDto) SetItemCategoryLevel2Name(v string) *TaobaoTbkOrderDetailsGetPublisherOrderDto {
    s.ItemCategoryLevel2Name = &v
    return s
}
func (s *TaobaoTbkOrderDetailsGetPublisherOrderDto) SetMarketingType(v string) *TaobaoTbkOrderDetailsGetPublisherOrderDto {
    s.MarketingType = &v
    return s
}
func (s *TaobaoTbkOrderDetailsGetPublisherOrderDto) SetModifiedTime(v string) *TaobaoTbkOrderDetailsGetPublisherOrderDto {
    s.ModifiedTime = &v
    return s
}
func (s *TaobaoTbkOrderDetailsGetPublisherOrderDto) SetTalentPid(v string) *TaobaoTbkOrderDetailsGetPublisherOrderDto {
    s.TalentPid = &v
    return s
}
func (s *TaobaoTbkOrderDetailsGetPublisherOrderDto) SetTbGmvTotalPrice(v string) *TaobaoTbkOrderDetailsGetPublisherOrderDto {
    s.TbGmvTotalPrice = &v
    return s
}
func (s *TaobaoTbkOrderDetailsGetPublisherOrderDto) SetExtraMktId(v string) *TaobaoTbkOrderDetailsGetPublisherOrderDto {
    s.ExtraMktId = &v
    return s
}
func (s *TaobaoTbkOrderDetailsGetPublisherOrderDto) SetUntts(v string) *TaobaoTbkOrderDetailsGetPublisherOrderDto {
    s.Untts = &v
    return s
}
func (s *TaobaoTbkOrderDetailsGetPublisherOrderDto) SetPubSharePreFeeForCommission(v string) *TaobaoTbkOrderDetailsGetPublisherOrderDto {
    s.PubSharePreFeeForCommission = &v
    return s
}
func (s *TaobaoTbkOrderDetailsGetPublisherOrderDto) SetPubShareFeeForCommission(v string) *TaobaoTbkOrderDetailsGetPublisherOrderDto {
    s.PubShareFeeForCommission = &v
    return s
}
func (s *TaobaoTbkOrderDetailsGetPublisherOrderDto) SetSubsidyInfoDtoList(v []TaobaoTbkOrderDetailsGetSubsidyDetailDTO) *TaobaoTbkOrderDetailsGetPublisherOrderDto {
    s.SubsidyInfoDtoList = &v
    return s
}
func (s *TaobaoTbkOrderDetailsGetPublisherOrderDto) SetPubShareRateForSdy(v string) *TaobaoTbkOrderDetailsGetPublisherOrderDto {
    s.PubShareRateForSdy = &v
    return s
}
func (s *TaobaoTbkOrderDetailsGetPublisherOrderDto) SetTkTotalRateForSdy(v string) *TaobaoTbkOrderDetailsGetPublisherOrderDto {
    s.TkTotalRateForSdy = &v
    return s
}
func (s *TaobaoTbkOrderDetailsGetPublisherOrderDto) SetPubSharePreFeeForSdy(v string) *TaobaoTbkOrderDetailsGetPublisherOrderDto {
    s.PubSharePreFeeForSdy = &v
    return s
}
func (s *TaobaoTbkOrderDetailsGetPublisherOrderDto) SetPubShareFeeForSdy(v string) *TaobaoTbkOrderDetailsGetPublisherOrderDto {
    s.PubShareFeeForSdy = &v
    return s
}
func (s *TaobaoTbkOrderDetailsGetPublisherOrderDto) SetAlimmShareInfoDto(v TaobaoTbkOrderDetailsGetAlimmShareInfoDTO) *TaobaoTbkOrderDetailsGetPublisherOrderDto {
    s.AlimmShareInfoDto = &v
    return s
}
func (s *TaobaoTbkOrderDetailsGetPublisherOrderDto) SetPlatformSpecialServiceRate(v string) *TaobaoTbkOrderDetailsGetPublisherOrderDto {
    s.PlatformSpecialServiceRate = &v
    return s
}
func (s *TaobaoTbkOrderDetailsGetPublisherOrderDto) SetPlatformSpecialServiceFee(v string) *TaobaoTbkOrderDetailsGetPublisherOrderDto {
    s.PlatformSpecialServiceFee = &v
    return s
}
func (s *TaobaoTbkOrderDetailsGetPublisherOrderDto) SetPlatformSpecialShareInfoDto(v TaobaoTbkOrderDetailsGetPlatformSpecialShareInfoDTO) *TaobaoTbkOrderDetailsGetPublisherOrderDto {
    s.PlatformSpecialShareInfoDto = &v
    return s
}

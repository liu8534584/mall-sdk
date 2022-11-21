package pddResponse

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

// Response pdd.ddk.order.list.increment.get（最后更新时间段增量同步推广订单信息）
type PddDdkOrderListIncrementGetRsp struct {
	ErrorResponse        PddErrorResponse     `json:"error_response"`
	OrderListGetResponse OrderListGetResponse `json:"order_list_get_response"`
}

type OrderListGetResponse struct {
	TotalCount int64       `json:"total_count"`
	OrderList  []OrderList `json:"order_list"`
	RequestID  string      `json:"request_id"`
}

type OrderList struct {
	GoodsPrice                 int64   `json:"goods_price"`
	OrderGroupSuccessTime      int64   `json:"order_group_success_time"`
	OrderModifyAt              int64   `json:"order_modify_at"`
	AuthDuoId                  int64   `json:"auth_duo_id"`
	CPANew                     int64   `json:"cpa_new"`
	GroupId                    float64 `json:"group_id"`
	OrderStatusDesc            string  `json:"order_status_desc"`
	ZsDuoId                    int64   `json:"zs_duo_id"`
	OrderId                    string  `json:"order_id" `
	OrderSn                    string  `json:"order_sn" `
	BatchNo                    string  `json:"batch_no" `
	GoodsId                    int64   `json:"goods_id" `
	GoodsName                  string  `json:"goods_name" `
	GoodsSign                  string  `json:"goods_sign" `
	GoodsQuantity              int64   `json:"goods_quantity" `
	GoodsThumbnailURL          string  `json:"goods_thumbnail_url" `
	IsDirect                   uint8   `json:"is_direct" `
	MallId                     int64   `json:"mall_id" `
	FailReason                 string  `json:"fail_reason" `
	NoSubsidyReason            string  `json:"no_subsidy_reason" `
	OrderAmount                int64   `json:"order_amount" `
	OrderCreateTime            int64   `json:"order_create_time" `
	OrderPayTime               int64   `json:"order_pay_time" `
	OrderReceiveTime           int64   `json:"order_receive_time" `
	OrderSettleTime            int64   `json:"order_settle_time" `
	OrderVerifyTime            int64   `json:"order_verify_time" `
	OrderStatus                int8    `json:"order_status" `
	PriceCompareStatus         uint8   `json:"price_compare_status" `
	PromotionAmount            int64   `json:"promotion_amount" `
	PromotionRate              uint32  `json:"promotion_rate" `
	PId                        string  `json:"p_id" `
	SubsidyAmount              int64   `json:"subsidy_amount" `
	SubsidyDuoAmountLevel      int64   `json:"subsidy_duo_amount_level" `
	SubsidyDuoAmountTenMillion int64   `json:"subsidy_duo_amount_ten_million" `
	SubsidyType                uint8   `json:"subsidy_type" `
	CustomParameters           string  `json:"custom_parameters" `
	Type                       uint8   `json:"type" `
	CatIds                     []int64 `json:"cat_ids" `
}

func (p OrderList) MarshalJSON() ([]byte, error) {
	type Order OrderList
	return json.Marshal(&struct {
		Order
		GoodsId string `json:"goods_id"`
		MallId  string `json:"mall_id"`
		CatIds  string `json:"cat_ids"`
	}{
		Order:   (Order)(p),
		GoodsId: strconv.FormatInt(p.GoodsId, 10),
		MallId:  strconv.FormatInt(p.MallId, 10),
		CatIds:  strings.Replace(fmt.Sprint(p.CatIds), " ", ",", -1),
	})
}

package request

// 订单详情
type OrderDetail struct {
	ItemId      int `json:"item_id"`
	AttrValueId int `json:"attr_value_id"`
	PayNum      int `json:"pay_num"`
}

// 提交订单
type SubmitOrderReq struct {
	Realname     string        `json:"realname"`
	UserPhone    string        `json:"user_phone"`
	UserAddress  string        `json:"user_address"`
	OrderDetails []OrderDetail `json:"order_details"`
}

// 申请退款
type ApplyRefundReq struct {
	OrderId       int         `json:"order_id"`
	RefundType    uint8       `json:"refund_type"`    // 退款类型(1:仅退款,2:退货退款)
	RefundImg     interface{} `json:"refund_img"`     // 退款图片
	RefundExplain string      `json:"refund_explain"` // 退款申请说明：衣服开线了...
	RefundReason  string      `json:"refund_reason"`  // 退款原因：不想要了
	RefundPrice   float64     `json:"refund_price"`   // 申请退款金额
}

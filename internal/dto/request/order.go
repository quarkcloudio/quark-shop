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
	OrderId             int    `json:"order_id"`
	RefundReasonImg     string `json:"refund_reason_img"`     // 退款图片
	RefundReasonExplain string `json:"refund_reason_explain"` // 退款申请说明：衣服开线了...
	RefundReason        string `json:"refund_reason"`         // 退款原因：不想要了
}

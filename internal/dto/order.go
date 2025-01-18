package dto

import (
	"github.com/quarkcloudio/quark-go/v3/utils/datetime"
)

// 订单信息
type OrderDTO struct {
	Id               int               `json:"id"`                    // 订单ID
	OrderNo          string            `json:"order_no"`              // 订单号
	Uid              int               `json:"uid"`                   // 用户id
	Realname         string            `json:"realname"`              // 用户姓名
	UserPhone        string            `json:"user_phone"`            // 用户电话
	UserAddress      string            `json:"user_address"`          // 详细地址
	TotalNum         int               `json:"total_num"`             // 订单商品总数
	TotalPrice       float64           `json:"total_price"`           // 订单总价
	PayPrice         float64           `json:"pay_price"`             // 实际支付金额
	Paid             bool              `json:"paid"`                  // 支付状态
	PayTime          datetime.Datetime `json:"pay_time"`              // 支付时间
	PayType          string            `json:"pay_type"`              // 支付方式
	PayTypeText      string            `json:"pay_type_text"`         // 支付方式文本
	OrderDetails     []OrderDetailDTO  `json:"orderDetails"`          // 订单详细信息
	Status           int               `json:"status"`                // 订单状态
	StatusText       string            `json:"status_text"`           // 订单状态文本
	RefundStatus     uint8             `json:"refund_status"`         // 退款状态
	RefundStatusText string            `json:"refund_status_text"`    // 退款状态文本
	RefundImg        string            `json:"refund_reason_img"`     // 退款图片
	RefundExplain    string            `json:"refund_reason_explain"` // 退款用户说明
	RefuseReason     string            `json:"refuse_reason"`         // 不退款的理由
	RefundTime       datetime.Datetime `json:"refund_time" `          // 申请退款时间
	RefundPrice      float64           `json:"refund_price"`          // 申请退款金额
	RefundExpress    string            `json:"refund_express"`        // 退货快递公司
	RefundExpressNo  string            `json:"refund_express_no"`     // 退货退款上门取件快递单号
	RefundNum        int               `json:"refund_num"`            // 退货退款上门取件退货数量
	RefundPhone      string            `json:"refund_phone"`          // 退货退款上门取件联系电话
	RefundType       uint8             `json:"refund_type"`           // 退款类型(1:仅退款,2:退货退款)
	RefundedPrice    float64           `json:"refunded_price"`        // 已退款金额
	RefundedTime     datetime.Datetime `json:"refunded_time"`         // 同意退款时间
	Remark           string            `json:"remark"`                // 管理员备注
	MerchantId       int               `json:"merchant_id"`           // 预留字段:商户ID
	IsMerchantCheck  bool              `json:"is_merchant_check"`     // 是否已核销
	Cost             float64           `json:"cost"`                  // 成本价
	VerifyCode       string            `json:"verify_code"`           // 核销码
	ShippingType     uint8             `json:"shipping_type"`         // 配送方式
	ClerkId          int               `json:"clerk_id"`              // 店员id/核销员id
	IsCancel         bool              `json:"is_cancel"`             // 是否取消订单
	IsAllRefund      bool              `json:"is_all_refund"`         // 是否全部退款
	IsSystemDel      bool              `json:"is_system_del"`         // 是否为后台删除
	CreatedAt        datetime.Datetime `json:"created_at"`            // 下单时间
	UpdatedAt        datetime.Datetime `json:"updated_at"`            // 记录更新时间
}

// 订单详情信息
type OrderDetailDTO struct {
	Id            int          `json:"id"`            // 主键
	OrderId       int          `json:"order_id"`      // 订单id
	ItemId        int          `json:"item_id"`       // 商品ID
	ItemInfo      ItemDTO      `json:"itemInfo"`      // 商品信息
	OrderNo       string       `json:"order_no"`      // 订单号
	Name          string       `json:"name"`          // 商品名称
	AttrValueId   int          `json:"attr_value_id"` // 规格属性值id
	AttrValueInfo AttrValueDTO `json:"attrInfo"`      // 商品规格信息
	Image         string       `json:"image"`         // 商品图片
	SKU           string       `json:"sku"`           // 商品sku
	Price         float64      `json:"price"`         // 商品价格
	PayNum        int          `json:"pay_num"`       // 购买数量
}

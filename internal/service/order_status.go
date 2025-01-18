package service

import (
	"strconv"
	"strings"

	"github.com/quarkcloudio/quark-go/v3/dal/db"
	"github.com/quarkcloudio/quark-smart/v2/internal/model"
)

type ChangeTypeItem = struct {
	Type    string
	Message string
}

type OrderStatusService struct {
	OrderId           int
	CreateOrder       ChangeTypeItem
	PaySuccess        ChangeTypeItem
	DeliveryGoods     ChangeTypeItem
	TakeDelivery      ChangeTypeItem
	CheckOrderOver    ChangeTypeItem
	ApplyRefund       ChangeTypeItem
	RefundPrice       ChangeTypeItem
	CancelRefundOrder ChangeTypeItem
}

// 操作类型
// create_order:订单生成;
// pay_success:用户付款成功;
// delivery_goods:已发货 快递公司：圆通速递 快递单号：YT46466545445555;
// take_delivery:已收货;
// check_order_over:用户评价;
// apply_refund:用户申请退款，原因：收货地址填错了;
// refund_price:退款给用户：124.63元;
// cancel_refund_order:不退款原因:不符合退款要求
func NewOrderStatusService(orderId int) *OrderStatusService {
	service := &OrderStatusService{
		orderId,
		ChangeTypeItem{"create_order", "订单生成"},
		ChangeTypeItem{"pay_success", "用户付款成功"},
		ChangeTypeItem{"delivery_goods", "已发货 快递公司：{company} 快递单号：{number}"},
		ChangeTypeItem{"take_delivery", "已收货"},
		ChangeTypeItem{"check_order_over", "用户评价"},
		ChangeTypeItem{"apply_refund", "用户申请退款，原因：{reason}"},
		ChangeTypeItem{"refund_price", "退款给用户：{price}元"},
		ChangeTypeItem{"cancel_refund_order", "不退款原因：{reason}"},
	}
	return service
}

// 根据订单id获取订单记录列表
func (p *OrderStatusService) GetList() (statuses []model.OrderStatus, err error) {
	err = db.Client.Where("order_id = ?", p.OrderId).Find(&statuses).Error
	return
}

// 创建订单记录
func (p *OrderStatusService) Store(changeType string, changeMessage string) (err error) {
	err = db.Client.Create(&model.OrderStatus{
		OrderId:       p.OrderId,
		ChangeType:    changeType,
		ChangeMessage: changeMessage,
	}).Error
	return
}

// 变更为订单生成状态
func (p *OrderStatusService) ChangeToCreateOrderStatus() (err error) {
	return p.Store(p.CreateOrder.Type, p.CreateOrder.Message)
}

// 变更为用户付款成功状态
func (p *OrderStatusService) ChangeToPaySuccessStatus() (err error) {
	return p.Store(p.PaySuccess.Type, p.PaySuccess.Message)
}

// 变更为已发货状态
func (p *OrderStatusService) ChangeToDeliveryGoodsStatus(company string, number string) (err error) {
	message := strings.ReplaceAll(p.DeliveryGoods.Message, "{company}", company)
	message = strings.ReplaceAll(message, "{number}", number)
	return p.Store(p.DeliveryGoods.Type, message)
}

// 变更为已收货状态
func (p *OrderStatusService) ChangeToTakeDeliveryStatus() (err error) {
	return p.Store(p.TakeDelivery.Type, p.TakeDelivery.Message)
}

// 变更为用户已评价，订单已完成状态
func (p *OrderStatusService) ChangeToCheckOrderOverStatus() (err error) {
	return p.Store(p.CheckOrderOver.Type, p.CheckOrderOver.Message)
}

// 变更为申请退款状态
func (p *OrderStatusService) ChangeToApplyRefundStatus(reason string) (err error) {
	message := strings.ReplaceAll(p.ApplyRefund.Message, "{reason}", reason)
	return p.Store(p.ApplyRefund.Type, message)
}

// 变更为给用户退款状态
func (p *OrderStatusService) ChangeToRefundPriceStatus(price float64) (err error) {
	message := strings.ReplaceAll(p.RefundPrice.Message, "{price}", strconv.FormatFloat(price, 'f', 2, 64))
	return p.Store(p.RefundPrice.Type, message)
}

// 变更为不退款状态
func (p *OrderStatusService) ChangeToCancelRefundOrderStatus(reason string) (err error) {
	message := strings.ReplaceAll(p.CancelRefundOrder.Message, "{reason}", reason)
	return p.Store(p.CancelRefundOrder.Type, message)
}

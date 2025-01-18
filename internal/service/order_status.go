package service

import (
	"github.com/quarkcloudio/quark-smart/v2/internal/model"
)

type OrderStatusService struct{}

func NewOrderStatusService() *OrderStatusService {
	return &OrderStatusService{}
}

// 获取订单状态文本
func (p *OrderStatusService) GetStatusText(order model.Order) string {
	statusText := ""
	if order.Paid == 0 {
		return "待支付"
	}

	// -1:申请退款;-2:退款成功;0:待发货;1:待收货;2:已收货,待评价;3:已完成
	if order.ShippingType == 1 {
		switch order.Status {
		case -2:
			statusText = "退款成功"
		case -1:
			statusText = "申请退款"
		case 0:
			statusText = "待发货"
		case 1:
			statusText = "待收货"
		case 2:
			statusText = "已收货,待评价"
		case 3:
			statusText = "已完成"
		}
	}

	// -1:申请退款;-2:退款成功;0:待核销;1:待核销;2:已核销;3:已完成
	if order.ShippingType == 2 {
		switch order.Status {
		case -2:
			statusText = "退款成功"
		case -1:
			statusText = "申请退款"
		case 0:
			statusText = "待核销"
		case 1:
			statusText = "待核销"
		case 2:
			statusText = "已核销"
		case 3:
			statusText = "已完成"
		}
	}

	return statusText
}

// 获取退款状态文本
func (p *OrderStatusService) GetRefundStatusText(order model.Order) string {
	statusText := ""
	if order.Paid == 0 {
		return ""
	}

	// 0:未退款,1:申请中,2:已退款
	switch order.RefundStatus {
	case 0:
		statusText = "未退款"
	case 1:
		statusText = "申请中"
	case 2:
		statusText = "已退款"
	}

	return statusText
}

// 获取支付方式文本
func (p *OrderStatusService) GetPayTypeText(payType string) string {
	text := ""

	// WECHAT_PAY,ALI_PAY,OFFLINE_PAY,YUE_PAY
	switch payType {
	case "WECHAT_PAY":
		text = "微信支付"
	case "ALI_PAY":
		text = "支付宝支付"
	case "OFFLINE_PAY":
		text = "线下支付"
	case "YUE_PAY":
		text = "余额支付"
	}

	return text
}

package action

import (
	"strconv"

	"github.com/quarkcloudio/quark-go/v3"
	"github.com/quarkcloudio/quark-go/v3/template/admin/component/message"
	"github.com/quarkcloudio/quark-go/v3/template/admin/resource"
	"github.com/quarkcloudio/quark-go/v3/template/admin/resource/actions"
	"github.com/quarkcloudio/quark-smart/v2/internal/service"
	"gorm.io/gorm"
)

type OrderAgreeRefundAction struct {
	actions.ModalForm
}

// 订单退款
func OrderAgreeRefund() *OrderAgreeRefundAction {
	return &OrderAgreeRefundAction{}
}

// 初始化
func (p *OrderAgreeRefundAction) Init(ctx *quark.Context) interface{} {

	// 设置按钮文字
	p.Name = "<%= ((paid===1 && refund_status<=1) && '立即退款') %>"

	// 类型
	p.Type = "link"

	// 设置按钮大小,large | middle | small | default
	p.Size = "small"

	// 关闭时销毁 Modal 里的元素
	p.DestroyOnClose = true

	//  执行成功后刷新的组件
	p.Reload = "table"

	// 在表格行内展示
	p.SetOnlyOnIndexTableRow(true)

	// 行为接口接收的参数
	p.SetApiParams([]string{
		"id",
		"pay_price",
	})

	return p
}

// 字段
func (p *OrderAgreeRefundAction) Fields(ctx *quark.Context) []interface{} {
	field := &resource.Field{}
	return []interface{}{
		field.Hidden("id", "ID"),
		field.Display("订单号", "${order_no}"),
		field.Display("退款信息", "付款金额(${pay_price})  已退金额(${has_refund_price})  剩余可退(${can_refund_price})"),
		field.Number("refund_price", "退款金额").
			SetRequired().
			SetPlaceholder("请输入退款金额"),
	}
}

// 表单数据（异步获取）
func (p *OrderAgreeRefundAction) Data(ctx *quark.Context) map[string]interface{} {
	id, _ := strconv.Atoi(ctx.Query("id").(string))
	order, _ := service.NewOrderService().GetOrderById(id)
	return map[string]interface{}{
		"id":               id,
		"pay_price":        order.PayPrice,
		"has_refund_price": order.RefundPrice,
		"can_refund_price": order.PayPrice - order.RefundPrice,
	}
}

// 执行行为句柄
func (p *OrderAgreeRefundAction) Handle(ctx *quark.Context, query *gorm.DB) error {
	var refundReq struct {
		Id          int     `json:"id"`
		RefundPrice float64 `json:"refund_price"`
	}
	if err := ctx.Bind(&refundReq); err != nil {
		return ctx.JSON(200, message.Error(err.Error()))
	}
	if err := service.NewOrderService().AgreeRefund(refundReq.Id, refundReq.RefundPrice); err != nil {
		return ctx.JSON(200, message.Error(err.Error()))
	}
	return ctx.JSON(200, message.Success("操作成功"))
}

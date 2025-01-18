package action

import (
	"github.com/quarkcloudio/quark-go/v3"
	"github.com/quarkcloudio/quark-go/v3/template/admin/component/message"
	"github.com/quarkcloudio/quark-go/v3/template/admin/resource"
	"github.com/quarkcloudio/quark-go/v3/template/admin/resource/actions"
	"github.com/quarkcloudio/quark-smart/v2/internal/service"
	"gorm.io/gorm"
)

type OrderRefuseRefundAction struct {
	actions.ModalForm
}

// 订单退款
func OrderRefuseRefund() *OrderRefuseRefundAction {
	return &OrderRefuseRefundAction{}
}

// 初始化
func (p *OrderRefuseRefundAction) Init(ctx *quark.Context) interface{} {

	// 设置按钮文字
	p.Name = "<%= ((paid===1 && refund_status===1) && '不退款') %>"

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
func (p *OrderRefuseRefundAction) Fields(ctx *quark.Context) []interface{} {
	field := &resource.Field{}
	return []interface{}{
		field.Hidden("id", "ID"),
		field.Display("订单号", "${order_no}"),
		field.TextArea("refuse_reason", "不退款原因").
			SetRequired().
			SetPlaceholder("请输入不退款原因"),
	}
}

// 执行行为句柄
func (p *OrderRefuseRefundAction) Handle(ctx *quark.Context, query *gorm.DB) error {
	var refundReq struct {
		Id           int    `json:"id"`
		RefuseReason string `json:"refuse_reason"`
	}
	if err := ctx.Bind(&refundReq); err != nil {
		return ctx.JSON(200, message.Error(err.Error()))
	}
	if err := service.NewOrderService().RefuseRefund(refundReq.Id, refundReq.RefuseReason); err != nil {
		return ctx.JSON(200, message.Error(err.Error()))
	}
	return ctx.JSON(200, message.Success("操作成功"))
}

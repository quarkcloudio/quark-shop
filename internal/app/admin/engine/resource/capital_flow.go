package resource

import (
	"github.com/quarkcloudio/quark-go/v3"
	"github.com/quarkcloudio/quark-go/v3/app/admin/searches"
	"github.com/quarkcloudio/quark-go/v3/template/admin/component/form/fields/selectfield"
	"github.com/quarkcloudio/quark-go/v3/template/admin/resource"
	"github.com/quarkcloudio/quark-smart/v2/internal/app/admin/engine/action"
	"github.com/quarkcloudio/quark-smart/v2/internal/model"
	"gorm.io/gorm"
)

type CapitalFlow struct {
	resource.Template
}

// 初始化
func (p *CapitalFlow) Init(ctx *quark.Context) interface{} {

	// 标题
	p.Title = "资金流水"

	// 模型
	p.Model = &model.CapitalFlow{}

	// 分页
	p.PageSize = 10

	return p
}

// 查询
func (p *CapitalFlow) Query(ctx *quark.Context, query *gorm.DB) *gorm.DB {

	return query.Select(
		"capital_flows.id",
		"capital_flows.flow_no",
		"capital_flows.nickname",
		"capital_flows.phone",
		"capital_flows.price",
		"capital_flows.trading_type",
		"capital_flows.pay_type",
		"capital_flows.mark",
		"capital_flows.created_at",
		"orders.order_no",
	).
		Joins("JOIN orders ON orders.id = capital_flows.order_id")
}

func (p *CapitalFlow) Fields(ctx *quark.Context) []interface{} {
	field := &resource.Field{}

	return []interface{}{
		field.Hidden("id", "ID"),

		field.Text("flow_no", "交易单号").
			SetColumnWidth(180).
			SetEllipsis(true),

		field.Text("order_no", "关联订单").
			SetColumnWidth(180).
			SetEllipsis(true),

		field.Datetime("created_at", "交易时间").
			SetColumnWidth(160),

		field.Text("price", "交易金额").
			SetColumnWidth(100),

		field.Text("nickname", "交易用户").
			SetColumnWidth(100),

		field.Select("pay_type", "支付方式").
			SetOptions([]selectfield.Option{
				{
					Label: "微信支付",
					Value: "WECHAT_PAY",
				},
				{
					Label: "支付宝支付",
					Value: "ALI_PAY",
				},
				{
					Label: "线下支付",
					Value: "OFFLINE_PAY",
				},
				{
					Label: "余额支付",
					Value: "YUE_PAY",
				},
			}).
			SetColumnWidth(100),

		field.TextArea("mark", "备注").
			SetEllipsis(true),
	}
}

// 搜索
func (p *CapitalFlow) Searches(ctx *quark.Context) []interface{} {
	return []interface{}{
		searches.DatetimeRange("capital_flows.created_at", "订单时间"),
		searches.Input("orders.order_no", "关联订单"),
		searches.Input("capital_flows.nickname", "交易用户"),
	}
}

// 行为
func (p *CapitalFlow) Actions(ctx *quark.Context) []interface{} {
	return []interface{}{
		action.BillMark("备注"),
	}
}

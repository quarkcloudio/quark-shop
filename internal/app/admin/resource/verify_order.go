package resource

import (
	"github.com/quarkcloudio/quark-go/v3"
	"github.com/quarkcloudio/quark-go/v3/app/admin/searches"
	"github.com/quarkcloudio/quark-go/v3/template/admin/resource"
	"github.com/quarkcloudio/quark-smart/v2/internal/model"
	"gorm.io/gorm"
)

type VerifyOrder struct {
	resource.Template
}

// 初始化
func (p *VerifyOrder) Init(ctx *quark.Context) interface{} {

	// 标题
	p.Title = "核销记录"

	// 模型
	p.Model = &model.Order{}

	// 分页
	p.PageSize = 10

	return p
}

// 查询类型
func (p *VerifyOrder) Query(ctx *quark.Context, query *gorm.DB) *gorm.DB {
	return query.Where("shipping_type = ?", 2).Where("status = ?", 2)
}

func (p *VerifyOrder) Fields(ctx *quark.Context) []interface{} {
	field := &resource.Field{}

	return []interface{}{
		field.Hidden("id", "ID"),

		field.Text("title", "订单号"),

		field.Text("name", "商品信息"),

		field.Text("user_info", "用户信息"),

		field.Text("description", "支付金额"),

		field.Text("type", "核销人员"),

		field.Text("status", "订单状态"),

		field.Text("created_at", "下单时间"),

		field.Text("created_at", "核销时间"),
	}
}

// 搜索
func (p *VerifyOrder) Searches(ctx *quark.Context) []interface{} {
	return []interface{}{
		searches.Input("order_no", "订单号"),
		searches.DatetimeRange("pay_time", "支付时间"),
	}
}

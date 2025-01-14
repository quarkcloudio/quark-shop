package action

import (
	"strconv"

	"github.com/quarkcloudio/quark-go/v3"
	"github.com/quarkcloudio/quark-go/v3/template/admin/component/order"
	"github.com/quarkcloudio/quark-go/v3/template/admin/resource/actions"
	"github.com/quarkcloudio/quark-smart/v2/internal/service"
)

type OrderDetailAction struct {
	actions.Drawer
}

// 创建-抽屉类型
func OrderDetail() *OrderDetailAction {
	return &OrderDetailAction{}
}

// 初始化
func (p *OrderDetailAction) Init(ctx *quark.Context) interface{} {

	// 文字
	p.Name = "订单详情"

	// 类型
	p.Type = "link"

	// 设置按钮大小,large | middle | small | default
	p.Size = "small"

	// 关闭时销毁 Drawer 里的子元素
	p.DestroyOnClose = true

	// 执行成功后刷新的组件
	p.Reload = "table"

	// 宽度
	p.Width = 1000

	// 设置展示位置
	p.SetOnlyOnIndexTableRow(true)

	return p
}

// 内容
func (p *OrderDetailAction) GetBody(ctx *quark.Context) interface{} {
	// 基本信息
	info := order.
		NewInfo().
		SetDataIndex("baseInfo").
		SetColon(false).
		SetLayout("vertical").
		SetItems([]order.Item{
			order.NewItem("status_text", "订单状态"),
			order.NewItem("pay_price", "实际支付"),
			order.NewItem("refund_price", "实际退款"),
			order.NewItem("pay_type_text", "支付方式"),
			order.NewItem("pay_time", "支付时间"),
		})

	// 用户信息
	userInfo := order.
		NewInfo().
		SetDataIndex("userInfo").
		SetTitle("用户信息").
		SetColumn(2).
		SetItems([]order.Item{
			order.NewItem("username", "用户名"),
			order.NewItem("phone", "联系电话"),
		})

	// 订单信息
	orderInfo := order.
		NewInfo().
		SetDataIndex("orderInfo").
		SetTitle("订单信息").
		SetColumn(4).
		SetItems([]order.Item{
			order.NewItem("created_at", "下单时间"),
			order.NewItem("total_num", "商品总数"),
			order.NewItem("total_price", "商品总价"),
			order.NewItem("pay_price", "实际支付"),
		})

	detailInfo := []*order.Info{
		userInfo,
		orderInfo,
	}

	// 商品信息
	itemInfo := order.
		NewTable().
		SetDataIndex("itemInfo").
		SetColumn([]order.Column{
			order.NewColumn("image", "商品图片", "image"),
			order.NewColumn("name", "商品名称"),
			order.NewColumn("sku", "商品规格"),
			order.NewColumn("price", "支付价格"),
			order.NewColumn("pay_num", "购买数量"),
		})

	// 订单记录
	statusInfo := order.
		NewTable().
		SetDataIndex("statusInfo").
		SetColumn([]order.Column{
			order.NewColumn("order_id", "订单ID"),
			order.NewColumn("change_message", "操作记录"),
			order.NewColumn("created_at", "操作时间"),
		})
	return order.
		New().
		SetInitApi("/api/admin/order/action/order-detail-action/values?id=${id}").
		SetInfo(info).
		SetDetailInfo(detailInfo).
		SetItemInfo(itemInfo).
		SetStatusInfo(statusInfo)
}

// 表单数据（异步获取）
func (p *OrderDetailAction) Data(ctx *quark.Context) map[string]interface{} {
	id, _ := strconv.Atoi(ctx.Query("id").(string))
	order, _ := service.NewOrderService().GetOrderById(id)
	user, _ := service.NewUserService().GetInfoById(order.Uid)
	statuses, _ := service.NewOrderService().GetStatuses(id)
	return map[string]interface{}{
		"baseInfo":   order,
		"userInfo":   user,
		"orderInfo":  order,
		"itemInfo":   order.OrderDetails,
		"statusInfo": statuses,
	}
}

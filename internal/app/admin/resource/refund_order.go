package resource

import (
	"fmt"

	"github.com/quarkcloudio/quark-go/v3"
	"github.com/quarkcloudio/quark-go/v3/app/admin/searches"
	"github.com/quarkcloudio/quark-go/v3/template/admin/resource"
	"github.com/quarkcloudio/quark-smart/v2/internal/app/admin/action"
	"github.com/quarkcloudio/quark-smart/v2/internal/model"
	"github.com/quarkcloudio/quark-smart/v2/internal/service"
	"github.com/quarkcloudio/quark-smart/v2/pkg/utils"
	"gorm.io/gorm"
)

type RefundOrder struct {
	resource.Template
}

// 初始化
func (p *RefundOrder) Init(ctx *quark.Context) interface{} {

	// 标题
	p.Title = "售后订单"

	// 模型
	p.Model = &model.Order{}

	// 分页
	p.PageSize = 10

	return p
}

// 查询类型
func (p *RefundOrder) Query(ctx *quark.Context, query *gorm.DB) *gorm.DB {
	activeKey := ctx.QueryParam("activeKey")
	query.Where("status < ?", 0)
	switch activeKey {
	case "all":
		// 全部
	case "1":
		// 申请中
		query.Where("refund_status = ?", 1)
	case "2":
		// 已退款
		query.Where("refund_status = ?", 2)
	}
	return query
}

// 菜单
func (p *RefundOrder) Menus(ctx *quark.Context) interface{} {
	return map[string]interface{}{
		"type": "tab",
		"items": []map[string]string{
			{
				"key":   "all",
				"label": "全部",
			},
			{
				"key":   "1",
				"label": "申请中",
			},
			{
				"key":   "2",
				"label": "已退款",
			},
		},
	}
}

func (p *RefundOrder) Fields(ctx *quark.Context) []interface{} {
	field := &resource.Field{}

	return []interface{}{
		field.Hidden("id", "ID"),

		field.Text("refund_no", "退款单号"),

		field.Text("order_no", "原订单号"),

		field.Text("name", "商品信息", func(row map[string]interface{}) interface{} {
			result := ""
			orderDetails, err := service.NewOrderService().GetOrderDetailsByOrderId(row["id"])
			if err != nil {
				return result
			}
			for k, orderDetail := range orderDetails {
				name := orderDetail.Name
				image := utils.GetImagePath(orderDetail.Image)
				style := ""
				if k != 0 {
					style = "margin-top:5px"
				}
				price := orderDetail.Price * float64(orderDetail.PayNum)
				title := fmt.Sprintf("商品名称：%s\r\n规格名称：%s\r\n支付价格：¥%.2f\r\n购买数量：%d", name, orderDetail.SKU, price, orderDetail.PayNum)
				result = result + fmt.Sprintf("<div title='%s' style='%s'><img src='%s' height=40 width=40 /> %s</div>", title, style, image, name)
			}
			return result
		}).SetColumnWidth(250),

		field.Text("user_info", "用户信息", func(row map[string]interface{}) interface{} {
			userInfo, err := service.NewUserService().GetInfoById(row["uid"])
			if err != nil {
				return nil
			}
			return fmt.Sprintf("用户ID：%d</br>用户账号：%s</br>用户昵称：%s", userInfo.Id, userInfo.Username, userInfo.Nickname)
		}),

		field.Text("total_pay", "支付金额"),

		field.Text("refund_reason_time", "发起退款时间"),

		// 0:未退款,1:申请中,2:已退款
		field.Text("refund_status", "退款状态", func(row map[string]interface{}) interface{} {
			result := ""
			switch row["refund_status"] {
			case 0:
				result = "未退款"
			case 1:
				result = "申请中"
			case 2:
				result = "已退款"
			}
			return result
		}),

		field.Text("status", "订单状态", func(row map[string]interface{}) interface{} {
			if row["paid"].(uint8) == 0 {
				return "未付款"
			}
			result := ""
			switch row["status"] {
			case -2:
				result = "退款成功"
			case -1:
				result = "申请退款"
			case 0:
				result = "待发货"
			case 1:
				result = "待收货"
			case 2:
				result = "已收货,待评价"
			case 3:
				result = "已完成"
			}
			return result
		}),

		field.Text("refund_reason", "退款信息"),
	}
}

// 搜索
func (p *RefundOrder) Searches(ctx *quark.Context) []interface{} {
	return []interface{}{
		searches.Input("refund_no", "退款单号"),
		searches.Input("order_no", "订单号"),
		searches.DatetimeRange("refund_reason_time", "发起时间"),
	}
}

// 行为
func (p *RefundOrder) Actions(ctx *quark.Context) []interface{} {
	return []interface{}{
		action.OrderAgreeRefund(),
		action.OrderRefuseRefund(),
		action.OrderDetail(),
	}
}

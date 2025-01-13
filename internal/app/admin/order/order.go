package order

import "github.com/quarkcloudio/quark-go/v3/template/admin/component/component"

type Component struct {
	component.Element
	InitApi         string     `json:"initApi"`
	Icon            string     `json:"icon"`
	OrderNoText     string     `json:"orderNoText"`
	OrderDetailText string     `json:"orderDetailText"`
	OrderItemText   string     `json:"orderItemText"`
	OrderStatusText string     `json:"orderStatusText"`
	OrderNo         string     `json:"orderNo"`
	Info            Info       `json:"info"`
	DetailInfo      DetailInfo `json:"detailInfo"`
	ItemInfo        ItemInfo   `json:"itemInfo"`
	StatusInfo      StatusInfo `json:"statusInfo"`
}

// 初始化组件
func New() *Component {
	return (&Component{}).Init()
}

// 初始化
func (p *Component) Init() *Component {
	p.Component = "order"
	p.SetKey(component.DEFAULT_KEY, component.DEFAULT_CRYPT)
	return p
}

// Set style.
func (p *Component) SetStyle(style map[string]interface{}) *Component {
	p.Style = style
	return p
}

// 初始化数据接口
func (p *Component) SetInitApi(initApi string) *Component {
	p.InitApi = initApi
	return p
}

// 设置图标地址
func (p *Component) SetIcon(icon string) *Component {
	p.Icon = icon
	return p
}

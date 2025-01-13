package order

import "github.com/quarkcloudio/quark-go/v3/template/admin/component/component"

type Component struct {
	component.Element
	InitApi         string `json:"initApi"`
	Icon            string `json:"icon"`
	OrderNoText     string `json:"orderNoText"`
	OrderDetailText string `json:"orderDetailText"`
	OrderItemText   string `json:"orderItemText"`
	OrderStatusText string `json:"orderStatusText"`
	OrderNo         string `json:"orderNo"`
	Info            Info   `json:"info"`
	DetailInfo      Detail `json:"detailInfo"`
	ItemInfo        Table  `json:"itemInfo"`
	StatusInfo      Table  `json:"statusInfo"`
}

type Column struct {
	Title     string `json:"title"`
	DataIndex string `json:"dataIndex"`
}

type Item struct {
	Key   string `json:"key"`
	Label string `json:"label"`
}

// 初始化组件
func New() *Component {
	return (&Component{}).Init()
}

func NewColumn(dataIndex string, title string) Column {
	return Column{title, dataIndex}
}

func NewItem(key string, label string) Item {
	return Item{key, label}
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

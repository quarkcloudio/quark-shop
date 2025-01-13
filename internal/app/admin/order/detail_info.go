package order

import "github.com/quarkcloudio/quark-go/v3/template/admin/component/component"

type DetailInfo struct {
	component.Element
	InitApi         string `json:"initApi"`
	Icon            string `json:"icon"`
	OrderNoText     string `json:"orderNoText"`
	OrderDetailText string `json:"orderDetailText"`
	OrderItemText   string `json:"orderItemText"`
	OrderStatusText string `json:"orderStatusText"`
	OrderNo         string `json:"orderNo"`
	Info            string `json:"info"`
	DetailInfo      string `json:"detailInfo"`
	ItemInfo        string `json:"itemInfo"`
	StatusInfo      string `json:"statusInfo"`
}

// 初始化组件
func NewDetailInfo() *DetailInfo {
	return (&DetailInfo{}).Init()
}

// 初始化
func (p *DetailInfo) Init() *DetailInfo {
	p.Component = "detailInfo"
	p.SetKey(component.DEFAULT_KEY, component.DEFAULT_CRYPT)

	return p
}

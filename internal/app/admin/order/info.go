package order

import "github.com/quarkcloudio/quark-go/v3/template/admin/component/component"

type Info struct {
	component.Element
	Title  string `json:"title"`
	Column int    `json:"column"`
	Layout string `json:"layout"`
	Colon  bool   `json:"colon"`
	Items  []Item `json:"items"`
}

// 初始化组件
func NewInfo() *Info {
	return (&Info{}).Init()
}

// 初始化
func (p *Info) Init() *Info {
	p.Component = "info"
	p.Colon = false
	p.Column = 5
	p.Layout = "vertical"
	p.SetKey(component.DEFAULT_KEY, component.DEFAULT_CRYPT)
	return p
}

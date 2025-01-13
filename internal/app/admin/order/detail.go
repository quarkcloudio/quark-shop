package order

type Detail struct {
	InfoItems  []Info      `json:"infoItems"`
	DataSource interface{} `json:"dataSource"`
}

// 初始化组件
func NewDetail(infoItems []Info, dataSource interface{}) Detail {
	return Detail{infoItems, dataSource}
}

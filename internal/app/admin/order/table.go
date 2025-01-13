package order

type Table struct {
	Columns    []Column    `json:"columns"`
	DataSource interface{} `json:"dataSource"`
}

// 初始化组件
func NewTable(columns []Column, dataSource interface{}) Table {
	return Table{columns, dataSource}
}

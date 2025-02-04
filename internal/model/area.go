package model

// 地区模型
type Area struct {
	Code     int    `json:"code"`
	Name     string `json:"name"`
	Level    int    `json:"level"`
	Pcode    int    `json:"pcode"`
	Category int    `json:"category"`
}

package response

// 用户地址列表
type UserAddressIndexResp struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Phone     string `json:"phone"`
	Address   string `json:"address"`
	IsDefault int    `json:"is_default"`
}

// 用户地址详情
type UserAddressDetailResp struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Phone     string `json:"phone"`
	Province  int    `json:"province"`
	City      int    `json:"city"`
	District  int    `json:"district"`
	Address   string `json:"address"`
	IsDefault int    `json:"is_default"`
}

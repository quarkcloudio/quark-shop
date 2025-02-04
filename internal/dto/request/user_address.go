package request

// 用户地址列表
type UserAddressIndexReq struct {
	PageReq
	Uid int
}

// 用户地址id
type UserAddressIdReq struct {
	Id int `json:"id" query:"id"`
}

// 新增用户地址
type UserAddressAddReq struct {
	Uid       int
	Name      string `json:"name"`
	Phone     string `json:"phone"`
	Province  int    `json:"province"`
	City      int    `json:"city"`
	District  int    `json:"district"`
	Address   string `json:"address"`
	IsDefault string `json:"is_default"`
}

// 更新用户地址
type UserAddressSaveReq struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Phone     string `json:"phone"`
	Province  int    `json:"province"`
	City      int    `json:"city"`
	District  int    `json:"district"`
	Address   string `json:"address"`
	IsDefault string `json:"is_default"`
}

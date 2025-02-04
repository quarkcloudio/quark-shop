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

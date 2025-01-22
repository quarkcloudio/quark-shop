package response

// 分页返回
type PageResp struct {
	List  interface{} `json:"list"`
	Total int         `json:"total"`
}

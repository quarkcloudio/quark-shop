package request

type AreaOptionReq struct {
	Pcode int `query:"pcode" default:"0"`
}

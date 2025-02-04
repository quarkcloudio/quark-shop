package handler

import (
	"github.com/quarkcloudio/quark-go/v3"
	"github.com/quarkcloudio/quark-smart/v2/internal/dto/request"
	"github.com/quarkcloudio/quark-smart/v2/internal/dto/response"
	"github.com/quarkcloudio/quark-smart/v2/internal/service"
)

type Area struct{}

// 地区选项
func (p *Area) Options(ctx *quark.Context) error {
	var param request.AreaOptionReq
	if err := ctx.Bind(&param); err != nil {
		return ctx.JSONError(err.Error())
	}

	options := make([]response.Option, 0)

	areas := service.NewAreaService().GetListByPcode(param.Pcode, 3)
	for _, area := range areas {
		options = append(options, response.Option{
			Label: area.Name,
			Value: area.Code,
		})
	}

	return ctx.JSONOk("ok", options)
}

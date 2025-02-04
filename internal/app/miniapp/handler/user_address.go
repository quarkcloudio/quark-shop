package handler

import (
	"github.com/quarkcloudio/quark-go/v3"
	"github.com/quarkcloudio/quark-smart/v2/internal/dto/request"
	"github.com/quarkcloudio/quark-smart/v2/internal/dto/response"
	"github.com/quarkcloudio/quark-smart/v2/internal/service"
)

type UserAddress struct{}

// 用户地址列表
func (p *UserAddress) Index(ctx *quark.Context) error {

	var param request.UserAddressIndexReq
	if err := ctx.Bind(&param); err != nil {
		return ctx.JSONError(err.Error())
	}

	// 获取用户id
	uid, err := service.NewAuthService(ctx).GetUid()
	if err != nil {
		return ctx.JSONError(err.Error())
	}
	param.Uid = uid

	list := make([]response.UserAddressIndexResp, 0)

	userAddresses, total := service.NewUserAddressService().GetUserAddressesByUid(param)
	for _, userAddress := range userAddresses {
		address := service.NewAreaService().GetDetailByCode(userAddress.Province).Name
		address += service.NewAreaService().GetDetailByCode(userAddress.City).Name
		address += service.NewAreaService().GetDetailByCode(userAddress.District).Name
		list = append(list, response.UserAddressIndexResp{
			Id:        userAddress.Id,
			Name:      userAddress.Name,
			Phone:     userAddress.Phone,
			Address:   address + userAddress.Address,
			IsDefault: userAddress.IsDefault,
		})
	}

	return ctx.JSONOk("ok", response.PageResp{
		List:  list,
		Total: total,
	})
}

// 用户地址详情
func (p *UserAddress) Detail(ctx *quark.Context) error {

	var param request.UserAddressIdReq
	if err := ctx.Bind(&param); err != nil {
		return ctx.JSONError(err.Error())
	}

	userAddress := service.NewUserAddressService().GetUserAddressById(param.Id)

	return ctx.JSONOk("ok", userAddress)
}

// 新增用户地址
func (p *UserAddress) Create(ctx *quark.Context) error {
	return ctx.JSONOk("ok")
}

// 更新用户地址
func (p *UserAddress) Update(ctx *quark.Context) error {
	return ctx.JSONOk("ok")
}

// 删除用户地址
func (p *UserAddress) Delete(ctx *quark.Context) error {
	return ctx.JSONOk("ok")
}

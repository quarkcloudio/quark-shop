package handler

import (
	"github.com/quarkcloudio/quark-go/v3"
	"github.com/quarkcloudio/quark-smart/v2/internal/dto/request"
	"github.com/quarkcloudio/quark-smart/v2/internal/dto/response"
	"github.com/quarkcloudio/quark-smart/v2/internal/model"
	"github.com/quarkcloudio/quark-smart/v2/internal/service"
	"github.com/quarkcloudio/quark-smart/v2/pkg/utils"
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
func (p *UserAddress) Add(ctx *quark.Context) error {
	var param request.UserAddressAddReq
	if err := ctx.Bind(&param); err != nil {
		return ctx.JSONError(err.Error())
	}

	if param.Name == "" {
		return ctx.JSONError("请填写收货人姓名")
	}
	if param.Phone == "" || !utils.CheckRegex("^(13[0-9]|14[01456879]|15[0-35-9]|16[2567]|17[0-8]|18[0-9]|19[0-35-9])\\d{8}$", param.Phone) {
		return ctx.JSONError("联系电话格式不正确")
	}
	if param.Province == 0 || param.City == 0 || param.District == 0 {
		return ctx.JSONError("请选择完整地区信息")
	}
	if param.Address == "" {
		return ctx.JSONError("请填写详细地址")
	}

	uid, _ := service.NewAuthService(ctx).GetUid()

	if err := service.NewUserAddressService().CreateUserAddress(model.UserAddress{
		Uid:       uid,
		Name:      param.Name,
		Phone:     param.Phone,
		Province:  param.Province,
		City:      param.City,
		District:  param.District,
		Address:   param.Address,
		IsDefault: param.IsDefault,
	}); err != nil {
		return ctx.JSONError("新增用户地址失败")
	}

	return ctx.JSONOk("操作成功")
}

// 更新用户地址
func (p *UserAddress) Save(ctx *quark.Context) error {
	var param request.UserAddressSaveReq
	if err := ctx.Bind(&param); err != nil {
		return ctx.JSONError(err.Error())
	}

	if param.Phone != "" && !utils.CheckRegex("^(13[0-9]|14[01456879]|15[0-35-9]|16[2567]|17[0-8]|18[0-9]|19[0-35-9])\\d{8}$", param.Phone) {
		return ctx.JSONError("联系电话格式不正确")
	}

	if err := service.NewUserAddressService().UpdateUserAddressById(model.UserAddress{
		Id:        param.Id,
		Name:      param.Name,
		Phone:     param.Phone,
		Province:  param.Province,
		City:      param.City,
		District:  param.District,
		Address:   param.Address,
		IsDefault: param.IsDefault,
	}); err != nil {
		return ctx.JSONError("更新用户地址失败")
	}

	return ctx.JSONOk("操作成功")
}

// 删除用户地址
func (p *UserAddress) Delete(ctx *quark.Context) error {
	var param request.UserAddressIdReq
	if err := ctx.Bind(&param); err != nil {
		return ctx.JSONError(err.Error())
	}

	if err := service.NewUserAddressService().DeleteUserAddressById(param.Id); err != nil {
		return ctx.JSONError("删除用户地址失败")
	}

	return ctx.JSONOk("操作成功")
}

package service

import (
	"github.com/quarkcloudio/quark-go/v3/dal/db"
	"github.com/quarkcloudio/quark-smart/v2/internal/dto/request"
	"github.com/quarkcloudio/quark-smart/v2/internal/dto/response"
	"github.com/quarkcloudio/quark-smart/v2/internal/model"
)

type UserAddressService struct{}

func NewUserAddressService() *UserAddressService {
	return &UserAddressService{}
}

// 获取用户地址列表
func (p *UserAddressService) GetUserAddressesByUid(param request.UserAddressIndexReq) ([]model.UserAddress, int) {
	userAddresses := make([]model.UserAddress, 0)
	var count int64

	db.Client.Model(model.UserAddress{}).
		Where("uid = ?", param.Uid).
		Count(&count).
		Offset((param.Page - 1) * param.PageSize).
		Limit(param.PageSize).
		Find(&userAddresses)

	return userAddresses, int(count)
}

// 获取用户地址详情
func (p *UserAddressService) GetUserAddressById(id int) response.UserAddressDetailResp {
	var userAddress response.UserAddressDetailResp

	db.Client.Model(model.UserAddress{}).Where("id = ?", id).Last(&userAddress)

	return userAddress
}

// 新增用户地址
func (p *UserAddressService) CreateUserAddress(userAddress model.UserAddress) error {
	return db.Client.Model(&model.UserAddress{}).Create(&userAddress).Error
}

// 更新用户地址
func (p *UserAddressService) UpdateUserAddressById(userAddress model.UserAddress) error {
	return db.Client.Model(&model.UserAddress{}).Where("id = ?", userAddress.Id).Updates(&userAddress).Error
}

// 删除用户地址
func (p *UserAddressService) DeleteUserAddressById(id int) error {
	return db.Client.Model(&model.UserAddress{}).Where("id = ?", id).Delete(&model.UserAddress{}).Error
}

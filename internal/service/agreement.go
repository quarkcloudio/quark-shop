package service

import (
	"github.com/quarkcloudio/quark-go/v3/dal/db"
	"github.com/quarkcloudio/quark-smart/v2/internal/model"
)

type AgreementService struct{}

func NewAgreementService() *AgreementService {
	return &AgreementService{}
}

// 获取协议列表
func (p *AgreementService) GetList() (list []model.Agreement, err error) {
	err = db.Client.Model(model.Agreement{}).Where("status = ?", 1).Find(&list).Error
	return
}

// 根据id获取协议内容
func (p *AgreementService) GetInfoById(id int) (data model.Agreement, err error) {
	err = db.Client.Where("status = ?", 1).Where("id = ?", id).First(&data).Error
	return
}

package service

import (
	"github.com/quarkcloudio/quark-go/v3/dal/db"
	"github.com/quarkcloudio/quark-smart/v2/internal/model"
)

type AreaService struct{}

func NewAreaService() *AreaService {
	return &AreaService{}
}

// 根据pcode获取地区列表
func (p *AreaService) GetListByPcode(pcode, level int) []model.Area {
	areas := make([]model.Area, 0)

	query := db.Client.Model(model.Area{}).Where("pcode = ?", pcode)

	if level > 0 {
		query.Where("level <= ?", level)
	}

	query.Find(&areas)

	return areas
}

// 根据code获取地区详情
func (p *AreaService) GetDetailByCode(code int) model.Area {
	var area model.Area

	db.Client.Model(model.Area{}).Where("code = ?", code).Find(&area)

	return area
}
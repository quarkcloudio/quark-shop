package model

import (
	"github.com/quarkcloudio/quark-go/v3/dal/db"
	appmodel "github.com/quarkcloudio/quark-go/v3/model"
	"github.com/quarkcloudio/quark-go/v3/service"
	"github.com/quarkcloudio/quark-go/v3/utils/datetime"
)

// 协议
type Agreement struct {
	Id        int               `json:"id" gorm:"primaryKey;autoIncrement;column:id;comment:自增ID"`
	Title     string            `json:"title" gorm:"column:title;type:varchar(200);not null;default:'';comment:协议名称"`
	Content   string            `json:"content" gorm:"column:content;type:text;default:NULL;comment:协议内容"`
	Sort      int               `json:"sort" gorm:"column:sort;type:int(10);not null;default:0;comment:排序倒序"`
	Status    int               `json:"status" gorm:"column:status;type:tinyint(1) unsigned;not null;default:1;comment:1:显示,0:不显示"`
	CreatedAt datetime.Datetime `json:"created_at"`
	UpdatedAt datetime.Datetime `json:"updated_at"`
}

// Seeder
func (m *Agreement) Seeder() {

	// 如果菜单已存在，不执行Seeder操作
	if service.NewMenuService().IsExist(110) {
		return
	}

	// 创建菜单
	menuSeeders := []*appmodel.Menu{
		{Id: 110, Name: "协议设置", GuardName: "admin", Icon: "", Type: 2, Pid: 8, Sort: 0, Path: "/api/admin/agreement/form", Show: 1, IsEngine: 1, IsLink: 0, Status: 1},
	}
	db.Client.Create(&menuSeeders)
}

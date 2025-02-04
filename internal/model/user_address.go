package model

import (
	"github.com/quarkcloudio/quark-go/v3/utils/datetime"
	"gorm.io/gorm"
)

// 用户地址模型
type UserAddress struct {
	Id        int               `json:"id" gorm:"autoIncrement"`
	Uid       int               `json:"uid"`
	Name      string            `json:"name"`
	Phone     string            `json:"phone"`
	Province  int               `json:"province"`
	City      int               `json:"city"`
	District  int               `json:"district"`
	Address   string            `json:"address"`
	IsDefault string            `json:"is_default"`
	CreatedAt datetime.Datetime `json:"created_at"`
	UpdatedAt datetime.Datetime `json:"updated_at"`
	DeletedAt gorm.DeletedAt    `json:"deleted_at"`
}

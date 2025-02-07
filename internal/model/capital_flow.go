package model

import (
	"github.com/quarkcloudio/quark-go/v3/dal/db"
	appmodel "github.com/quarkcloudio/quark-go/v3/model"
	"github.com/quarkcloudio/quark-go/v3/service"
	"github.com/quarkcloudio/quark-go/v3/utils/datetime"
)

type CapitalFlow struct {
	Id          int               `json:"id" gorm:"primaryKey;autoIncrement;column:id;comment:编号"`
	FlowNo      string            `json:"flow_no" gorm:"column:flow_id;type:varchar(32);not null;default:'';comment:流水号"`
	OrderId     string            `json:"order_id" gorm:"column:order_id;type:varchar(50);not null;default:'';comment:关联id"`
	Uid         int               `json:"uid" gorm:"column:uid;type:int(11);not null;default:0;comment:用户id"`
	Nickname    string            `json:"nickname" gorm:"column:nickname;type:varchar(255);not null;default:'';comment:昵称"`
	Phone       string            `json:"phone" gorm:"column:phone;type:varchar(20);not null;default:'';comment:电话"`
	Price       float64           `json:"price" gorm:"column:price;type:decimal(12,2);not null;default:0.00;comment:交易金额"`
	TradingType int8              `json:"trading_type" gorm:"column:trading_type;type:tinyint(1);not null;default:0;comment:交易类型:1=支付订单,2=订单退款,3=充值订单,4=充值退款,5=抽奖红包,6=佣金提现,7=购买会员"`
	PayType     string            `json:"pay_type" gorm:"column:pay_type;type:varchar(32);not null;default:'';comment:支付类型:WECHAT_PAY,ALI_PAY,BALANCE_PAY,BANK_PAY,CASH_PAY,COUPON_PAY,INTEGRAL_PAY,POINT_PAY,OTHER_PAY"`
	Mark        string            `json:"mark" gorm:"column:mark;type:varchar(500);not null;default:'';comment:备注"`
	CreatedAt   datetime.Datetime `json:"created_at" gorm:"type:datetime(0);"`
	UpdatedAt   datetime.Datetime `json:"updated_at" gorm:"type:datetime(0);"` // 记录更新时间
}

// Seeder
func (m *CapitalFlow) Seeder() {

	// 如果菜单已存在，不执行Seeder操作
	if service.NewMenuService().IsExist(97) {
		return
	}

	// 创建菜单
	menuSeeders := []*appmodel.Menu{
		{Id: 97, Name: "财务管理", GuardName: "admin", Icon: "icon-moneycollect", Type: 1, Pid: 0, Sort: 0, Path: "/bill", Show: 1, IsEngine: 0, IsLink: 0, Status: 1},
		{Id: 98, Name: "资金流水", GuardName: "admin", Icon: "", Type: 2, Pid: 97, Sort: 0, Path: "/api/admin/capitalFlow/index", Show: 1, IsEngine: 1, IsLink: 0, Status: 1},
		{Id: 99, Name: "账单记录", GuardName: "admin", Icon: "", Type: 2, Pid: 97, Sort: 0, Path: "/api/admin/billRecord/index", Show: 1, IsEngine: 1, IsLink: 0, Status: 1},
	}
	db.Client.Create(&menuSeeders)
}

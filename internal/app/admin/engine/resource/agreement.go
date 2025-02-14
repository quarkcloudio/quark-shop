package resource

import (
	"strconv"

	"github.com/quarkcloudio/quark-go/v3"
	"github.com/quarkcloudio/quark-go/v3/app/admin/actions"
	"github.com/quarkcloudio/quark-go/v3/dal/db"
	"github.com/quarkcloudio/quark-go/v3/template/admin/component/tabs"
	"github.com/quarkcloudio/quark-go/v3/template/admin/resource"
	"github.com/quarkcloudio/quark-smart/v2/internal/model"
	"github.com/quarkcloudio/quark-smart/v2/internal/service"
	"gorm.io/gorm"
)

type Agreement struct {
	resource.Template
}

// 初始化
func (p *Agreement) Init(ctx *quark.Context) interface{} {
	p.Form.SetLabelCol(map[string]interface{}{
		"span": 0,
	}).SetWrapperCol(map[string]interface{}{
		"span": 24,
	}).SetButtonWrapperCol(map[string]interface{}{
		"offset": 10,
		"span":   14,
	})

	// 标题
	p.Title = "协议设置"

	// 模型
	p.Model = &model.Agreement{}

	return p
}

// 字段
func (p *Agreement) Fields(ctx *quark.Context) []interface{} {
	field := &resource.Field{}
	agreements, _ := service.NewAgreementService().GetList()

	tabPanes := []interface{}{}
	for _, agreement := range agreements {
		fields := []interface{}{}
		getField := field.Editor(strconv.Itoa(agreement.Id)).SetDefault(agreement.Content)
		fields = append(fields, getField)
		tabPane := (&tabs.TabPane{}).
			Init().
			SetTitle(agreement.Title).
			SetBody(fields)
		tabPanes = append(tabPanes, tabPane)
	}

	return tabPanes
}

// 行为
func (p *Agreement) Actions(ctx *quark.Context) []interface{} {
	return []interface{}{
		actions.FormSubmit(),
		actions.FormReset(),
		actions.FormBack(),
		actions.FormExtraBack(),
	}
}

func (p *Agreement) FormHandle(ctx *quark.Context, query *gorm.DB, data map[string]interface{}) error {
	result := true
	for k, v := range data {
		updateResult := db.Client.Model(&model.Agreement{}).Where("id", k).Update("content", v)
		if updateResult.Error != nil {
			result = false
		}
	}
	if !result {
		return ctx.CJSONError("操作失败，请重试")
	}
	return ctx.CJSONOk("操作成功")
}

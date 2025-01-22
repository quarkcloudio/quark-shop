package handler

import (
	"github.com/quarkcloudio/quark-go/v3"
	"github.com/quarkcloudio/quark-smart/v2/internal/dto/request"
	"github.com/quarkcloudio/quark-smart/v2/internal/dto/response"
	"github.com/quarkcloudio/quark-smart/v2/internal/service"
	"github.com/quarkcloudio/quark-smart/v2/pkg/utils"
)

// 结构体
type Item struct{}

// 商品列表
func (p *Item) Index(ctx *quark.Context) error {
	var param request.ItemIndexQueryReq
	if err := ctx.Bind(&param); err != nil {
		return ctx.JSONError(err.Error())
	}

	// 构建排序规则
	param.OrderRule = param.OrderByColumn + " ASC"
	if param.IsDesc {
		param.OrderRule = param.OrderByColumn + " DESC"
	}

	list := make([]response.ItemIndexResp, 0)

	items, total := service.NewItemService().GetItemPage(param)
	for _, item := range items {
		list = append(list, response.ItemIndexResp{
			Id:         item.Id,
			Name:       item.Name,
			Image:      utils.GetImagePath(item.Image),
			Price:      item.Price,
			FictiSales: item.FictiSales,
		})
	}

	return ctx.JSONOk("ok", response.PageResp{
		List:  list,
		Total: total,
	})
}

// 商品详情
func (p *Item) Detail(ctx *quark.Context) error {
	return ctx.JSONOk("Hello, world!")
}

// 商品分类
func (p *Item) Category(ctx *quark.Context) error {
	itemService := service.NewItemService()
	itemCategories := itemService.GetCategoriesByPid(0)
	for index, itemCategory := range itemCategories {
		itemCategory.Title = "全部商品"
		itemCategory.CoverId = utils.GetImagePath(itemCategory.CoverId)
		itemCategories[index].Children = append(itemCategories[index].Children, itemCategory)
		children := itemService.GetCategoriesByPid(itemCategory.Id)
		for _, child := range children {
			child.CoverId = utils.GetImagePath(child.CoverId)
			itemCategories[index].Children = append(itemCategories[index].Children, child)
		}
	}
	return ctx.JSONOk("ok", itemCategories)
}

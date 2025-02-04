package router

import (
	"github.com/quarkcloudio/quark-go/v3"
	"github.com/quarkcloudio/quark-smart/v2/internal/app/miniapp/handler"
	"github.com/quarkcloudio/quark-smart/v2/internal/middleware"
)

// 注册MiniApp路由
func MiniAppRegister(b *quark.Engine) {

	// 不需要认证路由组
	g := b.Group("/api/miniapp")
	g.GET("/index/index", (&handler.Index{}).Index)
	g.POST("/register/register", (&handler.Register{}).Register)
	g.POST("/login/login", (&handler.Login{}).Login)
	g.GET("/login/mockLogin", (&handler.Login{}).MockLogin)

	// 轮播组
	g.GET("/index/banner", (&handler.Index{}).Banner) // 轮播列表

	// 商品组
	g.GET("/item/category", (&handler.Item{}).Category) // 商品分类
	g.GET("/item/index", (&handler.Item{}).Index)       // 商品列表
	g.GET("/item/detail", (&handler.Item{}).Detail)     // 商品详情

	// 需要登录认证路由组
	ag := b.Group("/api/miniapp", middleware.MiniAppMiddleware)
	ag.GET("/user/index", (&handler.User{}).Index)
	ag.POST("/user/save", (&handler.User{}).Save)
	ag.POST("/user/delete", (&handler.User{}).Delete)
	
	ag.GET("/area/options", (&handler.Area{}).Options) // 地区选项/省市县三级联动

	ag.GET("/userAddress/index", (&handler.UserAddress{}).Index) // 用户地址列表
	ag.GET("/userAddress/detail", (&handler.UserAddress{}).Detail) // 用户地址详情
	ag.POST("/userAddress/create", (&handler.UserAddress{}).Create) // 创建用户地址
	ag.POST("/userAddress/update", (&handler.UserAddress{}).Update) // 更新用户地址
	ag.POST("/userAddress/delete", (&handler.UserAddress{}).Delete) // 删除用户地址

	// 订单路由组
	ag.GET("/order/index", (&handler.Order{}).Index)
	ag.GET("/order/detail", (&handler.Order{}).Detail)
	ag.POST("/order/submit", (&handler.Order{}).Submit)
	ag.GET("/order/cancel", (&handler.Order{}).Cancel)
}

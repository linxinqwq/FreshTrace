package router

import (
	"backed/api/handler"
	"backed/api/middleware"
	"github.com/gin-gonic/gin"
)

// MainRouter 项目总路由
func MainRouter(engine *gin.Engine) {
	// 配置全局跨域中间键
	engine.Use(middleware.CorsHandler)

	// 用户模块路由
	userRouter(engine.Group("/user"))

	// 项目总路由
	mainRouter(engine.Group("/main"))
}

// userRouter 用户模块路由
func userRouter(router *gin.RouterGroup) {
	// 用户注册接口
	router.POST("/register", handler.UserRegister)
	// 用户登录接口
	router.POST("/login", handler.UserLogin)
	// 用户找回接口
	router.POST("/retrieve", handler.UserRetrieve)
	// 用户找回接口信息错误
	router.POST("/retrieve_result", handler.UserRetrieveResult)
	// 用户更新信息接口
	router.POST("/update_user_info", middleware.JwtTokenValid, handler.UserUpdateInfo)
}

// mainRouter 项目总路由
func mainRouter(router *gin.RouterGroup) {
	// 文件上传接口
	router.POST("/upload", middleware.JwtTokenValid, handler.FileUpload)
	// 添加商品
	router.POST("/add_goods", middleware.JwtTokenValid, handler.AddGoods)
	// 获取商品列表(商家)
	router.GET("/get_manage_list", middleware.JwtTokenValid, handler.GetUserManageGoods)
	// 修改商品信息
	router.POST("/update_goods_info", middleware.JwtTokenValid, handler.UpdateGoodsInfo)
	// 下架商品
	router.POST("/remove_goods", middleware.JwtTokenValid, handler.RemoveGoods)
	// 获取销售榜前三
	router.GET("/get_sales_top_three", middleware.JwtTokenValid, handler.GetSalesTopThree)
	// 获取商品详情
	router.GET("/get_goods_detail", middleware.JwtTokenValid, handler.GetGoodInfo)
	// 购买商品接口
	router.POST("/buy_goods", middleware.JwtTokenValid, handler.BuyGoods)
	// 获取购买记录
	router.GET("/get_buy_record", middleware.JwtTokenValid, handler.GetUserOrders)
	// 获取商品列表(用户)
	router.GET("/get_goods_list", middleware.JwtTokenValid, handler.GetGoodsList)
	// 获取未审核的商品
	router.GET("/get_unchecked_goods", middleware.JwtTokenValid, handler.GetUncheckedGoods)
	// 修改审核状态
	router.POST("/update_check_status", middleware.JwtTokenValid, handler.UpdateCheckStatus)
	// 商家给商品更新进程
	router.POST("/update_goods_process", middleware.JwtTokenValid, handler.UpdateGoodsProcess)
	// 加入购物车
	router.POST("/add_to_cart", middleware.JwtTokenValid, handler.AddToCart)
	// 获取用户购物车
	router.GET("/get_user_cart", middleware.JwtTokenValid, handler.GetUserCart)
	// 移除购物车
	router.POST("/remove_from_cart", middleware.JwtTokenValid, handler.RemoveFromCart)
	// 实现购物车购买
	router.POST("/buy_from_cart", middleware.JwtTokenValid, handler.BuyFromCart)
	// 获取所有用户信息
	router.GET("/get_all_user_info", middleware.JwtTokenValid, handler.GetAllUserInfo)
	// 修改用户状态
	router.POST("/update_user_status", middleware.JwtTokenValid, handler.UpdateUserStatus)
	// 获取当前水果种类
	router.GET("/get_fruit_kind", middleware.JwtTokenValid, handler.GetFruitKind)
	// 添加当前水果种类
	router.POST("/add_fruit_kind", middleware.JwtTokenValid, handler.AddFruitKind)
	// 删除种类
	router.POST("/remove_fruit_kind", middleware.JwtTokenValid, handler.RemoveFruitKind)
	// 获取所有水果种类
	router.GET("/get_all_fruit_kind", middleware.JwtTokenValid, handler.GetAllFruitKind)
	// 获取所有订单
	router.GET("/get_all_order", middleware.JwtTokenValid, handler.GetAllOrder)
	// 审核订单
	router.POST("/check_order", middleware.JwtTokenValid, handler.CheckOrder)
}
